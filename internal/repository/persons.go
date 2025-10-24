package repository

import (
	"context"
	"split-bill-backend/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

func QueryCreateNewPerson(ctx context.Context, db *pgxpool.Pool, Person *entity.Persons) error {
	_, err := db.Exec(ctx,
		`INSERT INTO peoples (name, owner_id)
		VALUES($1, $2)
	`,
		Person.Name,
		Person.OwnerID,
	)

	return err
}

func QueryGetPersonsByUserID(ctx context.Context, db *pgxpool.Pool, id *string) ([]entity.PersonsOutput, error) {
	rows, err := db.Query(
		ctx,
		`SELECT id, name
        FROM peoples
        WHERE owner_id = $1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []entity.PersonsOutput
	for rows.Next() {
		var person entity.PersonsOutput
		if err := rows.Scan(
			&person.ID,
			&person.Name,
		); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return persons, nil
}

func QueryDeletePersonByID(ctx context.Context, db *pgxpool.Pool, person_id uint) error {
	_, err := db.Exec(ctx,
		`DELETE FROM peoples WHERE id = $1`,
		person_id,
	)

	return err
}

func QueryDeletepersonByUserID(ctx context.Context, db *pgxpool.Pool, person_id uint) error {
	_, err := db.Exec(ctx,
		`DELETE FROM peoples WHERE created_by = $1`,
		person_id,
	)

	return err
}

func QueryEditName(ctx context.Context, db *pgxpool.Pool, person *entity.Persons) error {
	_, err := db.Exec(ctx,
		`UPDATE peoples
		SET name = $1
		WHERE id = $2`,
		person.Name,
		person.ID,
	)

	return err
}
