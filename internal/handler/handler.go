package handler

import (
	"context"
	"net/http"
	"split-bill-backend/config"
	"split-bill-backend/internal/handler/controllers"
)

func Setup(cfg *config.Config, ctx context.Context) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc(GetUserByEmail, controllers.GetUserByEmail(ctx, cfg.Client))
	mux.HandleFunc(PostNewUser, controllers.AddNewUser(ctx, cfg))
	mux.HandleFunc(DeleteUserByEmail, controllers.DeleteUserEmail(ctx, cfg.Client))

	mux.HandleFunc(GetAllBill, controllers.GetAllBillsHandler(ctx, cfg.Client))
	mux.HandleFunc(PostNewBill, controllers.AddNewBill(ctx, cfg))
	mux.HandleFunc(DeleteBillByID, controllers.DeleteBillByID(ctx, cfg.Client))
	mux.HandleFunc(UpdateBillTitle, controllers.EditBill(ctx, cfg.Client))
	return mux
}
