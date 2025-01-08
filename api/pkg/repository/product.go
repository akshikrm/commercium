package repository

import (
	"akshidas/e-com/pkg/types"
	"akshidas/e-com/pkg/utils"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"
)

type product struct {
	store *sql.DB
}

func (p *product) Create(product *types.NewProductRequest) (*types.OneProduct, error) {
	query := "INSERT INTO products(name, slug, price, image, description, category_id, product_id, price_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *"
	row := p.store.QueryRow(query,
		product.Name,
		product.Slug,
		product.Price,
		product.PrimaryImage,
		product.Description,
		product.CategoryID,
		product.ProductID,
		product.PriceID,
	)

	savedProduct, err := scanProductRow(row)
	if err == sql.ErrNoRows {
		return nil, utils.NotFound
	}
	if err != nil {
		log.Printf("failed to create new product %s due to %s", product.Name, err)
		return nil, utils.ServerError
	}
	return savedProduct, nil
}

func (p *product) Update(pid int, product *types.NewProductRequest) (*types.OneProduct, error) {
	query := "UPDATE products SET name=$1, slug=$2, price=$3, image=$4, description=$5, category_id=$6 WHERE id=$7 AND deleted_at IS NULL RETURNING *"
	row := p.store.QueryRow(query,
		product.Name,
		product.Slug,
		product.Price,
		product.PrimaryImage,
		product.Description,
		product.CategoryID,
		pid,
	)

	savedProduct, err := scanProductRow(row)
	if err == sql.ErrNoRows {
		return nil, utils.NotFound
	}

	if err != nil {
		log.Printf("failed to update product %s due to %s", product.Name, err)
		return nil, utils.ServerError
	}
	return savedProduct, nil
}

func (p *product) GetAll(filter url.Values) ([]*types.ProductsList, error) {
	query := buildFilterQuery("SELECT p.id, p.name, p.slug, p.price, p.image, p.description, p.created_at, c.id as c_id, c.name as c_name,c.slug as c_slug,c.description as c_description FROM products p INNER JOIN product_categories c ON p.category_id=c.id AND c.enabled='t' where p.deleted_at IS NULL", filter)

	rows, err := p.store.Query(query)

	if err == sql.ErrNoRows {
		return nil, utils.NotFound
	}

	if err != nil {
		log.Printf("failed to get all products due to %s", err)
		return nil, utils.ServerError
	}

	products := []*types.ProductsList{}
	for rows.Next() {
		product := types.ProductsList{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
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
			return nil, utils.ServerError
		}
		products = append(products, &product)
	}
	return products, nil
}

func (m *product) GetOne(id int) (*types.OneProduct, error) {
	query := "SELECT * FROM products WHERE id=$1 AND deleted_at IS NULL"
	row := m.store.QueryRow(query, id)

	product, err := scanProductRow(row)
	if err == sql.ErrNoRows {
		log.Printf("product with id %d not found due to %s", id, err)
		return nil, utils.NotFound
	}
	if err != nil {
		return nil, utils.ServerError
	}

	return product, nil
}

func (m *product) Delete(id int) error {
	query := "UPDATE products SET deleted_at=$1 WHERE id=$2"
	if _, err := m.store.Exec(query, time.Now(), id); err != nil {
		log.Printf("failed to products %d due to %s", id, err)
		return utils.ServerError
	}
	return nil
}

func (m *product) InsertImages(productID uint32, uris []string) error {
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
		return utils.ServerError
	}

	if affected, err := result.RowsAffected(); err != nil {
		return utils.ServerError
	} else {
		if affected != int64(uriCount) {
			return utils.ServerError
		}
	}

	return nil
}

func scanProductRows(rows *sql.Rows) ([]*types.OneProduct, error) {
	products := []*types.OneProduct{}
	for rows.Next() {
		product := types.OneProduct{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Slug,
			&product.Price,
			&product.Image,
			&product.Description,
			&product.CategoryID,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func scanProductRow(rows *sql.Row) (*types.OneProduct, error) {
	product := types.OneProduct{}
	err := rows.Scan(
		&product.ID,
		&product.ProductID,
		&product.Name,
		&product.Slug,
		&product.Price,
		&product.PriceID,
		&product.Image,
		&product.Description,
		&product.CategoryID,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func newProduct(store *sql.DB) *product {
	return &product{
		store: store,
	}
}
