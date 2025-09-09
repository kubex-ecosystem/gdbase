package models

import (
	m "github.com/kubex-ecosystem/gdbase/internal/models/products"
	"gorm.io/gorm"
)

type ProductModel = m.Product
type ProductService = m.IProductService
type ProductRepo = m.IProductRepo

func NewProductService(productRepo ProductRepo) ProductService {
	return m.NewProductService(productRepo)
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return m.NewProductRepo(db)
}
