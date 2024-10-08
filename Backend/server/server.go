package server

import (
	"context"
	"fmt"
	"log"
	"main/database"
	"main/handlers"
	"main/helpers"
	"main/models"
	"main/pool"
	"main/repos"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// I added this middleware to prevent browser
// conflicts when running both frontend and backend parts of my app at the same time
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Start(accessConfig *models.AccessConfig, refreshConfig *models.RefreshConfig) {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    accessConfig.Port,
		Handler: enableCORS(mux),
	}

	userRepo := repos.NewUserRepo(db)
	itemRepo := repos.NewItemRepo(db)
	supplierRepo := repos.NewSupplierRepo(db)

	//This is a slice of all supplier IDs taken from my database, these Ids are identical to the ones
	//provided by the foodAPI endpoint:
	allSupplierIDs, err := supplierRepo.GetSupplierIDs()

	/*The commented out function below was called once to populate the relevant database table
	I'll leave it here to help clarify that I followed a similar approach when previously
	populating parts of the database that are currently static (unchanging)

	//supplierCategoryMap := helpers.GetSupplierCategories(allSupplierIDs)

	//err = supplierRepo.PopulateSupplierCategoryJunction(supplierCategoryMap)
	//if err != nil {
	//	fmt.Println("error calling repos PopulateSupplierCategoryJunction's function", err)
	//	return
	//}
	*/

	//tokenService := service.NewTokenService(accessConfig, refreshConfig)
	authHandler := handlers.NewAuthHandler(accessConfig, refreshConfig)
	userHandler := handlers.NewUserHandler(&userRepo)
	supplierHandler := handlers.NewSupplierHandler(&supplierRepo) //
	//itemHandler := handlers.NewItemHandler(&itemRepo)

	//allSuppliersFromDB, err := supplierRepo.GetAllSuppliers()

	mux.HandleFunc("GET /suppliers", supplierHandler.GetSuppliersHandler)

	//The channel will be used to store items IDs and prices
	itemIDsAndPricesCh := make(chan map[int]float32)

	// Worker pool channels
	resultCh := make(chan any)
	errorCh := make(chan error)

	// Initialize the worker pool
	workerPool := pool.NewWorkerPool(resultCh, errorCh).WithBrokerCount(5) // You can adjust the broker count

	// Start the worker pool
	workerPool.Start()

	// Appending Worker 1: Fetches item prices every minute
	workerPool.Append(func() (any, error) {
		for {
			// Fetching item IDs and itemIDsAndPrices from the foodAPI endpoint
			itemIDsAndPrices := helpers.GetItemPrices(allSupplierIDs)
			if err != nil {
				errorCh <- fmt.Errorf("worker 1 - Fetch Error: %v", err)
				return nil, err
			}

			//Send fetched itemIDsAndPrices to the next Worker who needs to use them via channel
			itemIDsAndPricesCh <- itemIDsAndPrices

			// Sleep for 1 minute and 5 seconds before fetching again
			time.Sleep(65 * time.Second)
		}
	})

	// Appending Worker 2: Checks and updates the item prices in the database
	workerPool.Append(func() (any, error) {
		for {
			//getting the item IDs and prices map from the previous worker's result channel
			itemIDsAndPrices := <-itemIDsAndPricesCh

			for itemID, ItemPrice := range itemIDsAndPrices {
				err = itemRepo.UpdateItemPrice(itemID, ItemPrice)
				if err != nil {
					errorCh <- fmt.Errorf("worker 2 - Item update Error: %v", err)
					return nil, err
				}
			}
			//below is a channel used to make sure that the next worker only broadcasts updated item prices
			resultCh <- true
			// Send updated itemIDsAndPrices to the next Worker who needs to use them via channel
			itemIDsAndPricesCh <- itemIDsAndPrices
			//I think there's no need for sleeping as reading from channels is a
			//blocking operation:
			//time.Sleep(1 * time.Minute)
		}
	})

	// Appending Worker 3: modify it to send updated price info to the frontend
	//workerPool.Append(func() (any, error) {
	//	for {
	//		finishedUpdatingPrices := <-resultCh
	//
	//		//if the item prices were updated, then broadcast them to clients
	//		if finishedUpdatingPrices != false && finishedUpdatingPrices != nil {
	//			itemIDsAndPrices := <-itemIDsAndPricesCh
	//			handlers.BroadcastJSON(itemIDsAndPrices)
	//		}
	//
	//		// Sleep for 1 minute before sending the next update
	//		//time.Sleep(60 * time.Second)
	//	}
	//})

	// Wait for jobs and handle results/errors
	go func() {
		for {
			select {
			case res := <-resultCh:
				fmt.Println("Result:", res)
			case err = <-errorCh:
				fmt.Println("Error:", err) // Already handled in the error goroutine
			}
		}
	}()

	// No Shutdown was provided for workers as they need to run indefinitely

	mux.HandleFunc("POST /email-register-check", userHandler.CheckIfEmailExists)

	//mux.HandleFunc("POST /login", authHandler.Login)
	//mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("GET /refresh", authHandler.Refresh)

	//The /profile is the protected API endpoint that requires a valid access token
	//mux.HandleFunc("GET /profile", service.ProtectedAPIRouteMiddleware(tokenService, userHandler.GetUserInfo))

	//mux.HandleFunc("GET /item", itemHandler.GetAllItems)

	// return a user's fields in JSON format
	//http.HandleFunc("GET /user", userHandler.GetUserInfo)
	//
	////create new user
	//http.HandleFunc("POST /user", userHandler.CreateUser)

	// Creating a channel to listen for OS signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	//The port used here is 8082 as provided by the .env that is accessed by the config.go
	// Starting the server in a goroutine
	go func() {
		log.Println("Starting server on", accessConfig.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", accessConfig.Port, err)
		}
	}()

	// Block until we receive a signal on the stop channel
	<-stopChan
	log.Println("Shutting down server...")

	// insertMenuItem a timer to wait for existing connections to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err = server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
