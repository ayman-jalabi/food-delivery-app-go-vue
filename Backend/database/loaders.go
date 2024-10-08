package database

import (
	"errors"
	"fmt"
	"main/models"
	"main/repos"
)

// InsertIngredients I used this function once only to insert all the ingredients into my database
func InsertIngredients(allIngredients []string, ingredientsRepo *repos.IngredientsRepo) error {
	err := errors.New("error")
	//counter := 0
	for _, ingredient := range allIngredients {
		//counter++
		//fmt.Println(counter, ingredient)
		err = ingredientsRepo.InsertIngredient(ingredient)
		if err != nil {
			fmt.Println("error inserting ingredient", ingredient, err)
			return err
		}
	}
	return nil
}

// InsertSuppliers I used this function once only to insert all the suppliers into my database
func InsertSuppliers(allSuppliers []models.SupplierJson, suppliersRepo repos.SupplierRepo) error {
	err := errors.New("error")
	//counter := 0

	for _, supplier := range allSuppliers {
		//counter++
		//fmt.Println(counter, supplier)
		//fmt.Println(supplier.Type)

		err = suppliersRepo.InsertSupplier(supplier)
		if err != nil {
			fmt.Println("error inserting supplier", supplier, err)
			return err
		}

	}
	return nil
}

// InsertWorkingHours I used this function once only to insert all the working hours into my database
func InsertWorkingHours(allWorkingHours []models.WorkingHoursJson, supplierRepo repos.SupplierRepo) error {
	err := errors.New("error")
	counter := 0
	for _, hours := range allWorkingHours {
		counter++
		fmt.Println(counter, hours)
		err = supplierRepo.InsertWorkingHour(hours)
		if err != nil {
			fmt.Println("error inserting hours: ", hours, err)
			return err
		}
	}
	return nil
}

// InsertItems I used this function once only to insert all the suppliers into my database
func InsertItems(allItems []models.ItemJson, itemsRepo repos.ItemRepo) error {
	err := errors.New("error")
	//counter := 0

	for _, item := range allItems {
		//counter++
		//fmt.Println(counter, supplier)
		//fmt.Println(supplier.Type)

		err = itemsRepo.InsertItem(item)
		if err != nil {
			fmt.Println("error inserting item", item, err)
			return err
		}

	}
	return nil
}
