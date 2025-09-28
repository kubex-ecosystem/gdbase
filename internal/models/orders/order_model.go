// Package orders contains models and interfaces related to orders
package orders

import (
	"time"

	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

// PaymentMethod represents the payment methods available for an order
type PaymentMethod string

const (
	Cash       PaymentMethod = "cash"
	Credit     PaymentMethod = "credit"
	CreditCard PaymentMethod = "credit_card"
	BankSlip   PaymentMethod = "bank_slip"
	Transfer   PaymentMethod = "transfer"
)

// OrderStatus representa os status possíveis de um pedido
// (alinhar com enum do TypeScript/SQL)
type OrderStatus string

const (
	OrderStatusDraft     OrderStatus = "DRAFT"
	OrderStatusCreated   OrderStatus = "CREATED"
	OrderStatusApproved  OrderStatus = "APPROVED"
	OrderStatusShipped   OrderStatus = "SHIPPED"
	OrderStatusDelivered OrderStatus = "DELIVERED"
	OrderStatusCanceled  OrderStatus = "CANCELED"
)

// Address represents an address associated with an order
type Address struct {
	Street     string  `json:"street" xml:"street" yaml:"street" gorm:"column:street"`
	Number     string  `json:"number" xml:"number" yaml:"number" gorm:"column:number"`
	Complement *string `json:"complement,omitempty" xml:"complement,omitempty" yaml:"complement,omitempty" gorm:"column:complement"`
	District   string  `json:"district" xml:"district" yaml:"district" gorm:"column:district"`
	City       string  `json:"city" xml:"city" yaml:"city" gorm:"column:city"`
	State      string  `json:"state" xml:"state" yaml:"state" gorm:"column:state"`
	ZipCode    string  `json:"zipCode" xml:"zipCode" yaml:"zipCode" gorm:"column:zip_code"`
	IsDefault  bool    `json:"isDefault" xml:"isDefault" yaml:"isDefault" gorm:"column:is_default"`
	Type       string  `json:"type" xml:"type" yaml:"type" gorm:"column:type"`
}

// IOrderItem define a interface para manipulação de itens do pedido
type IOrderItem interface {
	TableName() string
	GetID() string
	GetProductID() string
	GetQuantity() int
	GetUnitPrice() t.Money
	GetDiscount() t.Money
	GetTotal() t.Money
	GetNotes() *string

	SetProductID(string)
	SetQuantity(int)
	SetUnitPrice(t.Money)
	SetDiscount(t.Money)
	SetTotal(t.Money)
	SetNotes(*string)
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID        string  `json:"id" xml:"id" yaml:"id" gorm:"column:id"`
	ProductID string  `json:"productId" xml:"productId" yaml:"productId" gorm:"column:product_id"`
	Quantity  int     `json:"quantity" xml:"quantity" yaml:"quantity" gorm:"column:quantity"`
	UnitPrice t.Money `json:"unitPrice" xml:"unitPrice" yaml:"unitPrice" gorm:"column:unit_price"`
	Discount  t.Money `json:"discount" xml:"discount" yaml:"discount" gorm:"column:discount"`
	Total     t.Money `json:"total" xml:"total" yaml:"total" gorm:"column:total"`
	Notes     *string `json:"notes,omitempty" xml:"notes,omitempty" yaml:"notes,omitempty" gorm:"column:notes"`
}

// Métodos de IOrderItem

func (oi *OrderItem) TableName() string      { return "order_items" }
func (oi *OrderItem) GetID() string          { return oi.ID }
func (oi *OrderItem) GetProductID() string   { return oi.ProductID }
func (oi *OrderItem) GetQuantity() int       { return oi.Quantity }
func (oi *OrderItem) GetUnitPrice() t.Money  { return oi.UnitPrice }
func (oi *OrderItem) GetDiscount() t.Money   { return oi.Discount }
func (oi *OrderItem) GetTotal() t.Money      { return oi.Total }
func (oi *OrderItem) GetNotes() *string      { return oi.Notes }
func (oi *OrderItem) SetProductID(v string)  { oi.ProductID = v }
func (oi *OrderItem) SetQuantity(v int)      { oi.Quantity = v }
func (oi *OrderItem) SetUnitPrice(v t.Money) { oi.UnitPrice = v }
func (oi *OrderItem) SetDiscount(v t.Money)  { oi.Discount = v }
func (oi *OrderItem) SetTotal(v t.Money)     { oi.Total = v }
func (oi *OrderItem) SetNotes(v *string)     { oi.Notes = v }

// IOrder define a interface para manipulação de pedidos
type IOrder interface {
	TableName() string
	GetID() string
	GetCode() *string
	GetClientID() string
	GetUserID() string
	GetItems() []IOrderItem
	GetSubtotal() t.Money
	GetDiscountValue() t.Money
	GetDiscountPercentage() float64
	GetShippingValue() t.Money
	GetTotal() t.Money
	GetStatus() OrderStatus
	GetPayments() []OrderPayment
	GetNotes() *string
	GetShippingAddress() *Address
	GetDeliveryDate() *time.Time
	GetSyncStatus() string
	GetSyncErrorMessage() *string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetSyncedAt() *time.Time

	SetCode(*string)
	SetClientID(string)
	SetUserID(string)
	SetItems([]IOrderItem)
	SetSubtotal(t.Money)
	SetDiscountValue(t.Money)
	SetDiscountPercentage(float64)
	SetShippingValue(t.Money)
	SetTotal(t.Money)
	SetStatus(OrderStatus)
	SetPayments([]OrderPayment)
	SetNotes(*string)
	SetShippingAddress(*Address)
	SetDeliveryDate(*time.Time)
	SetSyncStatus(string)
	SetSyncErrorMessage(*string)
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
	SetSyncedAt(*time.Time)
}

// OrderPayment represents a payment for an order
type OrderPayment struct {
	Method       PaymentMethod `json:"method" xml:"method" yaml:"method" gorm:"column:method"`
	Installments int           `json:"installments" xml:"installments" yaml:"installments" gorm:"column:installments"`
	DueDate      time.Time     `json:"dueDate" xml:"dueDate" yaml:"dueDate" gorm:"column:due_date"`
	Value        t.Money       `json:"value" xml:"value" yaml:"value" gorm:"column:value"`
}

// Order represents a complete order
type Order struct {
	ID                 string         `json:"id" xml:"id" yaml:"id" gorm:"column:id;primaryKey"`
	Code               *string        `json:"code,omitempty" xml:"code,omitempty" yaml:"code,omitempty" gorm:"column:code"`
	ClientID           string         `json:"clientId" xml:"clientId" yaml:"clientId" gorm:"column:client_id"`
	UserID             string         `json:"userId" xml:"userId" yaml:"userId" gorm:"column:user_id"`
	Items              []OrderItem    `json:"items" xml:"items" yaml:"items" gorm:"-"`
	Subtotal           t.Money        `json:"subtotal" xml:"subtotal" yaml:"subtotal" gorm:"column:subtotal"`
	DiscountValue      t.Money        `json:"discountValue" xml:"discountValue" yaml:"discountValue" gorm:"column:discount_value"`
	DiscountPercentage float64        `json:"discountPercentage" xml:"discountPercentage" yaml:"discountPercentage" gorm:"column:discount_percentage"`
	ShippingValue      t.Money        `json:"shippingValue" xml:"shippingValue" yaml:"shippingValue" gorm:"column:shipping_value"`
	Total              t.Money        `json:"total" xml:"total" yaml:"total" gorm:"column:total"`
	Status             OrderStatus    `json:"status" xml:"status" yaml:"status" gorm:"column:status"`
	Payments           []OrderPayment `json:"payments" xml:"payments" yaml:"payments" gorm:"-"`
	Notes              *string        `json:"notes,omitempty" xml:"notes,omitempty" yaml:"notes,omitempty" gorm:"column:notes"`
	ShippingAddress    *Address       `json:"shippingAddress,omitempty" xml:"shippingAddress,omitempty" yaml:"shippingAddress,omitempty" gorm:"-"`
	DeliveryDate       *time.Time     `json:"deliveryDate,omitempty" xml:"deliveryDate,omitempty" yaml:"deliveryDate,omitempty" gorm:"column:delivery_date"`
	SyncStatus         string         `json:"syncStatus" xml:"syncStatus" yaml:"syncStatus" gorm:"column:sync_status"`
	SyncErrorMessage   *string        `json:"syncErrorMessage,omitempty" xml:"syncErrorMessage,omitempty" yaml:"syncErrorMessage,omitempty" gorm:"column:sync_error_message"`
	CreatedAt          time.Time      `json:"createdAt" xml:"createdAt" yaml:"createdAt" gorm:"column:created_at"`
	UpdatedAt          time.Time      `json:"updatedAt" xml:"updatedAt" yaml:"updatedAt" gorm:"column:updated_at"`
	SyncedAt           *time.Time     `json:"syncedAt,omitempty" xml:"syncedAt,omitempty" yaml:"syncedAt,omitempty" gorm:"column:synced_at"`
}

// Métodos de IOrder

func (o *Order) TableName() string   { return "orders" }
func (o *Order) GetID() string       { return o.ID }
func (o *Order) GetCode() *string    { return o.Code }
func (o *Order) GetClientID() string { return o.ClientID }
func (o *Order) GetUserID() string   { return o.UserID }
func (o *Order) GetItems() []IOrderItem {
	items := make([]IOrderItem, len(o.Items))
	for i := range o.Items {
		items[i] = &o.Items[i]
	}
	return items
}
func (o *Order) GetSubtotal() t.Money           { return o.Subtotal }
func (o *Order) GetDiscountValue() t.Money      { return o.DiscountValue }
func (o *Order) GetDiscountPercentage() float64 { return o.DiscountPercentage }
func (o *Order) GetShippingValue() t.Money      { return o.ShippingValue }
func (o *Order) GetTotal() t.Money              { return o.Total }
func (o *Order) GetStatus() OrderStatus         { return o.Status }
func (o *Order) GetPayments() []OrderPayment    { return o.Payments }
func (o *Order) GetNotes() *string              { return o.Notes }
func (o *Order) GetShippingAddress() *Address   { return o.ShippingAddress }
func (o *Order) GetDeliveryDate() *time.Time    { return o.DeliveryDate }
func (o *Order) GetSyncStatus() string          { return o.SyncStatus }
func (o *Order) GetSyncErrorMessage() *string   { return o.SyncErrorMessage }
func (o *Order) GetCreatedAt() time.Time        { return o.CreatedAt }
func (o *Order) GetUpdatedAt() time.Time        { return o.UpdatedAt }
func (o *Order) GetSyncedAt() *time.Time        { return o.SyncedAt }
func (o *Order) SetCode(v *string)              { o.Code = v }
func (o *Order) SetClientID(v string)           { o.ClientID = v }
func (o *Order) SetUserID(v string)             { o.UserID = v }
func (o *Order) SetItems(v []IOrderItem) {
	items := make([]OrderItem, len(v))
	for i, item := range v {
		if oi, ok := item.(*OrderItem); ok {
			items[i] = *oi
		}
	}
	o.Items = items
}
func (o *Order) SetSubtotal(v t.Money)           { o.Subtotal = v }
func (o *Order) SetDiscountValue(v t.Money)      { o.DiscountValue = v }
func (o *Order) SetDiscountPercentage(v float64) { o.DiscountPercentage = v }
func (o *Order) SetShippingValue(v t.Money)      { o.ShippingValue = v }
func (o *Order) SetTotal(v t.Money)              { o.Total = v }
func (o *Order) SetStatus(v OrderStatus)         { o.Status = v }
func (o *Order) SetPayments(v []OrderPayment)    { o.Payments = v }
func (o *Order) SetNotes(v *string)              { o.Notes = v }
func (o *Order) SetShippingAddress(v *Address)   { o.ShippingAddress = v }
func (o *Order) SetDeliveryDate(v *time.Time)    { o.DeliveryDate = v }
func (o *Order) SetSyncStatus(v string)          { o.SyncStatus = v }
func (o *Order) SetSyncErrorMessage(v *string)   { o.SyncErrorMessage = v }
func (o *Order) SetCreatedAt(v time.Time)        { o.CreatedAt = v }
func (o *Order) SetUpdatedAt(v time.Time)        { o.UpdatedAt = v }
func (o *Order) SetSyncedAt(v *time.Time)        { o.SyncedAt = v }

// OrderDraft represents a draft version of an order
type OrderDraft struct {
	ClientID          string      `json:"clientId" xml:"clientId" yaml:"clientId" gorm:"-"`
	Items             []OrderItem `json:"items" xml:"items" yaml:"items" gorm:"-"`
	Notes             *string     `json:"notes,omitempty" xml:"notes,omitempty" yaml:"notes,omitempty" gorm:"-"`
	DeliveryDate      *time.Time  `json:"deliveryDate,omitempty" xml:"deliveryDate,omitempty" yaml:"deliveryDate,omitempty" gorm:"-"`
	PaymentMethod     *string     `json:"paymentMethod,omitempty" xml:"paymentMethod,omitempty" yaml:"paymentMethod,omitempty" gorm:"-"`
	Installments      *int        `json:"installments,omitempty" xml:"installments,omitempty" yaml:"installments,omitempty" gorm:"-"`
	ShippingAddressID *string     `json:"shippingAddressId,omitempty" xml:"shippingAddressId,omitempty" yaml:"shippingAddressId,omitempty" gorm:"-"`
}
