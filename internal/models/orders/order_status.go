package orders

// OrderStatusEnum represents the possible statuses of an order
type OrderStatusEnum string

const (
	Draft      OrderStatusEnum = "draft"
	Pending    OrderStatusEnum = "pending"
	Approved   OrderStatusEnum = "approved"
	Processing OrderStatusEnum = "processing"
	Shipped    OrderStatusEnum = "shipped"
	Delivered  OrderStatusEnum = "delivered"
	Cancelled  OrderStatusEnum = "cancelled"
	Rejected   OrderStatusEnum = "rejected"
	Returned   OrderStatusEnum = "returned"
)

// OrderStatusLabels maps order statuses to user-friendly labels
var OrderStatusLabels = map[OrderStatusEnum]string{
	Draft:      "Rascunho",
	Pending:    "Pendente",
	Approved:   "Aprovado",
	Processing: "Em processamento",
	Shipped:    "Enviado",
	Delivered:  "Entregue",
	Cancelled:  "Cancelado",
	Rejected:   "Rejeitado",
	Returned:   "Devolvido",
}

// OrderStatusColors maps order statuses to their corresponding colors
var OrderStatusColors = map[OrderStatusEnum]string{
	Draft:      "gray",
	Pending:    "yellow",
	Approved:   "blue",
	Processing: "purple",
	Shipped:    "indigo",
	Delivered:  "green",
	Cancelled:  "red",
	Rejected:   "pink",
	Returned:   "orange",
}
