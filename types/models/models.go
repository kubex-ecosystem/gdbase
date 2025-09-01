// Package models contains the data models for the application.
package models

import "time"

type AuditEvents struct {
	EntityType string    `json:"Entity_type" yaml:"Entity_type" xml:"Entity_type"`
	UserID     string    `json:"User_id" yaml:"User_id" xml:"User_id"`
	Changes    any       `json:"Changes" yaml:"Changes" xml:"Changes"`
	Action     string    `json:"Action" yaml:"Action" xml:"Action"`
	ID         any       `json:"Id" yaml:"Id" xml:"Id"`
	EntityID   any       `json:"Entity_id" yaml:"Entity_id" xml:"Entity_id"`
	CreatedAt  time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
}

type Customers struct {
	City             string    `json:"City" yaml:"City" xml:"City"`
	Phone            string    `json:"Phone" yaml:"Phone" xml:"Phone"`
	Code             string    `json:"Code" yaml:"Code" xml:"Code"`
	CreatedAt        time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	ID               any       `json:"Id" yaml:"Id" xml:"Id"`
	LastSyncAt       time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	Name             string    `json:"Name" yaml:"Name" xml:"Name"`
	Email            string    `json:"Email" yaml:"Email" xml:"Email"`
	Address          string    `json:"Address" yaml:"Address" xml:"Address"`
	LastPurchaseDate time.Time `json:"Last_purchase_date" yaml:"Last_purchase_date" xml:"Last_purchase_date"`
	IsActive         bool      `json:"Is_active" yaml:"Is_active" xml:"Is_active"`
	Document         string    `json:"Document" yaml:"Document" xml:"Document"`
	UpdatedAt        time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	PaymentTerms     string    `json:"Payment_terms" yaml:"Payment_terms" xml:"Payment_terms"`
	State            string    `json:"State" yaml:"State" xml:"State"`
	PostalCode       string    `json:"Postal_code" yaml:"Postal_code" xml:"Postal_code"`
	ExternalID       string    `json:"External_id" yaml:"External_id" xml:"External_id"`
	Country          string    `json:"Country" yaml:"Country" xml:"Country"`
	CreditLimit      float64   `json:"Credit_limit" yaml:"Credit_limit" xml:"Credit_limit"`
}

type Inventory struct {
	ExpirationDate time.Time `json:"Expiration_date" yaml:"Expiration_date" xml:"Expiration_date"`
	LotControl     string    `json:"Lot_control" yaml:"Lot_control" xml:"Lot_control"`
	LocationCode   string    `json:"Location_code" yaml:"Location_code" xml:"Location_code"`
	ProductID      any       `json:"Product_id" yaml:"Product_id" xml:"Product_id"`
	LastSyncAt     time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	CreatedAt      time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	MinimumLevel   float64   `json:"Minimum_level" yaml:"Minimum_level" xml:"Minimum_level"`
	LastCountDate  time.Time `json:"Last_count_date" yaml:"Last_count_date" xml:"Last_count_date"`
	MaximumLevel   float64   `json:"Maximum_level" yaml:"Maximum_level" xml:"Maximum_level"`
	IsActive       bool      `json:"Is_active" yaml:"Is_active" xml:"Is_active"`
	Quantity       float64   `json:"Quantity" yaml:"Quantity" xml:"Quantity"`
	ID             any       `json:"Id" yaml:"Id" xml:"Id"`
	WarehouseID    any       `json:"Warehouse_id" yaml:"Warehouse_id" xml:"Warehouse_id"`
	UpdatedAt      time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	Status         any       `json:"Status" yaml:"Status" xml:"Status"`
	ReorderPoint   float64   `json:"Reorder_point" yaml:"Reorder_point" xml:"Reorder_point"`
}

type InventoryMovements struct {
	Quantity          float64   `json:"Quantity" yaml:"Quantity" xml:"Quantity"`
	MovementType      string    `json:"Movement_type" yaml:"Movement_type" xml:"Movement_type"`
	PreviousQuantity  float64   `json:"Previous_quantity" yaml:"Previous_quantity" xml:"Previous_quantity"`
	ReferenceDocument string    `json:"Reference_document" yaml:"Reference_document" xml:"Reference_document"`
	InventoryID       any       `json:"Inventory_id" yaml:"Inventory_id" xml:"Inventory_id"`
	ID                any       `json:"Id" yaml:"Id" xml:"Id"`
	ProductID         any       `json:"Product_id" yaml:"Product_id" xml:"Product_id"`
	ExternalID        string    `json:"External_id" yaml:"External_id" xml:"External_id"`
	CreatedBy         string    `json:"Created_by" yaml:"Created_by" xml:"Created_by"`
	Reason            string    `json:"Reason" yaml:"Reason" xml:"Reason"`
	CreatedAt         time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	WarehouseID       any       `json:"Warehouse_id" yaml:"Warehouse_id" xml:"Warehouse_id"`
	LastSyncAt        time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
}

