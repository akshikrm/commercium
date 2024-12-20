package storage

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

type OrdersStorage struct {
	store *sql.DB
}

func (m *OrdersStorage) GetPurchaseByOrderID(id uint) ([]*types.PurchaseList, error) {
	query := `select ps.id, ps.quantity, ps.price, ps.created_at,pr.id as
	product_id, pr.name as product_name, pr.slug as product_slug from purchases
	ps INNER JOIN products pr on ps.product_id=pr.id where ps.order_id=$1;`
	rows, err := m.store.Query(query, id)
	if err != nil {
		log.Printf("failed to retrieve data due to %s", err)
		return nil, utils.ServerError
	}
	defer rows.Close()

	purchases := []*types.PurchaseList{}
	for rows.Next() {
		purchase := types.PurchaseList{}
		if err := rows.Scan(
			&purchase.ID,
			&purchase.Price,
			&purchase.Quantity,
			&purchase.CreatedAt,
			&purchase.Product.ID,
			&purchase.Product.Name,
			&purchase.Product.Slug,
		); err != nil {
			if err == sql.ErrNoRows {
				log.Printf("row cannot be read")
				return nil, utils.NotFound
			}
			return nil, utils.ServerError
		}
		purchases = append(purchases, &purchase)
	}
	return purchases, nil
}

func (m *OrdersStorage) GetOrdersByUserID(id uint) ([]*types.OrderList, error) {
	query := `SELECT
	    o.id AS order_id,
	    o.order_id AS external_order_id,
	    o.price AS total_price,
	    o.created_at AS order_date,
	    JSON_AGG(
	        JSON_BUILD_OBJECT(
	            'id', p.id,
	            'name', p.name,
	            'slug', p.slug,
	            'price', oi.price,
	            'quantity', oi.quantity
	        )
	    ) AS products
	FROM
	    orders o
	JOIN
	    purchases oi ON o.id = oi.order_id
	JOIN
	    products p ON oi.product_id = p.id
	WHERE
	    o.user_id = $1
	GROUP BY
	    o.id;
`

	rows, err := m.store.Query(query, id)
	if err != nil {
		log.Printf("failed to retrieve orders due to %s", err)
		return nil, utils.ServerError
	}
	defer rows.Close()

	orders := []*types.OrderList{}
	for rows.Next() {
		order := types.OrderList{}
		var products string
		if err := rows.Scan(
			&order.ID,
			&order.OrderID,
			&order.Price,
			&order.CreatedAt,
			&products,
		); err != nil {
			if err == sql.ErrNoRows {
				log.Println("row cannot be read")
				return nil, utils.NotFound
			}
			return nil, utils.ServerError
		}
		err = json.Unmarshal([]byte(products), &order.Products)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON: %v", err)
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

func (m *OrdersStorage) NewOrder(orderRequest *types.OrderRequest) (uint, error) {
	query := "INSERT INTO orders(order_id, user_id, price) VALUES($1, $2, $3) RETURNING(id)"

	row := m.store.QueryRow(query, orderRequest.OrderID, orderRequest.UserID, orderRequest.Price)
	var orderID uint
	if err := row.Scan(&orderID); err != nil {
		log.Printf("failed to place order :%s", err)
		return 0, utils.ServerError
	}
	return orderID, nil
}

func (m *OrdersStorage) NewPurchase(newPurchases []*types.PurchaseRequest) error {
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

func NewOrdersStorage(store *sql.DB) *OrdersStorage {
	return &OrdersStorage{store: store}
}
