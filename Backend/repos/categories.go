package repos

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"main/models"
)

type CategoryRepo struct {
	//you can add here the db open function. you can find this example on the documentation maybe page:
	Db *sql.DB
}

func NewCategoryRepo(db *sql.DB) CategoryRepo {
	return CategoryRepo{db}
}

func (cr CategoryRepo) Create(category models.Category) error {
	_, err := cr.Db.Exec(
		"INSERT INTO categories (category_id, category_name) values ($1, $2)",
		category.ID, category.Name,
	)

	return err
}

func (cr CategoryRepo) GetALl() ([]models.Category, error) {

	var categories []models.Category

	result, err := cr.Db.Query("SELECT category_id, category_name FROM categories")
	if err != nil {
		return nil, err
	}
	for result.Next() {
		category := models.Category{}
		err := result.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}
	result.Close()

	return categories, nil
}

func (cr CategoryRepo) DeleteItem(id int) error {
	stmt, err := cr.Db.Prepare("DELETE FROM categories WHERE category_id = $1;")
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

func (cr CategoryRepo) PaginatedGetSupplierCategories(supplierId, page, pageSize int) ([]models.Category, error) {
	var categories []models.Category
	offset := (page - 1) * pageSize

	rows, err := cr.Db.Query(`SELECT supplier_id, category_id 
	FROM supplier_category_junc s 
	WHERE supplier_id = $1
	LIMIT $2 OFFSET $3
`, pageSize, offset)
	if err != nil {
		fmt.Println("error getting suppliers", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		if err = rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
