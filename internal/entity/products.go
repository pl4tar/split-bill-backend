package entity

type ProductsIO struct {
	BillID   uint      `json:"bill_id,string" example:"1"`
	Products []Product `json:"products,omitempty"`
}

type Product struct {
	ID      uint      `json:"product_id,string" example:"1"`
	Name    string    `json:"product_name" example:"beer"`
	Price   float64   `json:"product_price" example:"12.99"`
	Count   uint      `json:"product_count" example:"2"`
	Persons []Persons `json:"persons,omitempty" example:"[{'id': 1, 'name': 'abc'}]"`
	PayerID uint      `json:"payer_id,string" example:"1"`
}
