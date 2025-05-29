package orders

import (
	"errors"
	"fmt"
)

type IOrderService interface {
	CreateOrder(order *Order) (*Order, error)
	GetOrderByID(id string) (*Order, error)
	UpdateOrder(order *Order) (*Order, error)
	DeleteOrder(id string) error
	ListOrders() ([]*Order, error)
}

type OrderService struct {
	repo IOrderRepo
}

func NewOrderService(repo IOrderRepo) IOrderService {
	return &OrderService{repo: repo}
}

func (os *OrderService) CreateOrder(order *Order) (*Order, error) {
	if order.ClientID == "" || order.UserID == "" || len(order.Items) == 0 {
		return nil, errors.New("missing required fields")
	}
	createdOrder, err := os.repo.Create(order)
	if err != nil {
		return nil, fmt.Errorf("error creating order: %w", err)
	}
	return createdOrder, nil
}

func (os *OrderService) GetOrderByID(id string) (*Order, error) {
	order, err := os.repo.FindOne("id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("error fetching order: %w", err)
	}
	return order, nil
}

func (os *OrderService) UpdateOrder(order *Order) (*Order, error) {
	updatedOrder, err := os.repo.Update(order)
	if err != nil {
		return nil, fmt.Errorf("error updating order: %w", err)
	}
	return updatedOrder, nil
}

func (os *OrderService) DeleteOrder(id string) error {
	err := os.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting order: %w", err)
	}
	return nil
}

func (os *OrderService) ListOrders() ([]*Order, error) {
	orders, err := os.repo.FindAll("status != ?", "cancelled")
	if err != nil {
		return nil, fmt.Errorf("error listing orders: %w", err)
	}
	return orders, nil
}
