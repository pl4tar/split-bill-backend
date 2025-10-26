package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"split-bill-backend/internal/entity"
	"split-bill-backend/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllProductsHandler(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("bill_id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user id is required"))

			return
		}

		products, err := repository.QueryGetProductsIOByBillID(ctx, db, id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		body, err := json.Marshal(products)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func AddNewProductHandler(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Content-Type must be application/json"))

			return
		}

		var product entity.ProductsIO
		body := json.NewDecoder(r.Body)
		if err := body.Decode(&product); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		defer r.Body.Close()

		if product.BillID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bill id is required"))

			return
		}

		// TODO Добавить валидацию
		// я конечно понимаю, что фронт не должен давать выбраит пользователя которого не существует,
		// но все же было бы хорошо добавить валидацию
		err := repository.QueryCreateNewProduct(ctx, db, &product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Product added successfully"))
	}
}

func DeleteProductByID(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Content-Type must be application/json"))
			return
		}

		var request struct {
			ProductID uint `json:"product_id,string"`
		}

		body := json.NewDecoder(r.Body)
		if err := body.Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		if request.ProductID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Product ID is required"))
			return
		}

		err := repository.QueryDeleteProductByID(r.Context(), db, request.ProductID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Product deleted successfully"))
	}
}
