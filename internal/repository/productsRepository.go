package repository

import (
	"context"

	"split-bill-backend/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

func QueryCreateNewProduct(ctx context.Context, db *pgxpool.Pool, product *entity.ProductInput) error {
	_, err := db.Exec(ctx,
		`INSERT INTO products (name, price, count)
	VALUES($1, $2, $3)`,
		product.Name,
		product.Price,
		product.Count,
	)

	return err
}
