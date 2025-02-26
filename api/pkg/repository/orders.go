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
				'quantity', o.quantity,
				'shipping_status', o.status
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
				'quantity', o.quantity,
				'shipping_status', o.status
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

const (
	PENDING    types.ShippingStatus = "pending"
	DELIVERED  types.ShippingStatus = "delivered"
	IN_TRANSIT types.ShippingStatus = "in-transit"
)

func (m *orders) UpdateOrderStatus(orderID uint, status types.ShippingStatus) bool {
	query := "UPDATE orders SET status=$1 WHERE id=$2"
	_, err := m.store.Exec(query, status, orderID)
	if err != nil {
		fmt.Printf("query failed: %s", err)
		return false
	}
	return true
}

func (m *orders) GetShippingInformation() ([]*types.ShippingInformation, bool) {
	query := `SELECT 
	o.id,
	o.status,
	o.quantity,
	o.Amount,
	t.transaction_id AS transaction_id,
	pr.id AS user_id,
	pr.first_name AS first_name,
	pr.last_name AS last_name,
	pr.email AS user_email,
	p.id AS product_id,
	p.name AS product_name,
	o.created_at AS created_at
	FROM 
		orders o 
	INNER JOIN 
		products p ON p.product_id=o.product_id 
	INNER JOIN 
		transactions t ON t.id=o.transaction_id 
	INNER JOIN 
		users u ON t.customer_id=u.customer_id 
	INNER JOIN 
		profiles pr ON pr.user_id=u.id
	WHERE
		t.status='completed';
	`

	rows, err := m.store.Query(query)
	if err != nil {
		log.Printf("query failed %s", err)
		return nil, false
	}
	shippingInformattions := []*types.ShippingInformation{}
	for rows.Next() {
		shippingInformation := new(types.ShippingInformation)
		if err := rows.Scan(
			&shippingInformation.ID,
			&shippingInformation.Status,
			&shippingInformation.Quantity,
			&shippingInformation.Amount,
			&shippingInformation.TransactionID,
			&shippingInformation.User.ID,
			&shippingInformation.User.FirstName,
			&shippingInformation.User.LastName,
			&shippingInformation.User.Email,
			&shippingInformation.Product.ID,
			&shippingInformation.Product.Name,
			&shippingInformation.CreatedAt,
		); err != nil {
			log.Printf("scan failed %s", err)
			return nil, false
		}
		shippingInformattions = append(shippingInformattions, shippingInformation)
	}
	return shippingInformattions, true
}

func newOrders(store *sql.DB) types.OrdersRepository {
	return &orders{store: store}
}
