package services

import (
	"fmt"

	"github.com/RomaBilka/BloGo/example-wire/internal/models"
)

type productRepository interface {
	CreateProduct(product models.Product) (models.ProductId, error)
	GetProductById(id models.ProductId) (models.Product, error)
	UpdateProduct(id models.ProductId, Product models.Product) error
	DeleteProduct(id models.ProductId) error
	GetProducts() ([]models.Product, error)
}

type ProductService struct {
	repository productRepository
}

func NewProductService(repository productRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) CreateProduct(product models.Product) (models.Product, error) {
	id, err := s.repository.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	product, err = s.repository.GetProductById(id)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *ProductService) GetProduct(id models.ProductId) (models.Product, error) {
	product, err := s.repository.GetProductById(id)
	if err != nil {
		return models.Product{}, err
	}

	if product.Id == 0 {
		return models.Product{}, fmt.Errorf("%s", "Bad request, Product not found")
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(id models.ProductId) error {
	err := s.repository.DeleteProduct(id)

	return err
}

func (s *ProductService) UpdateProduct(id models.ProductId, product models.Product) (models.Product, error) {
	exists, err := s.repository.GetProductById(id)
	if err != nil {
		return models.Product{}, err
	}
	if exists.Id == 0 {
		return models.Product{}, fmt.Errorf("%s", "Bad request, Product not found")
	}

	err = s.repository.UpdateProduct(id, product)
	if err != nil {
		return models.Product{}, err
	}

	product, err = s.repository.GetProductById(id)
	if err != nil {
		return models.Product{}, err
	}

	return product, err
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	product, err := s.repository.GetProducts()
	if err != nil {
		return []models.Product{}, err
	}

	return product, nil
}
