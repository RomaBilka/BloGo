package product

import (
	"database/sql"
)

func NewPostgreProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

type ProductRepository struct {
	db *sql.DB
}

func (r *ProductRepository) GetProductById(id ProductId) (Product, error) {
	product := Product{}

	err := r.db.QueryRow("SELECT * FROM products WHERE id=$1", id).Scan(&product.Id, &product.Name, &product.Description, &product.Image, &product.Price, &product.Count)

	if err != nil && err != sql.ErrNoRows {
		return Product{}, err
	}

	return product, nil
}

func (r *ProductRepository) UpdateProduct(id ProductId, product Product) error {
	err := r.db.QueryRow("UPDATE products SET name = $1, description = $2, image = $3, price = $4, count = $5  WHERE id=$6", product.Name, product.Description, product.Image, product.Price, product.Count, id).Err()

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r *ProductRepository) DeleteProduct(id ProductId) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id=$1", id)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r *ProductRepository) CreateProduct(product Product) (ProductId, error) {
	id := 0
	err := r.db.QueryRow("INSERT INTO products (name, description, image, price, count) VALUES ($1, $2, $3, $4, $5)  RETURNING id", product.Name, product.Description, product.Image, product.Price, product.Count).Scan(&id)
	if err != nil {
		return 0, err
	}

	return ProductId(id), nil
}

func (r *ProductRepository) GetProducts() ([]Product, error) {
	var products []Product

	rows, err := r.db.Query("SELECT * FROM Products")
	if err != nil {
		return []Product{}, err
	}
	defer rows.Close()

	for rows.Next() {
		product := Product{}
		if err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Image, &product.Price, &product.Count); err != nil && err != sql.ErrNoRows {
			return []Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}
