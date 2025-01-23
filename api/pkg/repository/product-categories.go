package repository

import (
	"akshidas/e-com/pkg/types"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"
)

type productCategories struct {
	store *sql.DB
}

func (p *productCategories) Create(newCategory *types.NewProductCategoryRequest) (*types.ProductCategory, bool) {
	query := "INSERT INTO product_categories(name, slug, description, enabled) VALUES($1, $2, $3, $4) RETURNING *"
	row := p.store.QueryRow(query, newCategory.Name, newCategory.Slug, newCategory.Description, newCategory.Enabled)

	savedCategory, err := scanCategoryRow(row)
	if err == sql.ErrNoRows {
		return nil, true
	}

	if err != nil {
		log.Printf("Failed to create category %s due to %s", newCategory.Name, err)
		return nil, false
	}

	return savedCategory, true

}

func (p *productCategories) GetNames() ([]*types.ProductCategoryName, bool) {
	query := "SELECT id, name, slug FROM  product_categories WHERE deleted_at IS NULL AND enabled='t'"
	rows, err := p.store.Query(query)
	if err == sql.ErrNoRows {
		return nil, true
	}
	if err != nil {
		log.Printf("failed to get product_categories due to %s", err)
		return nil, false
	}
	productsCategories := []*types.ProductCategoryName{}

	for rows.Next() {
		productCategory := &types.ProductCategoryName{}
		err := rows.Scan(
			&productCategory.ID,
			&productCategory.Name,
			&productCategory.Slug,
		)
		if err != nil {
			log.Printf("failed to scan category due to %s", err)
			return nil, false
		}
		productsCategories = append(productsCategories, productCategory)
	}
	return productsCategories, true
}

func (p *productCategories) GetAll(filter url.Values) ([]*types.ProductCategory, bool) {
	query := buildFilterQuery("SELECT * FROM product_categories as p WHERE deleted_at IS NULL", filter)
	fmt.Println(query)
	rows, err := p.store.Query(query)
	if err == sql.ErrNoRows {
		return nil, true
	}
	if err != nil {
		log.Printf("failed to get product_categories due to %s", err)
		return nil, false
	}
	productsCategories, err := scanCategoryRows(rows)
	if err != nil {
		log.Printf("failed to get all products due to %s", err)
		return nil, false
	}
	return productsCategories, true
}

func (p *productCategories) GetOne(id int) (*types.ProductCategory, bool) {
	query := "SELECT * FROM product_categories WHERE id=$1 AND deleted_at IS NULL"
	row := p.store.QueryRow(query, id)
	productCategory, err := scanCategoryRow(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, true
		}
		log.Printf("failed to get product category with %d due to %s", id, err)
		return nil, false
	}
	return productCategory, true
}

func (p *productCategories) Update(id int, updateProductCategory *types.UpdateProductCategoryRequest) (*types.ProductCategory, bool) {
	query := "UPDATE product_categories SET name=$1, slug=$2, description=$3, enabled=$4 WHERE id=$5 AND deleted_at IS NULL RETURNING *"
	row := p.store.QueryRow(
		query,
		updateProductCategory.Name,
		updateProductCategory.Slug,
		updateProductCategory.Description,
		updateProductCategory.Enabled,
		id,
	)
	productCategory, err := scanCategoryRow(row)
	if err != nil {
		log.Printf("failed to update product category with id %d due to %s", id, err)
		if err == sql.ErrNoRows {
			return nil, false
		}
		return nil, true
	}
	return productCategory, true
}

func (p *productCategories) Delete(id int) bool {
	query := "UPDATE product_categories set deleted_at=$1 where id=$2 AND deleted_at IS NULL"
	if _, err := p.store.Exec(query, time.Now(), id); err != nil {
		log.Printf("failed to delete product category with id %d due to %s", id, err)
		return false
	}
	return true
}

func scanCategoryRow(row *sql.Row) (*types.ProductCategory, error) {
	productCategory := &types.ProductCategory{}
	err := row.Scan(
		&productCategory.ID,
		&productCategory.Name,
		&productCategory.Slug,
		&productCategory.Enabled,
		&productCategory.Description,
		&productCategory.CreatedAt,
		&productCategory.UpdatedAt,
		&productCategory.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return productCategory, nil
}

func scanCategoryRows(rows *sql.Rows) ([]*types.ProductCategory, error) {
	productsCategories := []*types.ProductCategory{}

	for rows.Next() {
		productCategory := &types.ProductCategory{}
		err := rows.Scan(
			&productCategory.ID,
			&productCategory.Name,
			&productCategory.Slug,
			&productCategory.Enabled,
			&productCategory.Description,
			&productCategory.CreatedAt,
			&productCategory.UpdatedAt,
			&productCategory.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		productsCategories = append(productsCategories, productCategory)
	}
	return productsCategories, nil
}

func newProductCategory(store *sql.DB) types.ProductCategoriesRepository {
	return &productCategories{store: store}
}
