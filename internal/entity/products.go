package entity

type Products struct {
	ID    uint    `json:"product_id" example:"1"`
	Name  string  `json:"product_name" example:"beer"`
	Price float64 `json:"product_price" example:"12.99"`
	Count uint    `json:"product_count" example:"2"`
}

type ProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count uint    `json:"count"`
}
