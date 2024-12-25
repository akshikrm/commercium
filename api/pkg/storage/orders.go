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
	query := `
	 SELECT 
		t.id,
		t.transaction_id AS txn_id,
		t.status AS payment_status,
		t.invoice_number,
		t.grand_total,
		JSON_AGG(
			JSON_BUILD_OBJECT(
				'id',o.id,
				'product_id',p.id,
				'name',p.name,
				'price',p.price,
				'quantity', o.quantity
			)
		)
	AS orders,
	t.created_at
	FROM 
		transactions AS t 
	JOIN 
		orders AS o ON t.id=o.transaction_id
	JOIN 
		products as p on o.product_id=p.product_id 
	JOIN 
		users as u on t.customer_id=u.customer_id 
	WHERE 
		u.id=$1
	GROUP BY 
		t.id, t.transaction_id, t.status, t.tax, t.sub_total, t.grand_total, u.id;
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
			&order.TxnID,
			&order.PaymentStatus,
			&order.InvoiceNumber,
			&order.Total,
			&products,
			&order.CreatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				log.Println("row cannot be read")
				return nil, utils.NotFound
			}
			log.Printf("error: %s", err)
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

func (m *OrdersStorage) CreateOrder(orders []*types.NewOrder) error {
	query := "INSERT INTO orders(transaction_id, quantity, price_id, product_id, amount) VALUES"

	ordersLength := len(orders)
	for i, order := range orders {
		query = fmt.Sprintf("%s (%d, %d, '%s', '%s', '%s')", query,
			order.TransactionID,
			order.Quantity,
			order.PriceID,
			order.ProductID,
			order.Amount,
		)
		if i == ordersLength-1 {
			query = fmt.Sprintf("%s;", query)
		} else {
			query = fmt.Sprintf("%s,", query)
		}
	}
	result, err := m.store.Exec(query)
	if err != nil {
		log.Printf("failed to create orders %s", err)
		return utils.ServerError
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return utils.ServerError
	}

	if affected != int64(len(orders)) {
		return utils.ServerError
	}

	return nil

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

func (m *ProductStorage) GetOrderStatus(txnId string) string {
	query := "SELECT payment_status from transactions where transaction_id=$1"
	row := m.store.QueryRow(query, txnId)

	var transactionStatus string
	err := row.Scan(&transactionStatus)
	if err != nil {
		log.Printf("query failed %s", err)
		return ""
	}
	return transactionStatus
}

// func (m *PurchaseStorage) Update() {}
// func (m *PurchaseStorage) Delete() {}

func NewOrdersStorage(store *sql.DB) *OrdersStorage {
	return &OrdersStorage{store: store}
}
