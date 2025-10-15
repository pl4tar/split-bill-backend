package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func QueryCreateNewBill(ctx context.Context, db *pgxpool.Pool, user *entity.Users, billName string) error {
	_, err := db.Exec(ctx,
		`INSERT INTO bills (title, created_by)
		VALUES($1, $2)
	`, billName, user.ID)

	return err
}

func QueryGetBillsByUserID(ctx context.Context, db *pgxpool.Pool, user *entity.Users) (*entity.Bills, error) {
	var bill []Bills
	
	err := db.QueryRow(
		ctx,
		`SELECT id, title 
		FROM bills
		WHERE created_by = $1`,
		user.ID,
	).Scan(
		&bill.ID,
		&bill.title
	)

	if err != nil{
		return nil, err
	}

	return &bill, err
}
