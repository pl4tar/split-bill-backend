package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"split-bill-backend/entity"
)

func QueryCreateNewBill(ctx context.Context, db *pgxpool.Pool, user *entity.Users, billName string) error {
	_, err := db.Exec(ctx,
		`INSERT INTO bills (title, created_by)
		VALUES($1, $2)
	`, billName, user.ID)

	return err
}

func QueryGetBillsByUserID(ctx context.Context, db *pgxpool.Pool, user *entity.Users) ([]entity.Bills, error) {
    rows, err := db.Query(
        ctx,
        `SELECT id, title, created_by
        FROM bills
        WHERE created_by = $1`,
        user.ID,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var bills []entity.Bills
    for rows.Next() {
        var bill entity.Bills
        if err := rows.Scan(
			&bill.ID, 
			&bill.Title, 
			&bill.CreatedUserID,
			); err != nil {
            return nil, err
        }
        bills = append(bills, bill)
    }

    if rows.Err() != nil {
        return nil, rows.Err()
    }

    return bills, nil
}

func QueryDeleteBillByID(ctx context.Context, db *pgxpool.Pool, bill_id uint) error{
    _, err := db.Exec(ctx,
         `DELETE FROM bills WHERE id = $1`,
         bill_id,
        )
        
    return err
}
