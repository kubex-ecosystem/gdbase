package orders

import (
	"fmt"

	"gorm.io/gorm"
)

// Adiciona importação explícita do model Order

type IOrderRepo interface {
	Create(o *Order) (*Order, error)
	FindOne(where ...interface{}) (*Order, error)
	FindAll(where ...interface{}) ([]*Order, error)
	Update(o *Order) (*Order, error)
	Delete(id string) error
	Close() error
}

type OrderRepo struct {
	g *gorm.DB
}

func NewOrderRepo(db *gorm.DB) IOrderRepo {
	if db == nil {
		return nil
	}
	return &OrderRepo{db}
}

func (or *OrderRepo) Create(o *Order) (*Order, error) {
	if o == nil {
		return nil, fmt.Errorf("OrderRepo: Order is nil")
	}
	err := or.g.Create(o).Error
	if err != nil {
		return nil, fmt.Errorf("OrderRepo: failed to create Order: %w", err)
	}
	return o, nil
}

func (or *OrderRepo) FindOne(where ...interface{}) (*Order, error) {
	var o Order
	err := or.g.Where(where[0], where[1:]...).First(&o).Error
	if err != nil {
		return nil, fmt.Errorf("OrderRepo: failed to find Order: %w", err)
	}
	return &o, nil
}

func (or *OrderRepo) FindAll(where ...interface{}) ([]*Order, error) {
	var os []*Order
	err := or.g.Where(where[0], where[1:]...).Find(&os).Error
	if err != nil {
		return nil, fmt.Errorf("OrderRepo: failed to find all orders: %w", err)
	}
	return os, nil
}

func (or *OrderRepo) Update(o *Order) (*Order, error) {
	if o == nil {
		return nil, fmt.Errorf("OrderRepo: Order is nil")
	}
	err := or.g.Save(o).Error
	if err != nil {
		return nil, fmt.Errorf("OrderRepo: failed to update Order: %w", err)
	}
	return o, nil
}

func (or *OrderRepo) Delete(id string) error {
	err := or.g.Delete(&Order{}, id).Error
	if err != nil {
		return fmt.Errorf("OrderRepo: failed to delete Order: %w", err)
	}
	return nil
}

func (or *OrderRepo) Close() error {
	sqlDB, err := or.g.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
