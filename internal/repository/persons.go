package repository

import (
	"context"
	"split-bill-backend/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

func QueryCreateNewPerson(ctx context.Context, db *pgxpool.Pool, person *entity.Persons) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var personID uint
	err = tx.QueryRow(ctx,
		`INSERT INTO peoples (name)
		VALUES($1)
		RETURNING id`,
		person.Name,
	).Scan(&personID)

	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		`INSERT INTO bill_people (bill_id, person_id)
		VALUES($1, $2)`,
		person.BillID,
		personID,
	)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func QueryGetPersonsByBillID(ctx context.Context, db *pgxpool.Pool, billID string) ([]entity.PersonsOutput, error) {
	rows, err := db.Query(
		ctx,
		`SELECT p.id, p.name
		FROM peoples p
		INNER JOIN bill_people bp ON p.id = bp.person_id
		WHERE bp.bill_id = $1
		ORDER BY p.id`,
		billID,
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

func QueryGetPersonsForProduct(ctx context.Context, db *pgxpool.Pool, productID uint) ([]entity.Persons, error) {
	rows, err := db.Query(ctx,
		`SELECT p.id, p.name
        FROM peoples p
        INNER JOIN product_assignments pa ON p.id = pa.person_id
        WHERE pa.product_id = $1
        ORDER BY p.id`,
		productID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []entity.Persons
	for rows.Next() {
		var person entity.Persons
		if err := rows.Scan(
			&person.ID,
			&person.Name,
		); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}

func QueryDeletePersonByID(ctx context.Context, db *pgxpool.Pool, personID *uint) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		`DELETE FROM bill_people WHERE person_id = $1`,
		personID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		`DELETE FROM product_assignments WHERE person_id = $1`,
		personID,
	)
	if err != nil {
		return err
	}
	//TODO сделать при реализации debts
	// _, err = tx.Exec(ctx,
	// 	`DELETE FROM debts WHERE from_person_id = $1 OR to_person_id = $1`,
	// 	personID,
	// )
	// if err != nil {
	// 	return err
	// }

	_, err = tx.Exec(ctx,
		`DELETE FROM peoples WHERE id = $1`,
		personID,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func QueryDeleteAllPersonsByBillID(ctx context.Context, db *pgxpool.Pool, billID uint) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		`DELETE FROM product_assignments 
		WHERE person_id IN (SELECT person_id FROM bill_people WHERE bill_id = $1)`,
		billID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		`DELETE FROM debts 
		WHERE from_person_id IN (SELECT person_id FROM bill_people WHERE bill_id = $1)
		OR to_person_id IN (SELECT person_id FROM bill_people WHERE bill_id = $1)`,
		billID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		`DELETE FROM bill_people WHERE bill_id = $1`,
		billID,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
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
