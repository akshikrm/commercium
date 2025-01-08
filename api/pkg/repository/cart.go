package repository

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"database/sql"
	"log"
	"time"
)

type cart struct {
	store *sql.DB
}

func (c *cart) GetAll(userID uint32) ([]*types.CartList, bool) {
	query := "SELECT c.id, c.quantity, p.price_id, p.id, p.name, p.slug, p.price, p.description, p.image, c.created_at FROM carts c INNER JOIN products p ON c.product_id=p.id WHERE c.user_id=$1 AND c.deleted_at IS NULL"
	rows, err := c.store.Query(query, userID)
	if err == sql.ErrNoRows {
		return nil, true
	}
	if err != nil {
		log.Printf("failed to get all carts due to %s", err)
		return nil, false
	}
	carts := []*types.CartList{}
	for rows.Next() {
		cart := types.CartList{}
		err := rows.Scan(
			&cart.ID,
			&cart.Quantity,
			&cart.PriceID,
			&cart.Product.ID,
			&cart.Product.Name,
			&cart.Product.Slug,
			&cart.Product.Price,
			&cart.Product.Description,
			&cart.Product.Image,
			&cart.CreatedAt,
		)
		if err != nil {
			log.Printf("failed to scan carts due to %s", err)
			return nil, false
		}
		carts = append(carts, &cart)
	}
	return carts, true
}

func (c *cart) GetAllProductIDByUserID(userID uint32) ([]*uint32, bool) {
	query := "SELECT p.id FROM carts c INNER JOIN products p ON c.product_id=p.id WHERE c.user_id=$1 AND c.deleted_at IS NULL"
	rows, err := c.store.Query(query, userID)
	if err == sql.ErrNoRows {
		return nil, true
	}
	if err != nil {
		log.Printf("failed to get all carts due to %s", err)
		return nil, false
	}
	products := []*uint32{}
	for rows.Next() {
		var pid uint32
		err := rows.Scan(&pid)
		if err != nil {
			log.Printf("failed to scan carts due to %s", err)
			return nil, false
		}
		products = append(products, &pid)
	}
	return products, true
}

func (c *cart) GetOne(cid uint32) (*types.CartList, bool) {
	query := "SELECT c.id, c.quantity, p.id, p.name, p.slug, p.price, p.description, p.image, c.created_at FROM carts c INNER JOIN products p ON c.product_id=p.id WHERE c.id=$1 AND c.deleted_at IS NULL"
	row := c.store.QueryRow(query, cid)
	cart := types.CartList{}
	err := row.Scan(
		&cart.ID,
		&cart.Quantity,
		&cart.Product.ID,
		&cart.Product.Name,
		&cart.Product.Slug,
		&cart.Product.Price,
		&cart.Product.Description,
		&cart.Product.Image,
		&cart.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, true
	}

	if err != nil {
		log.Printf("failed to scan carts due to %s", err)
		return nil, false
	}
	return &cart, true
}

func (c *cart) CheckIfEntryExist(userID, productID uint32) bool {
	query := "select exists(select 1 from carts where user_id=$1 and product_id=$2 and deleted_at IS NULL)"
	row := c.store.QueryRow(query, userID, productID)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		log.Printf("failed to scan due to %s", err)
		return false
	}
	return exists
}

func (c *cart) UpdateQuantity(updateQuantity *types.CreateCartRequest) bool {
	query := "UPDATE carts SET quantity=quantity+$1 WHERE user_id=$2 and product_id=$3"
	if _, err := c.store.Exec(query, updateQuantity.Quantity, updateQuantity.UserID, updateQuantity.ProductID); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Printf("Failed to update cart %d due to %s", updateQuantity.UserID, err)
		return false
	}
	return true
}

func (c *cart) Create(newCart *types.CreateCartRequest) (*types.Cart, bool) {
	query := "INSERT INTO carts(user_id, product_id, quantity) VALUES($1, $2, $3) RETURNING *"
	row := c.store.QueryRow(query, newCart.UserID, newCart.ProductID, newCart.Quantity)
	cart, err := scanNewCartRow(row)
	if err != nil {
		log.Printf("Failed to create new cart due to %s", err)
		return nil, false
	}
	return cart, true
}

func (c *cart) Update(cid uint32, updateCart *types.UpdateCartRequest) (*types.CartList, bool) {
	query := "UPDATE carts SET quantity=$1 WHERE id=$2 AND deleted_at IS NULL"
	if _, err := c.store.Exec(query, updateCart.Quantity, cid); err != nil {
		if err == sql.ErrNoRows {
			return nil, true
		}
		log.Printf("Failed to update cart %d due to %s", cid, err)
		return nil, false
	}
	return c.GetOne(cid)
}

func (c *cart) Delete(cid uint32) bool {
	query := "UPDATE carts set deleted_at=$1 where id=$2 AND deleted_at IS NULL"
	if _, err := c.store.Exec(query, time.Now(), cid); err != nil {
		log.Printf("failed to delete cart item with id %d due to %s", cid, err)
		return false
	}
	return true
}

func (c *cart) HardDeleteByUserID(userID string) bool {
	query := `
	DELETE FROM carts
	WHERE user_id IN (
	    SELECT u.id
	    FROM users u
	    WHERE u.customer_id = $1
	) AND deleted_at IS NULL;
	`
	if _, err := c.store.Exec(query, userID); err != nil {
		log.Printf("failed to delete cart item for customer %s due to %s", userID, err)
		return false
	}
	return true
}

func scanNewCartRow(row *sql.Row) (*types.Cart, error) {
	cart := types.Cart{}
	err := row.Scan(
		&cart.ID,
		&cart.UserID,
		&cart.ProductID,
		&cart.Quantity,
		&cart.CreatedAt,
		&cart.UpdatedAt,
		&cart.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func scanCartRows(rows *sql.Rows) ([]*types.Cart, error) {
	carts := []*types.Cart{}
	for rows.Next() {
		cart := types.Cart{}
		err := rows.Scan(
			&cart.ID,
			&cart.UserID,
			&cart.ProductID,
			&cart.Quantity,
			&cart.CreatedAt,
			&cart.UpdatedAt,
			&cart.DeletedAt,
		)
		if err != nil {
			log.Printf("failed to get all products due to %s", err)
			return nil, utils.ServerError
		}

		carts = append(carts, &cart)
	}
	return carts, nil
}

func newCart(store *sql.DB) types.CartRepository {
	return &cart{
		store: store,
	}
}
