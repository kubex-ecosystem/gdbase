package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	orders "github.com/kubex-ecosystem/gdbase/internal/models/orders"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}
	db.AutoMigrate(&orders.Order{})
	return db
}

func TestOrderRepo_CRUD(t *testing.T) {
	db := setupTestDB(t)
	repo := orders.NewOrderRepo(db)

	order := &orders.Order{
		ID:        "order1",
		ClientID:  "client1",
		UserID:    "user1",
		Status:    orders.OrderStatusCreated,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// Create
	created, err := repo.Create(order)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, created.ID)

	// FindOne
	found, err := repo.FindOne("id = ?", order.ID)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, found.ID)

	// Update
	order.Status = orders.OrderStatusApproved
	updated, err := repo.Update(order)
	assert.NoError(t, err)
	assert.Equal(t, orders.OrderStatusApproved, updated.Status)

	// FindAll
	ordersList, err := repo.FindAll("user_id = ?", order.UserID)
	assert.NoError(t, err)
	assert.Len(t, ordersList, 1)

	// Delete
	err = repo.Delete(order.ID)
	assert.NoError(t, err)
}

func TestOrderService_CRUD(t *testing.T) {
	db := setupTestDB(t)
	repo := orders.NewOrderRepo(db)
	service := orders.NewOrderService(repo)

	order := &orders.Order{
		ID:        "order2",
		ClientID:  "client2",
		UserID:    "user2",
		Status:    orders.OrderStatusCreated,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// Create
	created, err := service.CreateOrder(order)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, created.ID)

	// GetOrderByID
	got, err := service.GetOrderByID(order.ID)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, got.ID)

	// UpdateOrder
	order.Status = orders.OrderStatusDelivered
	updated, err := service.UpdateOrder(order)
	assert.NoError(t, err)
	assert.Equal(t, orders.OrderStatusDelivered, updated.Status)

	// ListOrders
	list, err := service.ListOrders()
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	// DeleteOrder
	err = service.DeleteOrder(order.ID)
	assert.NoError(t, err)
}
