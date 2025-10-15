package entity

type Products struct {
	ID    uint
	Name  string
	Price float64
	Count uint
}

type ProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count uint    `json:"count"`
}
