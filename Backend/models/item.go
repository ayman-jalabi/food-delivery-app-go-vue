package models

type ItemJson struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Price       float32  `json:"price"`
	Image       string   `json:"image"`
	Type        string   `json:"type"` //category
	Ingredients []string `json:"ingredients"`
	SupplierID  int
}

type ItemJsonIDAndPrice struct {
	ID    int     `json:"id"`
	Price float32 `json:"price"`
}

type Item struct {
	ID          int
	Name        string
	Price       float32
	Ingredients []string
	SupplierID  int
	CategoryID  int
}

type MenusResponse struct {
	Menu []ItemJson `json:"menu"`
}

type ItemPriceResponse struct {
	Menu []ItemJsonIDAndPrice `json:"menu"`
}
