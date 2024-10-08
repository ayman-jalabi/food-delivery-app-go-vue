package repos

import (
	"database/sql"
	"fmt"
	"log"
	"main/models"
)

type SupplierRepo struct {
	//you can add here the db open function. you can find this example on the documentation maybe page:
	Db *sql.DB
}

func NewSupplierRepo(db *sql.DB) SupplierRepo {
	return SupplierRepo{db}
}

func (sr SupplierRepo) insertWorkingHours(hours models.WorkingHours) {
	err := sr.Db.QueryRow(`
		INSERT INTO working_hours (opening, closing)
		VALUES ($1, $2)
	`, hours.Opening, hours.Closing)
	if err != nil {
		fmt.Println(err)
	}
}

func (sr SupplierRepo) InsertWorkingHour(suppliers models.WorkingHoursJson) error {
	stmt, err := sr.Db.Prepare("INSERT INTO working_hours(opening, closing) VALUES($1,$2)")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec(suppliers.Opening, suppliers.Closing)
	if err != nil {
		fmt.Println(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("affected = %d\n", rowCnt)

	return nil
}

func (sr SupplierRepo) InsertSupplier(supplier models.SupplierJson) error {
	// Step 1: Retrieve the supplier_type_id from supplier_kind table
	var supplierTypeID int
	var workingHoursID int

	query := "SELECT supplier_type_id FROM supplier_kind WHERE supplier_type_name = $1"
	err := sr.Db.QueryRow(query, supplier.Type).Scan(&supplierTypeID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case where no matching supplier_type_name is found
			return fmt.Errorf("supplier type name not found: %v", supplier.Type)
		}
		return fmt.Errorf("error querying supplier_type_id: %v", err)
	}

	// Step 2: Query to get working_hours_id
	queryTwo := "SELECT working_hours_id FROM working_hours WHERE opening = $1 AND closing = $2"
	err = sr.Db.QueryRow(queryTwo, supplier.WorkingHours.Opening, supplier.WorkingHours.Closing).Scan(&workingHoursID)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no matching working hours found, insert new working hours and get the working_hours_id
			insertQuery := `
				INSERT INTO working_hours (opening, closing)
				VALUES ($1, $2)
				RETURNING working_hours_id
			`
			err = sr.Db.QueryRow(insertQuery, supplier.WorkingHours.Opening, supplier.WorkingHours.Closing).Scan(&workingHoursID)
			if err != nil {
				return fmt.Errorf("error inserting new working hours: %v", err)
			}
		} else {
			return fmt.Errorf("error querying working_hours_id: %v", err)
		}
	}

	// Step 3: Insert into the suppliers table using the retrieved supplier_type_id
	_, err = sr.Db.Exec(
		`INSERT INTO suppliers (supplier_id ,name, supplier_type_id, image_url, working_hours_id) 
			   VALUES ($1, $2, $3, $4, $5)`,
		supplier.ID, supplier.Name, supplierTypeID, supplier.Image, workingHoursID,
	)
	if err != nil {
		return fmt.Errorf("error inserting into suppliers: %v", err)
	}

	return nil
}

func (sr SupplierRepo) PopulateSupplierCategoryJunction(supplierCategorySlice []models.SupplierCategoryJunction) error {

	for _, supplierCategory := range supplierCategorySlice {
		_, err := sr.Db.Exec(
			`INSERT INTO supplier_category_junc (supplier_id ,category_id) 
			   VALUES ($1, $2)`,
			supplierCategory.SupplierID, supplierCategory.CategoryID,
		)
		if err != nil {
			return fmt.Errorf("error populating supplier_category_junc: %v", err)
		}
	}

	return nil
}

func (sr SupplierRepo) GetSupplierIDs() ([]int, error) {
	var allSupplierIDs []int
	stmt, err := sr.Db.Prepare("SELECT supplier_id FROM suppliers")
	if err != nil {
		fmt.Println("error getting supplier ID: ", err)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("error querying supplier IDs: ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var supplierID int
		if err = rows.Scan(&supplierID); err != nil {
			fmt.Println("error scanning supplier ID: ", err)
			return nil, err
		}
		allSupplierIDs = append(allSupplierIDs, supplierID)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("rows error: ", err)
		return nil, err
	}

	return allSupplierIDs, nil
}

func (sr SupplierRepo) GetAllSuppliers() (*[]models.Supplier, error) {
	var allSuppliers []models.Supplier
	stmt, err := sr.Db.Prepare(`SELECT supplier_id, supplier_type_name, name, image_url, h.opening, h.closing 
	FROM suppliers s
	JOIN supplier_kind k ON s.supplier_type_id = k.supplier_type_id
    JOIN working_hours h ON s.working_hours_id = h.working_hours_id
`)
	if err != nil {
		fmt.Println("error getting supplier info: ", err)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("error querying supplier info: ", err)
		return nil, err
	}

	defer rows.Close()

	//copy supplier info from the DB and append it to the allSuppliers slice
	for rows.Next() {
		supplier := models.Supplier{}

		if err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Type,
			&supplier.Image, &supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing); err != nil {
			fmt.Println("error scanning supplier data: ", err)
			return nil, err
		}

		allSuppliers = append(allSuppliers, supplier)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("rows error: ", err)
		return nil, err
	}

	return &allSuppliers, nil
}

func (sr SupplierRepo) GetSupplier() (*models.Supplier, error) {
	var supplier models.Supplier

	stmt, err := sr.Db.Prepare(`SELECT supplier_id, supplier_type_name, name, image_url, h.opening, h.closing 
	FROM suppliers s
	JOIN supplier_kind k ON s.supplier_type_id = k.supplier_type_id
    JOIN working_hours h ON s.working_hours_id = h.working_hours_id
`)
	if err != nil {
		fmt.Println("error getting supplier info: ", err)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("error querying supplier info: ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Type,
			&supplier.Image, &supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing); err != nil {
			fmt.Println("error scanning supplier data: ", err)
			return nil, err
		}

	}

	if err = rows.Err(); err != nil {
		fmt.Println("rows error: ", err)
		return nil, err
	}

	return &supplier, nil
}

func (sr SupplierRepo) PaginatedGetSuppliers(page, pageSize int) ([]models.Supplier, error) {
	var suppliers []models.Supplier
	offset := (page - 1) * pageSize

	rows, err := sr.Db.Query(`SELECT supplier_id, name, supplier_type_name,image_url, h.opening, h.closing 
	FROM suppliers s 
	JOIN supplier_kind k ON s.supplier_type_id = k.supplier_type_id
    JOIN working_hours h ON s.working_hours_id = h.working_hours_id
	LIMIT $1 OFFSET $2
`, pageSize, offset)
	if err != nil {
		fmt.Println("error getting suppliers", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var supplier models.Supplier
		if err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.Type, &supplier.Image, &supplier.WorkingHours.Opening, &supplier.WorkingHours.Closing); err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

func (sr SupplierRepo) GetTotalOfSuppliers() (int, error) {
	var total int
	query := "SELECT COUNT(*) FROM suppliers"

	// Execute the query
	err := sr.Db.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total suppliers: %w", err)
	}

	return total, nil
}
