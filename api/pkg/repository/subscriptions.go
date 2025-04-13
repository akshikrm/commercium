package repository

import "database/sql"

type subscriptions struct {
	store *sql.DB
}

func (s *subscriptions) GetSubscriptionsByUserID() {
	query := `SELECT 
		p.id,
		p.name,
		p.description,
		t.transaction_id 
	FROM 
	orders o 
	INNER JOIN 
	transactions AS t ON o.transaction_id=t.id 
	INNER JOIN products AS p ON o.product_id=p.product_id WHERE t.status='completed';`
	_, err := s.store.Query(query)

	if err != nil {

	}
}
