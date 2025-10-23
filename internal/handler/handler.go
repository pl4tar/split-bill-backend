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
	return mux
}
