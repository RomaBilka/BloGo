package product

import (
	"fmt"
)

type productRepository interface {
	CreateProduct(product Product) (ProductId, error)
	GetProductById(id ProductId) (Product, error)
	UpdateProduct(id ProductId, Product Product) error
	DeleteProduct(id ProductId) error
	GetProducts() ([]Product, error)
}

type ProductService struct {
	repository productRepository
}

func NewProductService(repository productRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) CreateProduct(product Product) (Product, error) {
	id, err := s.repository.CreateProduct(product)
	if err != nil {
		return Product{}, err
	}

	product, err = s.repository.GetProductById(id)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *ProductService) GetProduct(id ProductId) (Product, error) {
	product, err := s.repository.GetProductById(id)
	if err != nil {
		return Product{}, err
	}

	if product.Id == 0 {
		return Product{}, fmt.Errorf("%s", "Bad request, Product not found")
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(id ProductId) error {
	err := s.repository.DeleteProduct(id)

	return err
}

func (s *ProductService) UpdateProduct(id ProductId, product Product) (Product, error) {
	exists, err := s.repository.GetProductById(id)
	if err != nil {
		return Product{}, err
	}
	if exists.Id == 0 {
		return Product{}, fmt.Errorf("%s", "Bad request, Product not found")
	}

	err = s.repository.UpdateProduct(id, product)
	if err != nil {
		return Product{}, err
	}

	product, err = s.repository.GetProductById(id)
	if err != nil {
		return Product{}, err
	}

	return product, err
}

func (s *ProductService) GetProducts() ([]Product, error) {
	products, err := s.repository.GetProducts()
	if err != nil {
		return []Product{}, err
	}

	return products, nil
}
