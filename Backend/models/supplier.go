package models

type SuppliersResponse struct {
	Suppliers []SupplierJson `json:"suppliers"`
}

type WorkingHoursJson struct {
	Opening string `json:"opening"`
	Closing string `json:"closing"`
}

type SupplierJson struct {
	ID           int              `json:"id"`
	Name         string           `json:"name"`
	Type         string           `json:"type"`
	Image        string           `json:"image"`
	WorkingHours WorkingHoursJson `json:"workingHours"`
}

type Supplier struct {
	ID           int
	Name         string
	Type         string
	Image        string
	WorkingHours WorkingHours
}

type WorkingHours struct {
	Opening string
	Closing string
}

// SupplierType the supplier type ID,Name:
/*
	1,restaurant
	2,coffee_shop
	3,supermarket
	4,bar
	5,shop
*/
type SupplierType struct {
	SupplierID       int
	SupplierTypeName string
}

type SupplierCategoryJunction struct {
	SupplierID int
	CategoryID int
}
