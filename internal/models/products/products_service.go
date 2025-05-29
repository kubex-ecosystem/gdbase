package products

import (
	"errors"
	"fmt"
)

type IProductService interface {
	CreateProduct(product *Product) (*Product, error)
	GetProductByID(id string) (*Product, error)
	UpdateProduct(product *Product) (*Product, error)
	DeleteProduct(id string) error
	ListProducts() ([]*Product, error)
}

type ProductService struct {
	repo IProductRepo
}

func NewProductService(repo IProductRepo) IProductService {
	return &ProductService{repo: repo}
}

func (ps *ProductService) CreateProduct(product *Product) (*Product, error) {
	if product.Name == "" || product.Code == "" || product.SKU == "" {
		return nil, errors.New("missing required fields")
	}
	createdProduct, err := ps.repo.Create(product)
	if err != nil {
		return nil, fmt.Errorf("error creating product: %w", err)
	}
	return createdProduct, nil
}

func (ps *ProductService) GetProductByID(id string) (*Product, error) {
	product, err := ps.repo.FindOne("id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("error fetching product: %w", err)
	}
	return product, nil
}

func (ps *ProductService) UpdateProduct(product *Product) (*Product, error) {
	updatedProduct, err := ps.repo.Update(product)
	if err != nil {
		return nil, fmt.Errorf("error updating product: %w", err)
	}
	return updatedProduct, nil
}

func (ps *ProductService) DeleteProduct(id string) error {
	err := ps.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}
	return nil
}

func (ps *ProductService) ListProducts() ([]*Product, error) {
	products, err := ps.repo.FindAll("is_active = ?", true)
	if err != nil {
		return nil, fmt.Errorf("error listing products: %w", err)
	}
	return products, nil
}
