package handler

import "fmt"

var apiV1 = "/api/v1"

//Users
var (
	GetUserByEmail    = fmt.Sprintf("GET %s/user", apiV1)
	PostNewUser       = fmt.Sprintf("POST %s/user", apiV1)
	DeleteUserByEmail = fmt.Sprintf("DELETE %s/user", apiV1)
)

//Bills
var (
	GetAllBill      = fmt.Sprintf("GET %s/bills", apiV1)
	GetBillById     = fmt.Sprintf("GET %s/bills/{id}", apiV1)
	PostNewBill     = fmt.Sprintf("POST %s/bills", apiV1)
	DeleteBillByID  = fmt.Sprintf("DELETE %s/bills", apiV1)
	UpdateBillTitle = fmt.Sprintf("PATCH %s/bills", apiV1)
)

//Persons
var (
	GetAllPersons    = fmt.Sprintf("GET %s/persons", apiV1)
	PostNewPerson    = fmt.Sprintf("POST %s/persons", apiV1)
	DeletePerson     = fmt.Sprintf("DELETE %s/persons", apiV1)
	UpdatePersonName = fmt.Sprintf("PATCH %s/persons", apiV1)
)

// Products
var (
	GetAllProducts = fmt.Sprintf("GET %s/products", apiV1)
	PostNewProduct = fmt.Sprintf("POST %s/products", apiV1)
	DeleteProduct  = fmt.Sprintf("DELETE %s/products", apiV1)
)
