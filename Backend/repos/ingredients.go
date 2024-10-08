package repos

import (
	"database/sql"
)

type IngredientsRepo struct {
	//you can add here the db open function.
	Db *sql.DB
}

func NewIngredientsRepo(db *sql.DB) *IngredientsRepo {
	return &IngredientsRepo{db}
}

func (inr *IngredientsRepo) InsertIngredient(ingredient string) error {
	_, err := inr.Db.Exec(`
		INSERT INTO ingredients (ingredient_name)
		VALUES ($1)
	`, ingredient)

	if err != nil {
		return err
	}

	return nil
}

func (inr *IngredientsRepo) GetAllIngredientIds() ([]int, error) {
	var ingredientIds []int
	rows, err := inr.Db.Query("SELECT ingredient_id FROM ingredients")
	// Iterate over the rows and extract ingredient_id
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ingredientIds = append(ingredientIds, id)
	}

	// Check for errors from the iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ingredientIds, nil
}
