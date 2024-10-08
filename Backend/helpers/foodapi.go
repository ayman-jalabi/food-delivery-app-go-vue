package helpers

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/models"
	"net/http"
	"os"
	"time"
)

// SupplierReplyLimit add here the maximum number of suppliers per displayed page
var SupplierReplyLimit = 20

// FetchSuppliers fetches supplier data from the given API.
func FetchSuppliers(SupplierReplyLimit, SupplierPageNumber int) ([]models.SupplierJson, error) {
	// Construct the API URL with the provided limit and page parameters.
	apiURL := fmt.Sprintf("https://foodapi.golang.nixdev.co/suppliers?limit=%d&page=%d", SupplierReplyLimit, SupplierPageNumber)

	// insertMenuItem a new HTTP client with a timeout and skip TLS verification.
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // Skip TLS certificate verification (not recommended for production).
			},
		},
	}

	// Send a GET request to the API.
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 OK.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON data into the SuppliersResponse struct.
	var suppliersResponse models.SuppliersResponse
	if err = json.Unmarshal(body, &suppliersResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Return the suppliers' data.
	return suppliersResponse.Suppliers, nil
}

// FetchMenus fetches supplier menu data from the given API.
func FetchMenus(supplierID int) ([]models.ItemJson, error) {
	// Construct the API URL with the provided limit and page parameters.
	apiURL := fmt.Sprintf("https://foodapi.golang.nixdev.co/suppliers/%d/menu", supplierID)

	// insertMenuItem a new HTTP client with a timeout and skip TLS verification.
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // Skip TLS certificate verification (not recommended for production).
			},
		},
	}

	// Send a GET request to the API.
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 OK.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	//fmt.Println(string(body))

	// Parse the JSON data into the SuppliersResponse struct.
	var menusResponse models.MenusResponse
	if err = json.Unmarshal(body, &menusResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Return the Menus' data.
	return menusResponse.Menu, nil
}

// FetchItemPrices fetches a supplier's menu item's ID and price data from the given API.
func FetchItemPrices(supplierID int) ([]models.ItemJsonIDAndPrice, error) {
	// Construct the API URL with the provided limit and page parameters.
	apiURL := fmt.Sprintf("https://foodapi.golang.nixdev.co/suppliers/%d/menu", supplierID)

	// insertMenuItem a new HTTP client with a timeout and skip TLS verification.
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // Skip TLS certificate verification (not recommended for production).
			},
		},
	}

	// Send a GET request to the API.
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 OK.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON data
	var menusResponse models.ItemPriceResponse
	if err = json.Unmarshal(body, &menusResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Return the Menus' item ID and price data.
	return menusResponse.Menu, nil
}

func GetSuppliersAndTheirIds() ([]int, []models.SupplierJson) {

	onePageOfSuppliers := []models.SupplierJson{}
	AllSuppliers := []models.SupplierJson{}
	AllSupplierIds := []int{}
	err := errors.New("error")
	supplierPageNum := 1

	for supplierPageNum < 7 {
		onePageOfSuppliers, err = FetchSuppliers(20, supplierPageNum)
		if err != nil {
			fmt.Println("error fetching onePageOfSuppliers", err.Error())
		}
		for _, supplier := range onePageOfSuppliers {
			AllSupplierIds = append(AllSupplierIds, supplier.ID)
			AllSuppliers = append(AllSuppliers, supplier)
		}
		supplierPageNum++
	}
	return AllSupplierIds, AllSuppliers
}

func GetItemPrices(allSupplierIds []int) map[int]float32 {
	err := errors.New("error")
	singleCallMenus := []models.ItemJsonIDAndPrice{}
	//var itemIDAndPrice map[int]float32
	itemIDAndPrice := make(map[int]float32)

	for _, supplierId := range allSupplierIds {
		singleCallMenus, err = FetchItemPrices(supplierId)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error fetching supplier %v menu data: %v \n", supplierId, err)
			//fmt.Println("error fetching supplier data", err)
		}
		for _, item := range singleCallMenus {
			itemIDAndPrice[item.ID] = item.Price
		}
	}

	return itemIDAndPrice
}

func GetAllItems(allSupplierIds []int) []models.ItemJson {
	//counter := 0
	err := errors.New("error")
	singleCallMenus := []models.ItemJson{}
	allItems := []models.ItemJson{}

	for _, supplierId := range allSupplierIds {
		singleCallMenus, err = FetchMenus(supplierId)
		if err != nil {
			fmt.Fprintln(os.Stdout, "error fetching singleSupplierMenu: ", supplierId, "\n", err.Error())
		}
		for _, item := range singleCallMenus {
			//counter++
			//fmt.Println(counter)
			item.SupplierID = supplierId
			allItems = append(allItems, item)
		}
	}

	return allItems
}

func GetSupplierCategories(allSupplierIds []int) []models.SupplierCategoryJunction {
	err := errors.New("error")
	singleCallMenuItems := []models.ItemJson{}
	sliceOfSupplierIDAndCategories := []models.SupplierCategoryJunction{}

	// Create a map to track unique SupplierID and CategoryID combinations
	seen := make(map[string]struct{})

	for _, supplierId := range allSupplierIds {
		singleCallMenuItems, err = FetchMenus(supplierId)
		if err != nil {
			fmt.Fprintln(os.Stdout, "error fetching singleSupplierMenu: ", supplierId, "\n", err.Error())
			continue // Skip to the next supplierId on error
		}
		for _, item := range singleCallMenuItems {
			var categoryId int
			switch item.Type {
			case "pizza":
				categoryId = 1
			case "burger":
				categoryId = 2
			case "sushi":
				categoryId = 3
			case "frozen_meal":
				categoryId = 4
			case "appetizer":
				categoryId = 5
			case "dessert":
				categoryId = 6
			default:
				categoryId = 7
			}

			// Create a unique key for the SupplierID and CategoryID combination
			key := fmt.Sprintf("%d-%d", supplierId, categoryId)

			// Check if this combination has already been seen
			if _, exists := seen[key]; !exists {
				// Mark this combination as seen
				seen[key] = struct{}{}

				// Creating a new SupplierCategoryJunction for each unique entry
				supplierCategory := models.SupplierCategoryJunction{
					SupplierID: supplierId,
					CategoryID: categoryId,
				}

				// Appending the new entry to the slice
				sliceOfSupplierIDAndCategories = append(sliceOfSupplierIDAndCategories, supplierCategory)
			}
		}
	}

	return sliceOfSupplierIDAndCategories
}

func GetAllIngredients(allItems []models.ItemJson) []string {
	//counter := 0
	AllIngredients := []string{}

	for _, item := range allItems {
		aSingleItemsIngredients := item.Ingredients
		for _, ingredient := range aSingleItemsIngredients {
			//counter++
			//fmt.Println(counter)
			if StringContains(AllIngredients, ingredient) {
				continue
			}
			AllIngredients = append(AllIngredients, ingredient)
		}
	}
	return AllIngredients
}

func GetWorkingHours(allSuppliers []models.SupplierJson) []models.WorkingHoursJson {

	allWorkingHours := []models.WorkingHoursJson{}

	for _, supplier := range allSuppliers {
		if OpeningAndClosingHoursContains(allWorkingHours, supplier.WorkingHours.Opening, supplier.WorkingHours.Closing) {
			continue
		}
		allWorkingHours = append(allWorkingHours, supplier.WorkingHours)
	}
	fmt.Println(len(allWorkingHours))
	return allWorkingHours
}
