package repository

import (
	"context"
	"strconv"

	"split-bill-backend/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

func QueryCreateNewProduct(ctx context.Context, db *pgxpool.Pool, productIO *entity.ProductsIO) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, product := range productIO.Products {
		var productID uint
		err = tx.QueryRow(ctx,
			`INSERT INTO products (name, price, count, payer_id)
			VALUES($1, $2, $3, $4)
			RETURNING id`,
			product.Name,
			product.Price,
			product.Count,
			product.PayerID,
		).Scan(
			&productID,
		)

		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx,
			`INSERT INTO bill_products (bill_id, product_id)
			VALUES($1, $2)`,
			productIO.BillID,
			productID,
		)
		if err != nil {
			return err
		}

		for _, person := range product.Persons {
			_, err = tx.Exec(ctx,
				`INSERT INTO product_assignments (product_id, person_id)
				VALUES($1, $2)`,
				productID,
				person.ID,
			)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

func QueryGetProductsIOByBillID(ctx context.Context, db *pgxpool.Pool, billID *string) (*entity.ProductsIO, error) {
	rows, err := db.Query(ctx,
		`SELECT 
            p.id as product_id, 
            p.name as product_name, 
            p.price, 
            p.count, 
            p.payer_id,
            per.id as person_id,
            per.name as person_name
         FROM products p
         INNER JOIN bill_products bp ON p.id = bp.product_id
         LEFT JOIN product_assignments pa ON p.id = pa.product_id
         LEFT JOIN peoples per ON pa.person_id = per.id
         WHERE bp.bill_id = $1
         ORDER BY p.id, per.id`,
		billID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productsIO entity.ProductsIO
	billIDUint, _ := strconv.ParseUint(*billID, 10, 32)
	productsIO.BillID = uint(billIDUint)

	productsMap := make(map[uint]*entity.Product)

	for rows.Next() {
		var (
			productID   uint
			productName string
			price       float64
			count       uint
			payerID     uint
			personID    *uint
			personName  *string
		)

		if err := rows.Scan(
			&productID,
			&productName,
			&price,
			&count,
			&payerID,
			&personID,
			&personName,
		); err != nil {
			return nil, err
		}

		// Добавляем продукт если его еще нет в мапе
		if _, exists := productsMap[productID]; !exists {
			productsMap[productID] = &entity.Product{
				ID:      productID,
				Name:    productName,
				Price:   price,
				Count:   count,
				PayerID: payerID,
				Persons: []entity.Persons{},
			}
		}

		// Добавляем person если он есть
		if personID != nil && personName != nil {
			person := entity.Persons{
				ID:   *personID,
				Name: *personName,
			}
			productsMap[productID].Persons = append(productsMap[productID].Persons, person)
		}
	}

	// Конвертируем мапу в слайс
	for _, product := range productsMap {
		productsIO.Products = append(productsIO.Products, *product)
	}

	return &productsIO, nil
}

func QueryGetProductsByBillID(ctx context.Context, db *pgxpool.Pool, billID *string) ([]entity.Product, error) {
	rows, err := db.Query(ctx,
		`SELECT p.id, p.name, p.price, p.count, p.payer_id
		FROM products p
		INNER JOIN bill_products bp ON p.id = bp.product_id
		WHERE bp.bill_id = $1`,
		billID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Count,
			&product.PayerID,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func QueryGetProductPeople(ctx context.Context, db *pgxpool.Pool, productID uint) (*[]entity.Persons, error) {
	rows, err := db.Query(ctx,
		`SELECT p.id, p.name
		FROM product_assignments pa
		RiGHT JOIN persons p ON p.id == pa.person_id
		WHERE product_id = $1`,
		productID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []entity.Persons
	for rows.Next() {
		var person entity.Persons
		if err := rows.Scan(&person); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}

	return &persons, nil
}

func QueryDeleteProductByID(ctx context.Context, db *pgxpool.Pool, productID uint) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		`DELETE FROM product_assignments WHERE product_id = $1`,
		productID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		`DELETE FROM bill_products WHERE product_id = $1`,
		productID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		`DELETE FROM products WHERE id = $1`,
		productID,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