type OrderItems struct {
	SuggestionReason string    `json:"Suggestion_reason" yaml:"Suggestion_reason" xml:"Suggestion_reason"`
	Quantity         float64   `json:"Quantity" yaml:"Quantity" xml:"Quantity"`
	LastSyncAt       time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	IsSuggested      bool      `json:"Is_suggested" yaml:"Is_suggested" xml:"Is_suggested"`
	UpdatedAt        time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	ExternalID       string    `json:"External_id" yaml:"External_id" xml:"External_id"`
	OrderID          any       `json:"Order_id" yaml:"Order_id" xml:"Order_id"`
	CreatedAt        time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Discount         float64   `json:"Discount" yaml:"Discount" xml:"Discount"`
	ProductID        any       `json:"Product_id" yaml:"Product_id" xml:"Product_id"`
	ID               any       `json:"Id" yaml:"Id" xml:"Id"`
	UnitPrice        float64   `json:"Unit_price" yaml:"Unit_price" xml:"Unit_price"`
	Total            float64   `json:"Total" yaml:"Total" xml:"Total"`
}

type Orders struct {
	PredictionID             any       `json:"Prediction_id" yaml:"Prediction_id" xml:"Prediction_id"`
	ID                       any       `json:"Id" yaml:"Id" xml:"Id"`
	ShippingAmount           float64   `json:"Shipping_amount" yaml:"Shipping_amount" xml:"Shipping_amount"`
	Priority                 int       `json:"Priority" yaml:"Priority" xml:"Priority"`
	CustomerID               any       `json:"Customer_id" yaml:"Customer_id" xml:"Customer_id"`
	ShippingAddress          string    `json:"Shipping_address" yaml:"Shipping_address" xml:"Shipping_address"`
	EstimatedDeliveryDate    time.Time `json:"Estimated_delivery_date" yaml:"Estimated_delivery_date" xml:"Estimated_delivery_date"`
	OrderNumber              string    `json:"Order_number" yaml:"Order_number" xml:"Order_number"`
	DiscountAmount           float64   `json:"Discount_amount" yaml:"Discount_amount" xml:"Discount_amount"`
	ActualDeliveryDate       time.Time `json:"Actual_delivery_date" yaml:"Actual_delivery_date" xml:"Actual_delivery_date"`
	OrderDate                time.Time `json:"Order_date" yaml:"Order_date" xml:"Order_date"`
	TotalAmount              float64   `json:"Total_amount" yaml:"Total_amount" xml:"Total_amount"`
	Status                   any       `json:"Status" yaml:"Status" xml:"Status"`
	LastSyncAt               time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	TaxAmount                float64   `json:"Tax_amount" yaml:"Tax_amount" xml:"Tax_amount"`
	ExpectedMargin           float64   `json:"Expected_margin" yaml:"Expected_margin" xml:"Expected_margin"`
	CreatedAt                time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	UpdatedAt                time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	IsAutomaticallyGenerated bool      `json:"Is_automatically_generated" yaml:"Is_automatically_generated" xml:"Is_automatically_generated"`
	PaymentMethod            string    `json:"Payment_method" yaml:"Payment_method" xml:"Payment_method"`
	Notes                    string    `json:"Notes" yaml:"Notes" xml:"Notes"`
	ExternalID               string    `json:"External_id" yaml:"External_id" xml:"External_id"`
	PaymentStatus            any       `json:"Payment_status" yaml:"Payment_status" xml:"Payment_status"`
	FinalAmount              float64   `json:"Final_amount" yaml:"Final_amount" xml:"Final_amount"`
}

type PredictionDailyData struct {
	PredictionID    any       `json:"Prediction_id" yaml:"Prediction_id" xml:"Prediction_id"`
	PredictedStock  float64   `json:"Predicted_stock" yaml:"Predicted_stock" xml:"Predicted_stock"`
	CreatedAt       time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	LowerBound      float64   `json:"Lower_bound" yaml:"Lower_bound" xml:"Lower_bound"`
	PredictedDemand float64   `json:"Predicted_demand" yaml:"Predicted_demand" xml:"Predicted_demand"`
	ID              any       `json:"Id" yaml:"Id" xml:"Id"`
	DayDate         any       `json:"Day_date" yaml:"Day_date" xml:"Day_date"`
	UpperBound      float64   `json:"Upper_bound" yaml:"Upper_bound" xml:"Upper_bound"`
}

