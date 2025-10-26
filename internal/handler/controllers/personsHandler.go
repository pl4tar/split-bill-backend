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

func GetAllPersonsHandler(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("user_id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user id is required"))

			return
		}

		persons, err := repository.QueryGetPersonsByBillID(r.Context(), db, id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		body, err := json.Marshal(persons)
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

func AddNewPersonHandler(ctx context.Context, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Content-Type must be application/json"))

			return
		}
		var person entity.Persons
		body := json.NewDecoder(r.Body)
		if err := body.Decode(&person); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
		defer r.Body.Close()
		if person.Name == "" || person.BillID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Name or owner id is required"))

			return
		}

		_, err := repository.QueryGetBillByID(r.Context(), cfg.Client, &person.BillID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("unknown user"))

			return
		}

		err = repository.QueryCreateNewPerson(r.Context(), cfg.Client, &person)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
	}
}

func DeletePersonByID(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Content-Type must be application/json"))

			return
		}

		var request struct {
			Person_id uint `json:"id,string"`
		}

		body := json.NewDecoder(r.Body)
		if err := body.Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
		defer r.Body.Close()

		if request.Person_id == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No Content"))

			return
		}

		err := repository.QueryDeletePersonByID(r.Context(), db, &request.Person_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

	}
}

func EditPersonHandler(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Content-Type must be application/json"))

			return
		}

		var person entity.Persons
		body := json.NewDecoder(r.Body)
		if err := body.Decode(&person); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
		defer r.Body.Close()

		if person.ID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Person ID is required"))
			return
		}

		if person.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Person name is required"))
			return
		}

		if person.BillID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Person owner id is required"))
			return
		}

		_, userExist := repository.QueryGetUserByID(r.Context(), db, &person.BillID)
		if userExist != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("unknown user"))

			return
		}

		// TODO: Сделать проверку на существоваание во всех похожих хендлерах

		err := repository.QueryEditName(r.Context(), db, &person)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
	}
}
