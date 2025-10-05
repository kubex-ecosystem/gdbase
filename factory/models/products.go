package models

import (
	"context"

	m "github.com/kubex-ecosystem/gdbase/internal/models/products"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
)

type ProductModel = m.Product
type ProductService = m.IProductService
type ProductRepo = m.IProductRepo

func NewProductService(productRepo ProductRepo) ProductService {
	return m.NewProductService(productRepo)
}

func NewProductRepo(ctx context.Context, dbService *svc.DBServiceImpl) ProductRepo {
	return m.NewProductRepo(ctx, dbService)
}
