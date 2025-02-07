package repository

import (
	"akshidas/e-com/pkg/types"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/lib/pq"
)

type product struct {
	store *sql.DB
}

func (p *product) GetAll(filter url.Values) ([]*types.ProductsList, bool) {
	query := buildFilterQuery("SELECT p.id, p.name, p.type, p.slug, p.price, p.image[1], p.description, p.created_at, c.id as c_id, c.name as c_name,c.slug as c_slug,c.description as c_description FROM products p INNER JOIN product_categories c ON p.category_id=c.id AND c.enabled='t' where p.deleted_at IS NULL", filter)
	rows, err := p.store.Query(query)
	if err == sql.ErrNoRows {
		return nil, true
	}

	if err != nil {
		log.Printf("failed to get all products due to %s", err)
		return nil, false
	}

	products := []*types.ProductsList{}
	for rows.Next() {
		product := types.ProductsList{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Type,
			&product.Slug,
			&product.Price,
			&product.Image,
			&product.Description,
			&product.CreatedAt,
			&product.Category.ID,
			&product.Category.Name,
			&product.Category.Slug,
			&product.Category.Description,
		)
		if err != nil {
			log.Printf("failed to scan products due to %s", err)
			return nil, false
		}
		products = append(products, &product)
	}

	return products, true
}

func (m *product) GetOne(id int) (*types.OneProduct, bool) {
	query := `SELECT 
		id,
		category_id,
		name,
		slug,
		price,
		image,
		description,
		created_at,
		updated_at,
		deleted_at
	FROM 
		products 
	WHERE 
		id=$1 
	AND 
		deleted_at IS NULL`
	row := m.store.QueryRow(query, id)

	product := types.OneProduct{}
	err := row.Scan(
		&product.ID,
		&product.CategoryID,
		&product.Name,
		&product.Slug,
		&product.Price,
		pq.Array(&product.Image),
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.DeletedAt,
	)
	if err == sql.ErrNoRows {
		return nil, true
	}
	if err != nil {
		log.Printf("failed to get product with id %d due to %s", id, err)
		return nil, false
	}

	return &product, true
}

func (p *product) CreatePrice(createPrice *types.NewPrice) bool {
	query := `
	INSERT INTO 
		prices (price, label, price_id, product_id)
	VALUES
		($1, $2, $3, $4)`
	_, err := p.store.Exec(query,
		createPrice.Amount,
		createPrice.Label,
		createPrice.ID,
		createPrice.ProductID,
	)
	if err != nil {
		log.Printf("failed to add price to database due to %s", err)
		return false
	}
	return true
}

func (p *product) Create(product *types.NewProductRequest) (*types.OneProduct, bool) {
	query := `INSERT INTO products
		(name, slug,  image, description, category_id, product_id) 
	VALUES 
		($1, $2, $3, $4, $5, $6) 
	RETURNING 
		id, product_id, name, slug, image, description, category_id`

	row := p.store.QueryRow(query,
		product.Name,
		product.Slug,
		pq.Array(product.Image),
		product.Description,
		product.CategoryID,
		product.ProductID,
	)

	savedProduct := types.OneProduct{}
	err := row.Scan(
		&savedProduct.ID,
		&savedProduct.ProductID,
		&savedProduct.Name,
		&savedProduct.Slug,
		pq.Array(&savedProduct.Image),
		&savedProduct.Description,
		&savedProduct.CategoryID,
	)
	if err == sql.ErrNoRows {
		return nil, true
	}
	if err != nil {
		log.Printf("failed to create new product %s due to %s", product.Name, err)
		return nil, false
	}
	return &savedProduct, true
}

func (p *product) Update(pid int, product *types.NewProductRequest) (*types.OneProduct, bool) {
	query := `UPDATE 
		products 
	SET 
		name=$1,
		slug=$2,
		price=$3,
		image=$4,
		description=$5,
		category_id=$6 
	WHERE 
		id=$7 
	AND 
		deleted_at IS NULL 
	RETURNING 
			id, product_id, name, slug, price, price_id, image, description, category_id
	`
	row := p.store.QueryRow(query,
		product.Name,
		product.Slug,
		product.Price,
		pq.Array(product.Image),
		product.Description,
		product.CategoryID,
		pid,
	)

	savedProduct := types.OneProduct{}
	err := row.Scan(
		&savedProduct.ID,
		&savedProduct.ProductID,
		&savedProduct.Name,
		&savedProduct.Slug,
		&savedProduct.Price,
		&savedProduct.PriceID,
		pq.Array(&savedProduct.Image),
		&savedProduct.Description,
		&savedProduct.CategoryID,
	)
	if err == sql.ErrNoRows {
		return nil, true
	}

	if err != nil {
		log.Printf("failed to update product %s due to %s", product.Name, err)
		return nil, false
	}
	return &savedProduct, true
}

func (m *product) Delete(id int) bool {
	query := "UPDATE products SET deleted_at=$1 WHERE id=$2"
	if _, err := m.store.Exec(query, time.Now(), id); err != nil {
		log.Printf("failed to products %d due to %s", id, err)
		return false
	}
	return true
}

func (m *product) InsertImages(productID uint32, uris []string) bool {
	query := "INSERT INTO product_images(product_id, uri) VALUES"
	uriCount := len(uris)
	for i, uri := range uris {
		query = fmt.Sprintf("%s (%d, '%s')", query, productID, uri)
		if i == uriCount-1 {
			query = fmt.Sprintf("%s;", query)
		} else {
			query = fmt.Sprintf("%s,", query)
		}
	}
	fmt.Println(query)
	result, err := m.store.Exec(query)

	if err != nil {
		log.Printf("failed to insert images due to: %s", err)
		return false
	}

	if affected, err := result.RowsAffected(); err != nil {
		return false
	} else {
		if affected != int64(uriCount) {
			return false
		}
	}

	return true
}

func newProduct(store *sql.DB) *product {
	return &product{
		store: store,
	}
}
