package controllers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"split-bill-backend/config"
	"split-bill-backend/internal/entity"
	"split-bill-backend/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserByEmail(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("email is required"))

			return
		}

		user, err := repository.QueryGetUserByEmail(r.Context(), db, &email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		body, err := json.Marshal(user)
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

func AddNewUser(ctx context.Context, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Content-Type must be application/json"))

			return
		}
		var user entity.Users
		body := json.NewDecoder(r.Body)
		if err := body.Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
		defer r.Body.Close()

		if user.Email == "" || user.Password == "" || user.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("email or password is required"))

			return
		}

		if user.Email == "" && user.Password == "" && user.Name == "" {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("No Content"))

			return
		}

		_, exists := repository.QueryGetUserByEmail(r.Context(), cfg.Client, &user.Email)
		if exists == nil {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("email already exists"))
			return
		}

		err := repository.QuerySaveNewUser(r.Context(), cfg.Client, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
	}
}

func DeleteUserEmail(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Content-Type must be application/json"))

			return
		}

		var request struct {
			Email string `json:"email"`
		}

		body := json.NewDecoder(r.Body)
		if err := body.Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
		defer r.Body.Close()
		slog.Info(request.Email)

		if request.Email == "" {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("No Content"))

			return
		}

		err := repository.QueryDeleteUser(r.Context(), db, request.Email)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

	}
}
