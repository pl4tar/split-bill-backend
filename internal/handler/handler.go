package handler

import (
	"context"
	"log/slog"
	"net/http"
	"split-bill-backend/config"
	"split-bill-backend/internal/handler/controllers"
)

func Setup(cfg *config.Config, ctx context.Context) http.Handler {
	mux := http.NewServeMux()
	db := cfg.Client
	//Users
	mux.HandleFunc(GetUserByEmail, loggingMiddleware(corsMiddleware(controllers.GetUserByEmail(ctx, db))))
	mux.HandleFunc(PostNewUser, loggingMiddleware(corsMiddleware(controllers.AddNewUser(ctx, cfg))))
	mux.HandleFunc(DeleteUserByEmail, loggingMiddleware(corsMiddleware(controllers.DeleteUserEmail(ctx, db))))

	//Bills
	mux.HandleFunc(GetAllBill, loggingMiddleware(corsMiddleware(controllers.GetAllBillsHandler(ctx, db))))
	mux.HandleFunc(PostNewBill, loggingMiddleware(corsMiddleware(controllers.AddNewBill(ctx, cfg))))
	mux.HandleFunc(DeleteBillByID, loggingMiddleware(corsMiddleware(controllers.DeleteBillByID(ctx, db))))
	mux.HandleFunc(UpdateBillTitle, loggingMiddleware(corsMiddleware(controllers.EditBill(ctx, db))))

	//Persons
	mux.HandleFunc(GetAllPersons, loggingMiddleware(corsMiddleware(controllers.GetAllPersonsHandler(ctx, db))))
	mux.HandleFunc(PostNewPerson, loggingMiddleware(corsMiddleware(controllers.AddNewPersonHandler(ctx, cfg))))
	mux.HandleFunc(DeletePerson, loggingMiddleware(corsMiddleware(controllers.DeletePersonByID(ctx, db))))
	mux.HandleFunc(UpdatePersonName, loggingMiddleware(corsMiddleware(controllers.EditPersonHandler(ctx, db))))

	// Products
	mux.HandleFunc(GetAllProducts, loggingMiddleware(corsMiddleware(controllers.GetAllProductsHandler(ctx, db))))
	mux.HandleFunc(PostNewProduct, loggingMiddleware(corsMiddleware(controllers.AddNewProductHandler(ctx, db))))
	mux.HandleFunc(DeleteProduct, loggingMiddleware(corsMiddleware(controllers.DeleteProductByID(ctx, db))))

	// Calculation
	mux.HandleFunc(CalculateDebts, loggingMiddleware(corsMiddleware(controllers.CalculateDebtsHandler(ctx, cfg.Client))))
	return mux
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logging every request

		// не логируем метод OPTIONS
		if r.Method == http.MethodOptions {
			next(w, r)
			return
		}

		ip := r.Header.Get("X-Forwarded-For")

		userAgent := r.Header.Get("User-Agent")
		slog.Info("IP: %s, Method: %s, Route: %s, Query: %s, UserAgent: %s, AuthHeader: %s",
			ip, r.Method, r.URL.Path, r.URL.Query(), userAgent, r.Header.Get("Authorization"))

		next(w, r)
	}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := map[string]bool{
			"http://localhost:5173": true,
		}
		origin := r.Header.Get("Origin")
		if allowedOrigins[origin] {
			//w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}
