package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"split-bill-backend/config"
	"split-bill-backend/internal/entity"
	"split-bill-backend/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllBillsHandler(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("user_id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user id is required"))

			return
		}

		bills, err := repository.QueryGetBillsByUserID(r.Context(), db, &id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		body, err := json.Marshal(bills)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
		w.Write(body)
		w.WriteHeader(http.StatusOK)
	}
}

func AddNewBill(ctx context.Context, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Content-Type must be application/json"))

			return
		}
		var bill entity.Bills
		body := json.NewDecoder(r.Body)
		if err := body.Decode(&bill); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
		defer r.Body.Close()

		if bill.Title == "" || bill.CreatedUserID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Title or CreatedUserID is required"))

			return
		}

		_, err := repository.QueryGetUserByID(r.Context(), cfg.Client, &bill.CreatedUserID)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
		}

		err = repository.QueryCreateNewBill(r.Context(), cfg.Client, &bill)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
	}
}
