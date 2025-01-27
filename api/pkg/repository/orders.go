package repository

import (
	"akshidas/e-com/pkg/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

type orders struct {
	store *sql.DB
}

func (m *orders) GetPurchaseByOrderID(id uint) ([]*types.PurchaseList, bool) {
	query := `select ps.id, ps.quantity, ps.price, ps.created_at,pr.id as
	product_id, pr.name as product_name, pr.slug as product_slug from purchases
	ps INNER JOIN products pr on ps.product_id=pr.id where ps.order_id=$1;`
	rows, err := m.store.Query(query, id)
	if err != nil {
		log.Printf("failed to retrieve data due to %s", err)
		return nil, false
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
			log.Printf("failed to read row due to %s", err)
			return nil, false
		}
		purchases = append(purchases, &purchase)
	}
	return purchases, true
}

func (m *orders) getAllOrders(query string) ([]*types.OrderList, bool) {
	rows, err := m.store.Query(query)
	if err != nil {
		log.Printf("failed to retrieve orders due to %s", err)
		return nil, false
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
				return nil, true
			}
			log.Printf("error: %s", err)
			return nil, false
		}
		err = json.Unmarshal([]byte(products), &order.Products)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON: %v", err)
		}
		orders = append(orders, &order)
	}
	return orders, true

}

func (m *orders) GetOrdersByUserID(id uint32) ([]*types.OrderList, bool) {
	query := fmt.Sprintf(`
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
		u.id=%d
	GROUP BY 
		t.id, t.transaction_id, t.status, t.tax, t.sub_total, t.grand_total, u.id;
`, id)
	return m.getAllOrders(query)
}

func (m *orders) GetAllOrders() ([]*types.OrderList, bool) {
	query := fmt.Sprintf(`
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
	GROUP BY 
		t.id, t.transaction_id, t.status, t.tax, t.sub_total, t.grand_total, u.id;
`)
	fmt.Println(query)
	return m.getAllOrders(query)
}

func (m *orders) CreateOrder(orders []*types.NewOrder) bool {
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
		return false
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false
	}

	if affected != int64(len(orders)) {
		return false
	}

	return true

}

func (m *orders) NewOrder(orderRequest *types.OrderRequest) (uint, bool) {
	query := "INSERT INTO orders(order_id, user_id, price) VALUES($1, $2, $3) RETURNING(id)"

	row := m.store.QueryRow(query, orderRequest.OrderID, orderRequest.UserID, orderRequest.Price)
	var orderID uint
	if err := row.Scan(&orderID); err != nil {
		log.Printf("failed to place order :%s", err)
		return 0, false
	}
	return orderID, true
}

func (m *orders) NewPurchase(newPurchases []*types.PurchaseRequest) bool {
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
			return true
		}
		log.Printf("failed write purchase due to %s", err)
		return false
	}
	return true

}

func (m *product) GetOrderStatus(txnId string) string {
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

func newOrders(store *sql.DB) types.OrdersRepository {
	return &orders{store: store}
}
