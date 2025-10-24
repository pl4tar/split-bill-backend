package handler

import (
	"context"
	"net/http"
	"split-bill-backend/config"
	"split-bill-backend/internal/handler/controllers"
)

func Setup(cfg *config.Config, ctx context.Context) http.Handler {
	mux := http.NewServeMux()

	//Users
	mux.HandleFunc(GetUserByEmail, controllers.GetUserByEmail(ctx, cfg.Client))
	mux.HandleFunc(PostNewUser, controllers.AddNewUser(ctx, cfg))
	mux.HandleFunc(DeleteUserByEmail, controllers.DeleteUserEmail(ctx, cfg.Client))

	//Bills
	mux.HandleFunc(GetAllBill, controllers.GetAllBillsHandler(ctx, cfg.Client))
	mux.HandleFunc(PostNewBill, controllers.AddNewBill(ctx, cfg))
	mux.HandleFunc(DeleteBillByID, controllers.DeleteBillByID(ctx, cfg.Client))
	mux.HandleFunc(UpdateBillTitle, controllers.EditBill(ctx, cfg.Client))

	//Persons
	mux.HandleFunc(GetAllPersons, controllers.GetAllPersonsHandler(ctx, cfg.Client))
	mux.HandleFunc(PostNewPerson, controllers.AddNewPersonHandler(ctx, cfg))
	mux.HandleFunc(DeletePerson, controllers.DeletePersonByID(ctx, cfg.Client))
	mux.HandleFunc(UpdatePersonName, controllers.EditPersonHandler(ctx, cfg.Client))

	return mux
}
