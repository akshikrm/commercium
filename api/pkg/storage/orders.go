package storage

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"database/sql"
	"fmt"
	"log"
)

type PurchaseStorage struct {
	store *sql.DB
}

// func (m *PurchaseStorage) GetByUserID(userID uint32) ([]*types.Purchase, error) {
// 	query := "select * from purchases where user_id=$1 and deleted_at IS NULL"
// 	rows, err := m.store.Query(query, userID)
// 	if err != nil {
// 		log.Printf("failed to retreive data due to %s", err)
// 		return nil, utils.ServerError
// 	}
// 	defer rows.Close()
//
// 	purchases := []*types.Purchase{}
// 	for rows.Next() {
// 		purchase := types.Purchase{}
// 		if err := rows.Scan(
// 			&purchase.ID,
// 			&purchase.UserID,
// 			&purchase.ProductID,
// 			&purchase.CreatedAt,
// 			&purchase.UpdatedAt,
// 			&purchase.DeletedAt,
// 		); err != nil {
// 			if err == sql.ErrNoRows {
// 				log.Printf("row cannot be read")
// 				return nil, utils.NotFound
// 			}
// 			return nil, utils.ServerError
// 		}
// 		purchases = append(purchases, &purchase)
// 	}
//
// 	return purchases, nil
//
// }

func (m *PurchaseStorage) NewOrder(orderRequest *types.OrderRequest) (uint, error) {
	query := "INSERT INTO orders(order_id, user_id, price) VALUES($1, $2, $3) RETURNING(id)"

	row := m.store.QueryRow(query, orderRequest.OrderID, orderRequest.UserID, orderRequest.Price)
	var orderID uint
	if err := row.Scan(&orderID); err != nil {
		log.Printf("failed to place order :%s", err)
		return 0, utils.ServerError
	}
	return orderID, nil
}

func (m *PurchaseStorage) NewPurchase(newPurchases []*types.PurchaseRequest) error {
	query := "INSERT INTO purchases(order_id, product_id, quantity, price) VALUES"
	for i, purchaseRequest := range newPurchases {
		query = fmt.Sprintf("%s (%d, %d, %d, %d)",
			query,
			purchaseRequest.OrderID,
			purchaseRequest.ProductID,
			purchaseRequest.Quantity,
			purchaseRequest.Price,
		)
		if i < len(newPurchases)-1 {
			query = fmt.Sprintf("%s,", query)
		}
	}
	_, err := m.store.Exec(query)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("cannot scan the saved purchase due to %s", err)
			return utils.NotFound
		}
		log.Printf("failed write purchase due to %s", err)
		return utils.ServerError
	}
	return nil

}

// func (m *PurchaseStorage) Update() {}
// func (m *PurchaseStorage) Delete() {}

func NewPurchaseStorage(store *sql.DB) *PurchaseStorage {
	return &PurchaseStorage{store: store}
}
