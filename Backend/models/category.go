package models

type CategoryJson struct {
	Name string `json:"category_name"`
}

type Category struct {
	ID   string
	Name string
}

// Categories the category ID,Name:
/*
	1,pizza
	2,burger
	3,sushi
	4,frozen_meal
	5,appetizer
	6,dessert
	7,pastry
*/
type Categories struct {
	CategoryID   int
	CategoryName string
}
