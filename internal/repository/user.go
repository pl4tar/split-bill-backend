package repository

import (
	"context"
	"split-bill-backend/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

func QueryGetUserByEmail(ctx context.Context, db *pgxpool.Pool, email *string) (*entity.Users, error) {
	var user entity.Users

	err := db.QueryRow(ctx,
		`SELECT id, name, email FROM users
	WHERE email=$1`,
		email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func QueryGetUserByID(ctx context.Context, db *pgxpool.Pool, id *uint) (*entity.Users, error) {
	var user entity.Users

	err := db.QueryRow(ctx,
		`SELECT id, name, email FROM users
	WHERE id=$1`,
		id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func QueryUpdateUserPassword(ctx context.Context, db *pgxpool.Pool, userID int64, newHashedPassword string) error {
	_, err := db.Exec(
		ctx,
		`UPDATE users
		SET password = $1
		WHERE id = $2
	`,
		newHashedPassword,
		userID)

	return err
}

func QuerySaveNewUser(ctx context.Context, db *pgxpool.Pool, user *entity.Users) error {
	_, err := db.Exec(
		ctx,
		`INSERT INTO users (name, email, password)
		VALUES($1, $2, $3)`,
		user.Name,
		user.Email,
		user.Password,
	)

	return err
}

func QueryDeleteUser(ctx context.Context, db *pgxpool.Pool, mail string) error {
	_, err := db.Exec(ctx, "DELETE FROM users WHERE email = $1", mail)

	return err
}
