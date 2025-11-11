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

// GetAllBillsHandler
// @Summary Получение всех счетов по ID пользователя
// @Description Возвращает список счетов для указанного пользователя
// @Tags bills
// @Produce json
// @Param user_id query string true "ID пользователя"
// @Success 200 {array} entity.Bills "Список счетов"
// @Failure 400 {string} string "Ошибка запроса"
// @Router /bills [get]
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

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

// AddNewBill
// @Summary Создание нового счета
// @Description Создает новый счет для пользователя
// @Tags bills
// @Accept json
// @Produce json
// @Param bill body entity.Bills true "Данные счета"
// @Success 200 {string} string "Счет успешно создан"
// @Failure 400 {string} string "Ошибка запроса"
// @Router /bills [post]
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
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("unknown user"))

			return
		}

		err = repository.QueryCreateNewBill(r.Context(), cfg.Client, &bill)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
	}
}

// DeleteBillByID
// @Summary Удаление счета по ID
// @Description Удаляет счет по указанному идентификатору
// @Tags bills
// @Accept json
// @Produce json
// @Param request body entity.BillDel true "ID счета для удаления"
// @Success 200 {string} string "Счет успешно удален"
// @Failure 400 {string} string "Ошибка запроса"
// @Router /bills [delete]
func DeleteBillByID(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Content-Type must be application/json"))

			return
		}

		var request struct {
			Bill_id uint `json:"bill_id,string"`
		}

		body := json.NewDecoder(r.Body)
		if err := body.Decode(&request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
		defer r.Body.Close()

		if request.Bill_id == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No Content"))

			return
		}

		err := repository.QueryDeleteBillByID(r.Context(), db, request.Bill_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Bill deleted successfully"))
	}
}

// EditBill
// @Summary Редактирование счета
// @Description Обновляет данные счета
// @Tags bills
// @Accept json
// @Produce json
// @Param bill body entity.Bills true "Обновленные данные счета"
// @Success 200 {string} string "Счет успешно обновлен"
// @Failure 400 {string} string "Ошибка запроса"
// @Router /bills [put]
func EditBill(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
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

		if bill.ID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bill ID is required"))
			return
		}

		if bill.Title == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bill Title is required"))
			return
		}

		if bill.CreatedUserID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bill CreatedUserID is required"))
			return
		}

		_, err := repository.QueryGetUserByID(r.Context(), db, &bill.CreatedUserID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("unknown user"))

			return
		}

		err = repository.QueryEditTitle(r.Context(), db, &bill)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}
	}

}
