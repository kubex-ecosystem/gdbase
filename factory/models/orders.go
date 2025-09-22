// Package models provides the data models and services for all models domain.
package models

import (
	m "github.com/kubex-ecosystem/gdbase/internal/models/orders"
	"gorm.io/gorm"
)

type OrdersModel = m.Order
type OrdersDraft = m.OrderDraft
type OrdersService = m.IOrderService
type OrdersRepo = m.IOrderRepo

func NewOrdersService(ordersRepo OrdersRepo) OrdersService {
	return m.NewOrderService(ordersRepo)
}

func NewOrdersRepo(db *gorm.DB) OrdersRepo {
	return m.NewOrderRepo(db)
}