type Products struct {
	IsActive          bool      `json:"Is_active" yaml:"Is_active" xml:"Is_active"`
	Width             float64   `json:"Width" yaml:"Width" xml:"Width"`
	MaxStockThreshold int       `json:"Max_stock_threshold" yaml:"Max_stock_threshold" xml:"Max_stock_threshold"`
	Name              string    `json:"Name" yaml:"Name" xml:"Name"`
	SearchVector      any       `json:"Search_vector" yaml:"Search_vector" xml:"Search_vector"`
	ShelfLifeDays     int       `json:"Shelf_life_days" yaml:"Shelf_life_days" xml:"Shelf_life_days"`
	MinStockThreshold int       `json:"Min_stock_threshold" yaml:"Min_stock_threshold" xml:"Min_stock_threshold"`
	Weight            float64   `json:"Weight" yaml:"Weight" xml:"Weight"`
	UpdatedAt         time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	Length            float64   `json:"Length" yaml:"Length" xml:"Length"`
	Cost              float64   `json:"Cost" yaml:"Cost" xml:"Cost"`
	LastSyncAt        time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	ExternalID        string    `json:"External_id" yaml:"External_id" xml:"External_id"`
	CreatedAt         time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Barcode           string    `json:"Barcode" yaml:"Barcode" xml:"Barcode"`
	LeadTimeDays      int       `json:"Lead_time_days" yaml:"Lead_time_days" xml:"Lead_time_days"`
	ID                any       `json:"Id" yaml:"Id" xml:"Id"`
	ReorderPoint      int       `json:"Reorder_point" yaml:"Reorder_point" xml:"Reorder_point"`
	Sku               string    `json:"Sku" yaml:"Sku" xml:"Sku"`
	Height            float64   `json:"Height" yaml:"Height" xml:"Height"`
	Price             float64   `json:"Price" yaml:"Price" xml:"Price"`
	Category          string    `json:"Category" yaml:"Category" xml:"Category"`
	Manufacturer      string    `json:"Manufacturer" yaml:"Manufacturer" xml:"Manufacturer"`
	Description       string    `json:"Description" yaml:"Description" xml:"Description"`
}

type StockPredictions struct {
	PredictionDate           time.Time `json:"Prediction_date" yaml:"Prediction_date" xml:"Prediction_date"`
	ConfidenceLevel          any       `json:"Confidence_level" yaml:"Confidence_level" xml:"Confidence_level"`
	WarehouseID              any       `json:"Warehouse_id" yaml:"Warehouse_id" xml:"Warehouse_id"`
	ID                       any       `json:"Id" yaml:"Id" xml:"Id"`
	CurrentLevel             float64   `json:"Current_level" yaml:"Current_level" xml:"Current_level"`
	DaysUntilStockout        int       `json:"Days_until_stockout" yaml:"Days_until_stockout" xml:"Days_until_stockout"`
	SuggestedReorderQuantity float64   `json:"Suggested_reorder_quantity" yaml:"Suggested_reorder_quantity" xml:"Suggested_reorder_quantity"`
	PredictedLevel           float64   `json:"Predicted_level" yaml:"Predicted_level" xml:"Predicted_level"`
	UpdatedAt                time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	PredictionHorizonDays    int       `json:"Prediction_horizon_days" yaml:"Prediction_horizon_days" xml:"Prediction_horizon_days"`
	ProductID                any       `json:"Product_id" yaml:"Product_id" xml:"Product_id"`
	CreatedAt                time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
}

type SyncConfig struct {
	IsActive            bool      `json:"Is_active" yaml:"Is_active" xml:"Is_active"`
	EntityName          string    `json:"Entity_name" yaml:"Entity_name" xml:"Entity_name"`
	CreatedAt           time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	LastSyncTimestamp   time.Time `json:"Last_sync_timestamp" yaml:"Last_sync_timestamp" xml:"Last_sync_timestamp"`
	ID                  int       `json:"Id" yaml:"Id" xml:"Id"`
	SyncIntervalMinutes int       `json:"Sync_interval_minutes" yaml:"Sync_interval_minutes" xml:"Sync_interval_minutes"`
	ErrorCount          int       `json:"Error_count" yaml:"Error_count" xml:"Error_count"`
	UpdatedAt           time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
}

