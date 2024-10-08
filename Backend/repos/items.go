package repos

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"main/models"
)

type ItemRepo struct {
	Db *sql.DB
}

func NewItemRepo(db *sql.DB) ItemRepo {
	return ItemRepo{db}
}

func (ir ItemRepo) InsertItem(item models.ItemJson) error {

	var categoryTypeID int
	var ingredientID int
	//var originalName string
	//var counter int

	//Query to get item category and store it in the var categoryTypeID
	query := "SELECT category_id FROM categories WHERE category_name = $1"
	err := ir.Db.QueryRow(query, item.Type).Scan(&categoryTypeID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle case where no matching supplier_type_name is found
			return fmt.Errorf("item type name not found: %v", item.Type)
		}
		return fmt.Errorf("error querying category_id: %v", err)
	}

	// Insert item into the item table
	_, err = ir.Db.Exec(
		`INSERT INTO item (item_id, item_name, price, supplier_id, category_id, image_url) 
			   VALUES ($1, $2, $3, $4, $5, $6)
			   ON CONFLICT (item_id) DO NOTHING`,
		item.ID, item.Name, item.Price, item.SupplierID, categoryTypeID, item.Image,
	)
	if err != nil {
		return fmt.Errorf("error inserting into item: %v", err)
	}

	//insert item and ingredient IDs into the item_ingredients table
	for _, element := range item.Ingredients {
		ingredientIdQuery := "SELECT ingredient_id FROM ingredients WHERE ingredient_name = $1"
		err = ir.Db.QueryRow(ingredientIdQuery, element).Scan(&ingredientID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				// Handle case where no matching supplier_type_name is found
				return fmt.Errorf("ingredient not found: %v", element)
			}
			return fmt.Errorf("error querying ingredient_id: %v", err)
		}

		_, err = ir.Db.Exec(
			`INSERT INTO item_ingredients (item_id, ingredient_id) 
			   VALUES ($1, $2)
			   ON CONFLICT (item_id,ingredient_id) DO NOTHING`,
			item.ID, ingredientID,
		)
		if err != nil {
			return fmt.Errorf("error inserting into item_ingredients: %v", err)
		}
	}
	return nil
}

func (ir ItemRepo) GetAllItemIDAndPrice() ([]models.ItemJsonIDAndPrice, error) {

	var itemsIDsAndPrices []models.ItemJsonIDAndPrice

	result, err := ir.Db.Query("SELECT item_id, price FROM item")
	if err != nil {
		fmt.Println("error querying items:", err)
		return nil, err
	}
	for result.Next() {
		item := models.ItemJsonIDAndPrice{}
		err = result.Scan(&item.ID, &item.Price)
		if err != nil {
			fmt.Println("error getting item ID and price from database", err)
			return nil, err
		}

		itemsIDsAndPrices = append(itemsIDsAndPrices, item)
	}
	result.Close()

	return itemsIDsAndPrices, nil
}

func (ir ItemRepo) UpdateItemPrice(itemID int, newItemPrice float32) error {
	var oldItemPrice float32
	//Query to get item category and store it in the var categoryTypeID
	query := "SELECT price FROM item WHERE item_id = $1"
	err := ir.Db.QueryRow(query, itemID).Scan(&oldItemPrice)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			return fmt.Errorf("item ID not found: %v", itemID)
		}
		return fmt.Errorf("error querying category_id: %v", err)
	}

	// Comparing old price with new price and updating it if needed
	if oldItemPrice != newItemPrice {
		// If the price has changed, update it in the database
		updateQuery := "UPDATE item SET price = $1 WHERE item_id = $2"
		_, err = ir.Db.Exec(updateQuery, newItemPrice, itemID)
		if err != nil {
			return fmt.Errorf("error updating item price: %v", err)
		}
		//I can use the printf below to test that the prices are really changing
		//fmt.Printf("Item price updated for itemID %d: old price: %f, new price: %f\n", itemID, oldItemPrice, newItemPrice)
	}
	return nil
}

func (ir ItemRepo) DeleteItem(id int) error {
	stmt, err := ir.Db.Prepare("DELETE FROM item WHERE item_id = $1;")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	log.Printf("ID = %d, number of affected rows= %d\n", id, rowCnt)

	return err
}

//func (ir ItemRepo) GetAllItems() ([]models.Item, error) {
//	var allItems []models.Item
//
//	stmt, err := ir.Db.Prepare(`SELECT item_id, item_name, price, image_url, supplier_id, category_id
//	FROM item i
//	JOIN item_ingredients i_ing ON i.i = k.supplier_type_id
//    JOIN working_hours h ON s.working_hours_id = h.working_hours_id
//`)
//	if err != nil {
//		fmt.Println("error getting supplier info: ", err)
//		return nil, err
//	}
//
//	defer stmt.Close()
//
//	rows, err := stmt.Query()
//	if err != nil {
//		fmt.Println("error querying supplier info: ", err)
//		return nil, err
//	}
//
//	defer rows.Close()
//
//	//copy supplier info from the DB and append it to the allSuppliers slice
//	for rows.Next() {
//		item := models.Item{}
//
//		if err = rows.Scan(&item.Ingredients); err != nil {
//			fmt.Println("error scanning supplier data: ", err)
//			return nil, err
//		}
//
//		allSuppliers = append(allSuppliers, supplier)
//	}
//
//	if err = rows.Err(); err != nil {
//		fmt.Println("rows error: ", err)
//		return nil, err
//	}
//
//	return allSuppliers, nil
//}
