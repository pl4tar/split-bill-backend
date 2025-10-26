package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"split-bill-backend/internal/handler/count"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CalculateDebtsHandler(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		billIDStr := r.URL.Query().Get("bill_id")
		if billIDStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bill_id is required"))
			return
		}

		billID, err := strconv.ParseUint(billIDStr, 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid bill_id"))
			return
		}

		debtCalculation, err := count.CalculateDebts(r.Context(), db, uint(billID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(debtCalculation); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