type SyncLogs struct {
	RecordsProcessed int       `json:"Records_processed" yaml:"Records_processed" xml:"Records_processed"`
	RecordsUpdated   int       `json:"Records_updated" yaml:"Records_updated" xml:"Records_updated"`
	RecordsCreated   int       `json:"Records_created" yaml:"Records_created" xml:"Records_created"`
	RecordsFailed    int       `json:"Records_failed" yaml:"Records_failed" xml:"Records_failed"`
	StartTime        time.Time `json:"Start_time" yaml:"Start_time" xml:"Start_time"`
	ID               int       `json:"Id" yaml:"Id" xml:"Id"`
	EndTime          time.Time `json:"End_time" yaml:"End_time" xml:"End_time"`
	EntityName       string    `json:"Entity_name" yaml:"Entity_name" xml:"Entity_name"`
	ErrorMessage     string    `json:"Error_message" yaml:"Error_message" xml:"Error_message"`
	CreatedAt        time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Status           string    `json:"Status" yaml:"Status" xml:"Status"`
}

type TempInventory struct {
	UpdatedAt      time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	WarehouseID    any       `json:"Warehouse_id" yaml:"Warehouse_id" xml:"Warehouse_id"`
	ReorderPoint   float64   `json:"Reorder_point" yaml:"Reorder_point" xml:"Reorder_point"`
	Status         any       `json:"Status" yaml:"Status" xml:"Status"`
	Quantity       float64   `json:"Quantity" yaml:"Quantity" xml:"Quantity"`
	ID             any       `json:"Id" yaml:"Id" xml:"Id"`
	MaximumLevel   float64   `json:"Maximum_level" yaml:"Maximum_level" xml:"Maximum_level"`
	LastCountDate  time.Time `json:"Last_count_date" yaml:"Last_count_date" xml:"Last_count_date"`
	LastSyncAt     time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	MinimumLevel   float64   `json:"Minimum_level" yaml:"Minimum_level" xml:"Minimum_level"`
	CreatedAt      time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	ExpirationDate time.Time `json:"Expiration_date" yaml:"Expiration_date" xml:"Expiration_date"`
	ProductID      any       `json:"Product_id" yaml:"Product_id" xml:"Product_id"`
	LocationCode   string    `json:"Location_code" yaml:"Location_code" xml:"Location_code"`
	LotControl     string    `json:"Lot_control" yaml:"Lot_control" xml:"Lot_control"`
}

type UserPreferences struct {
	UpdatedAt       time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	PreferenceValue string    `json:"Preference_value" yaml:"Preference_value" xml:"Preference_value"`
	ID              any       `json:"Id" yaml:"Id" xml:"Id"`
	PreferenceKey   string    `json:"Preference_key" yaml:"Preference_key" xml:"Preference_key"`
	UserID          string    `json:"User_id" yaml:"User_id" xml:"User_id"`
	CreatedAt       time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
}

type Users struct {
	Name      string    `json:"Name" yaml:"Name" xml:"Name"`
	Email     string    `json:"Email" yaml:"Email" xml:"Email"`
	RoleID    int       `json:"Role_id" yaml:"Role_id" xml:"Role_id"`
	CreatedAt time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Username  string    `json:"Username" yaml:"Username" xml:"Username"`
	UpdatedAt time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	Active    bool      `json:"Active" yaml:"Active" xml:"Active"`
	Password  string    `json:"Password" yaml:"Password" xml:"Password"`
	ID        any       `json:"Id" yaml:"Id" xml:"Id"`
}

type Warehouses struct {
	LastSyncAt time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	Address    string    `json:"Address" yaml:"Address" xml:"Address"`
	CreatedAt  time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	UpdatedAt  time.Time `json:"Updated_at" yaml:"Updated_at" xml:"UpdatedAt"`
	IsActive   bool      `json:"Is_active" yaml:"Is_active" xml:"Is_active"`
	ExternalID string    `json:"External_id" yaml:"External_id" xml:"External_id"`
	Code       string    `json:"Code" yaml:"Code" xml:"Code"`
	PostalCode string    `json:"Postal_code" yaml:"Postal_code" xml:"Postal_code"`
	State      string    `json:"State" yaml:"State" xml:"State"`
	Name       string    `json:"Name" yaml:"Name" xml:"Name"`
	ID         any       `json:"Id" yaml:"Id" xml:"Id"`
	City       string    `json:"City" yaml:"City" xml:"City"`
	Country    string    `json:"Country" yaml:"Country" xml:"Country"`
}
