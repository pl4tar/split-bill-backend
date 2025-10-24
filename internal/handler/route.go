package handler

import "fmt"

var apiV1 = "/api/v1"

var (
	GetUserByEmail    = fmt.Sprintf("GET %s/user", apiV1)
	PostNewUser       = fmt.Sprintf("POST %s/user", apiV1)
	DeleteUserByEmail = fmt.Sprintf("DELETE %s/user", apiV1)
)

var (
	GetAllBill      = fmt.Sprintf("GET %s/bills", apiV1)
	GetBillById     = fmt.Sprintf("GET %s/bills/{id}", apiV1)
	PostNewBill     = fmt.Sprintf("POST %s/bills", apiV1)
	DeleteBillByID  = fmt.Sprintf("DELETE %s/bills", apiV1)
	UpdateBillTitle = fmt.Sprintf("PATCH %s/bills", apiV1)
)
