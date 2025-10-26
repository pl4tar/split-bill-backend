package handler

import (
	"context"
	"net/http"
	"split-bill-backend/config"
	"split-bill-backend/internal/handler/controllers"
)

func Setup(cfg *config.Config, ctx context.Context) http.Handler {
	mux := http.NewServeMux()
	db := cfg.Client
	//Users
	mux.HandleFunc(GetUserByEmail, controllers.GetUserByEmail(ctx, db))
	mux.HandleFunc(PostNewUser, controllers.AddNewUser(ctx, cfg))
	mux.HandleFunc(DeleteUserByEmail, controllers.DeleteUserEmail(ctx, db))

	//Bills
	mux.HandleFunc(GetAllBill, controllers.GetAllBillsHandler(ctx, db))
	mux.HandleFunc(PostNewBill, controllers.AddNewBill(ctx, cfg))
	mux.HandleFunc(DeleteBillByID, controllers.DeleteBillByID(ctx, db))
	mux.HandleFunc(UpdateBillTitle, controllers.EditBill(ctx, db))

	//Persons
	mux.HandleFunc(GetAllPersons, controllers.GetAllPersonsHandler(ctx, db))
	mux.HandleFunc(PostNewPerson, controllers.AddNewPersonHandler(ctx, cfg))
	mux.HandleFunc(DeletePerson, controllers.DeletePersonByID(ctx, db))
	mux.HandleFunc(UpdatePersonName, controllers.EditPersonHandler(ctx, db))

	// Products
	mux.HandleFunc(GetAllProducts, controllers.GetAllProductsHandler(ctx, db))
	mux.HandleFunc(PostNewProduct, controllers.AddNewProductHandler(ctx, db))
	mux.HandleFunc(DeleteProduct, controllers.DeleteProductByID(ctx, db))

	// Calculation
	mux.HandleFunc(CalculateDebts, controllers.CalculateDebtsHandler(ctx, cfg.Client))
	return mux
}
