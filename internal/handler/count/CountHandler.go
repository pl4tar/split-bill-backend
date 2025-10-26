package count

import (
	"context"
	"fmt"
	"math"
	"split-bill-backend/internal/entity"
	"split-bill-backend/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CalculateDebts(ctx context.Context, db *pgxpool.Pool, billID uint) (*entity.DebtCalculation, error) {
	products, err := repository.QueryGetProductsIOByBillID(ctx, db, fmt.Sprintf("%d", billID))
	if err != nil {
		return nil, err
	}

	personsMap := make(map[uint]string)
	persons, err := repository.QueryGetPersonsByBillID(ctx, db, fmt.Sprintf("%d", billID))
	if err != nil {
		return nil, err
	}
	for _, person := range persons {
		personsMap[person.ID] = person.Name
	}

	debts := calculateIndividualDebts(products, personsMap)

	optimizedDebts := performMutualDeduction(debts)

	return &entity.DebtCalculation{
		BillID: billID,
		Debts:  optimizedDebts,
	}, nil
}

func calculateIndividualDebts(productsIO *entity.ProductsIO, personsMap map[uint]string) []entity.Debt {
	var debts []entity.Debt

	for _, product := range productsIO.Products {
		if len(product.Persons) == 0 {
			continue
		}

		costPerPerson := math.Round((product.Price*float64(product.Count))/float64(len(product.Persons))*100) / 100
		for _, person := range product.Persons {
			if person.ID != product.PayerID {
				debts = append(debts, entity.Debt{
					FromPersonID:   person.ID,
					FromPersonName: person.Name,
					ToPersonID:     product.PayerID,
					ToPersonName:   personsMap[product.PayerID],
					Amount:         costPerPerson,
					Description:    product.Name,
				})
			}
		}
	}

	return debts
}

func performMutualDeduction(debts []entity.Debt) []entity.Debt {
	debtMap := make(map[string]float64)

	for _, debt := range debts {
		key := fmt.Sprintf("%d-%d", debt.FromPersonID, debt.ToPersonID)
		debtMap[key] += debt.Amount
	}

	for key1, amount1 := range debtMap {
		var from1, to1 uint
		fmt.Sscanf(key1, "%d-%d", &from1, &to1)

		key2 := fmt.Sprintf("%d-%d", to1, from1)
		if amount2, exists := debtMap[key2]; exists && amount2 > 0 && amount1 > 0 {
			if amount1 > amount2 {
				debtMap[key1] = amount1 - amount2
				debtMap[key2] = 0
			} else if amount2 > amount1 {
				debtMap[key1] = 0
				debtMap[key2] = amount2 - amount1
			} else {
				debtMap[key1] = 0
				debtMap[key2] = 0
			}
		}
	}

	var result []entity.Debt
	for key, amount := range debtMap {
		if amount > 0.01 {
			var fromID, toID uint
			fmt.Sscanf(key, "%d-%d", &fromID, &toID)

			var fromName, toName, foodName string
			for _, originalDebt := range debts {
				if originalDebt.FromPersonID == fromID && originalDebt.ToPersonID == toID {
					fromName = originalDebt.FromPersonName
					toName = originalDebt.ToPersonName
					foodName = originalDebt.Description
					break
				}
			}

			result = append(result, entity.Debt{
				FromPersonID:   fromID,
				FromPersonName: fromName,
				ToPersonID:     toID,
				ToPersonName:   toName,
				Amount:         math.Round(amount*100) / 100,
				Description:    foodName,
			})
		}
	}

	return result
}
