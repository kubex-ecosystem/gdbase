package main

import "time"

type Audit_events struct {
	Entity_type string    `json:"Entity_type" yaml:"Entity_type" xml:"Entity_type"`
	User_id     string    `json:"User_id" yaml:"User_id" xml:"User_id"`
	Changes     any       `json:"Changes" yaml:"Changes" xml:"Changes"`
	Action      string    `json:"Action" yaml:"Action" xml:"Action"`
	Id          any       `json:"Id" yaml:"Id" xml:"Id"`
	Entity_id   any       `json:"Entity_id" yaml:"Entity_id" xml:"Entity_id"`
	Created_at  time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
}

type Customers struct {
	City               string    `json:"City" yaml:"City" xml:"City"`
	Phone              string    `json:"Phone" yaml:"Phone" xml:"Phone"`
	Code               string    `json:"Code" yaml:"Code" xml:"Code"`
	Created_at         time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Id                 any       `json:"Id" yaml:"Id" xml:"Id"`
	Last_sync_at       time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	Name               string    `json:"Name" yaml:"Name" xml:"Name"`
	Email              string    `json:"Email" yaml:"Email" xml:"Email"`
	Address            string    `json:"Address" yaml:"Address" xml:"Address"`
	Last_purchase_date time.Time `json:"Last_purchase_date" yaml:"Last_purchase_date" xml:"Last_purchase_date"`
	Is_active          bool      `json:"Is_active" yaml:"Is_active" xml:"Is_active"`
	Document           string    `json:"Document" yaml:"Document" xml:"Document"`
	Updated_at         time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	Payment_terms      string    `json:"Payment_terms" yaml:"Payment_terms" xml:"Payment_terms"`
	State              string    `json:"State" yaml:"State" xml:"State"`
	Postal_code        string    `json:"Postal_code" yaml:"Postal_code" xml:"Postal_code"`
	External_id        string    `json:"External_id" yaml:"External_id" xml:"External_id"`
	Country            string    `json:"Country" yaml:"Country" xml:"Country"`
	Credit_limit       float64   `json:"Credit_limit" yaml:"Credit_limit" xml:"Credit_limit"`
}

type Inventory struct {
	Expiration_date time.Time `json:"Expiration_date" yaml:"Expiration_date" xml:"Expiration_date"`
	Lot_control     string    `json:"Lot_control" yaml:"Lot_control" xml:"Lot_control"`
	Location_code   string    `json:"Location_code" yaml:"Location_code" xml:"Location_code"`
	Product_id      any       `json:"Product_id" yaml:"Product_id" xml:"Product_id"`
	Last_sync_at    time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	Created_at      time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Minimum_level   float64   `json:"Minimum_level" yaml:"Minimum_level" xml:"Minimum_level"`
	Last_count_date time.Time `json:"Last_count_date" yaml:"Last_count_date" xml:"Last_count_date"`
	Maximum_level   float64   `json:"Maximum_level" yaml:"Maximum_level" xml:"Maximum_level"`
	Is_active       bool      `json:"Is_active" yaml:"Is_active" xml:"Is_active"`
	Quantity        float64   `json:"Quantity" yaml:"Quantity" xml:"Quantity"`
	Id              any       `json:"Id" yaml:"Id" xml:"Id"`
	Warehouse_id    any       `json:"Warehouse_id" yaml:"Warehouse_id" xml:"Warehouse_id"`
	Updated_at      time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	Status          any       `json:"Status" yaml:"Status" xml:"Status"`
	Reorder_point   float64   `json:"Reorder_point" yaml:"Reorder_point" xml:"Reorder_point"`
}

type Inventory_movements struct {
	Quantity           float64   `json:"Quantity" yaml:"Quantity" xml:"Quantity"`
	Movement_type      string    `json:"Movement_type" yaml:"Movement_type" xml:"Movement_type"`
	Previous_quantity  float64   `json:"Previous_quantity" yaml:"Previous_quantity" xml:"Previous_quantity"`
	Reference_document string    `json:"Reference_document" yaml:"Reference_document" xml:"Reference_document"`
	Inventory_id       any       `json:"Inventory_id" yaml:"Inventory_id" xml:"Inventory_id"`
	Id                 any       `json:"Id" yaml:"Id" xml:"Id"`
	Product_id         any       `json:"Product_id" yaml:"Product_id" xml:"Product_id"`
	External_id        string    `json:"External_id" yaml:"External_id" xml:"External_id"`
	Created_by         string    `json:"Created_by" yaml:"Created_by" xml:"Created_by"`
	Reason             string    `json:"Reason" yaml:"Reason" xml:"Reason"`
	Created_at         time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Warehouse_id       any       `json:"Warehouse_id" yaml:"Warehouse_id" xml:"Warehouse_id"`
	Last_sync_at       time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
}

type Order_items struct {
	Suggestion_reason string    `json:"Suggestion_reason" yaml:"Suggestion_reason" xml:"Suggestion_reason"`
	Quantity          float64   `json:"Quantity" yaml:"Quantity" xml:"Quantity"`
	Last_sync_at      time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	Is_suggested      bool      `json:"Is_suggested" yaml:"Is_suggested" xml:"Is_suggested"`
	Updated_at        time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	External_id       string    `json:"External_id" yaml:"External_id" xml:"External_id"`
	Order_id          any       `json:"Order_id" yaml:"Order_id" xml:"Order_id"`
	Created_at        time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Discount          float64   `json:"Discount" yaml:"Discount" xml:"Discount"`
	Product_id        any       `json:"Product_id" yaml:"Product_id" xml:"Product_id"`
	Id                any       `json:"Id" yaml:"Id" xml:"Id"`
	Unit_price        float64   `json:"Unit_price" yaml:"Unit_price" xml:"Unit_price"`
	Total             float64   `json:"Total" yaml:"Total" xml:"Total"`
}

type Orders struct {
	Prediction_id              any       `json:"Prediction_id" yaml:"Prediction_id" xml:"Prediction_id"`
	Id                         any       `json:"Id" yaml:"Id" xml:"Id"`
	Shipping_amount            float64   `json:"Shipping_amount" yaml:"Shipping_amount" xml:"Shipping_amount"`
	Priority                   int       `json:"Priority" yaml:"Priority" xml:"Priority"`
	Customer_id                any       `json:"Customer_id" yaml:"Customer_id" xml:"Customer_id"`
	Shipping_address           string    `json:"Shipping_address" yaml:"Shipping_address" xml:"Shipping_address"`
	Estimated_delivery_date    time.Time `json:"Estimated_delivery_date" yaml:"Estimated_delivery_date" xml:"Estimated_delivery_date"`
	Order_number               string    `json:"Order_number" yaml:"Order_number" xml:"Order_number"`
	Discount_amount            float64   `json:"Discount_amount" yaml:"Discount_amount" xml:"Discount_amount"`
	Actual_delivery_date       time.Time `json:"Actual_delivery_date" yaml:"Actual_delivery_date" xml:"Actual_delivery_date"`
	Order_date                 time.Time `json:"Order_date" yaml:"Order_date" xml:"Order_date"`
	Total_amount               float64   `json:"Total_amount" yaml:"Total_amount" xml:"Total_amount"`
	Status                     any       `json:"Status" yaml:"Status" xml:"Status"`
	Last_sync_at               time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	Tax_amount                 float64   `json:"Tax_amount" yaml:"Tax_amount" xml:"Tax_amount"`
	Expected_margin            float64   `json:"Expected_margin" yaml:"Expected_margin" xml:"Expected_margin"`
	Created_at                 time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Updated_at                 time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	Is_automatically_generated bool      `json:"Is_automatically_generated" yaml:"Is_automatically_generated" xml:"Is_automatically_generated"`
	Payment_method             string    `json:"Payment_method" yaml:"Payment_method" xml:"Payment_method"`
	Notes                      string    `json:"Notes" yaml:"Notes" xml:"Notes"`
	External_id                string    `json:"External_id" yaml:"External_id" xml:"External_id"`
	Payment_status             any       `json:"Payment_status" yaml:"Payment_status" xml:"Payment_status"`
	Final_amount               float64   `json:"Final_amount" yaml:"Final_amount" xml:"Final_amount"`
}

type Prediction_daily_data struct {
	Prediction_id    any       `json:"Prediction_id" yaml:"Prediction_id" xml:"Prediction_id"`
	Predicted_stock  float64   `json:"Predicted_stock" yaml:"Predicted_stock" xml:"Predicted_stock"`
	Created_at       time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Lower_bound      float64   `json:"Lower_bound" yaml:"Lower_bound" xml:"Lower_bound"`
	Predicted_demand float64   `json:"Predicted_demand" yaml:"Predicted_demand" xml:"Predicted_demand"`
	Id               any       `json:"Id" yaml:"Id" xml:"Id"`
	Day_date         any       `json:"Day_date" yaml:"Day_date" xml:"Day_date"`
	Upper_bound      float64   `json:"Upper_bound" yaml:"Upper_bound" xml:"Upper_bound"`
}

type Products struct {
	Is_active           bool      `json:"Is_active" yaml:"Is_active" xml:"Is_active"`
	Width               float64   `json:"Width" yaml:"Width" xml:"Width"`
	Max_stock_threshold int       `json:"Max_stock_threshold" yaml:"Max_stock_threshold" xml:"Max_stock_threshold"`
	Name                string    `json:"Name" yaml:"Name" xml:"Name"`
	Search_vector       any       `json:"Search_vector" yaml:"Search_vector" xml:"Search_vector"`
	Shelf_life_days     int       `json:"Shelf_life_days" yaml:"Shelf_life_days" xml:"Shelf_life_days"`
	Min_stock_threshold int       `json:"Min_stock_threshold" yaml:"Min_stock_threshold" xml:"Min_stock_threshold"`
	Weight              float64   `json:"Weight" yaml:"Weight" xml:"Weight"`
	Updated_at          time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	Length              float64   `json:"Length" yaml:"Length" xml:"Length"`
	Cost                float64   `json:"Cost" yaml:"Cost" xml:"Cost"`
	Last_sync_at        time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	External_id         string    `json:"External_id" yaml:"External_id" xml:"External_id"`
	Created_at          time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Barcode             string    `json:"Barcode" yaml:"Barcode" xml:"Barcode"`
	Lead_time_days      int       `json:"Lead_time_days" yaml:"Lead_time_days" xml:"Lead_time_days"`
	Id                  any       `json:"Id" yaml:"Id" xml:"Id"`
	Reorder_point       int       `json:"Reorder_point" yaml:"Reorder_point" xml:"Reorder_point"`
	Sku                 string    `json:"Sku" yaml:"Sku" xml:"Sku"`
	Height              float64   `json:"Height" yaml:"Height" xml:"Height"`
	Price               float64   `json:"Price" yaml:"Price" xml:"Price"`
	Category            string    `json:"Category" yaml:"Category" xml:"Category"`
	Manufacturer        string    `json:"Manufacturer" yaml:"Manufacturer" xml:"Manufacturer"`
	Description         string    `json:"Description" yaml:"Description" xml:"Description"`
}

type Stock_predictions struct {
	Prediction_date            time.Time `json:"Prediction_date" yaml:"Prediction_date" xml:"Prediction_date"`
	Confidence_level           any       `json:"Confidence_level" yaml:"Confidence_level" xml:"Confidence_level"`
	Warehouse_id               any       `json:"Warehouse_id" yaml:"Warehouse_id" xml:"Warehouse_id"`
	Id                         any       `json:"Id" yaml:"Id" xml:"Id"`
	Current_level              float64   `json:"Current_level" yaml:"Current_level" xml:"Current_level"`
	Days_until_stockout        int       `json:"Days_until_stockout" yaml:"Days_until_stockout" xml:"Days_until_stockout"`
	Suggested_reorder_quantity float64   `json:"Suggested_reorder_quantity" yaml:"Suggested_reorder_quantity" xml:"Suggested_reorder_quantity"`
	Predicted_level            float64   `json:"Predicted_level" yaml:"Predicted_level" xml:"Predicted_level"`
	Updated_at                 time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	Prediction_horizon_days    int       `json:"Prediction_horizon_days" yaml:"Prediction_horizon_days" xml:"Prediction_horizon_days"`
	Product_id                 any       `json:"Product_id" yaml:"Product_id" xml:"Product_id"`
	Created_at                 time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
}

type Sync_config struct {
	Is_active             bool      `json:"Is_active" yaml:"Is_active" xml:"Is_active"`
	Entity_name           string    `json:"Entity_name" yaml:"Entity_name" xml:"Entity_name"`
	Created_at            time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Last_sync_timestamp   time.Time `json:"Last_sync_timestamp" yaml:"Last_sync_timestamp" xml:"Last_sync_timestamp"`
	Id                    int       `json:"Id" yaml:"Id" xml:"Id"`
	Sync_interval_minutes int       `json:"Sync_interval_minutes" yaml:"Sync_interval_minutes" xml:"Sync_interval_minutes"`
	Error_count           int       `json:"Error_count" yaml:"Error_count" xml:"Error_count"`
	Updated_at            time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
}

type Sync_logs struct {
	Records_processed int       `json:"Records_processed" yaml:"Records_processed" xml:"Records_processed"`
	Records_updated   int       `json:"Records_updated" yaml:"Records_updated" xml:"Records_updated"`
	Records_created   int       `json:"Records_created" yaml:"Records_created" xml:"Records_created"`
	Records_failed    int       `json:"Records_failed" yaml:"Records_failed" xml:"Records_failed"`
	Start_time        time.Time `json:"Start_time" yaml:"Start_time" xml:"Start_time"`
	Id                int       `json:"Id" yaml:"Id" xml:"Id"`
	End_time          time.Time `json:"End_time" yaml:"End_time" xml:"End_time"`
	Entity_name       string    `json:"Entity_name" yaml:"Entity_name" xml:"Entity_name"`
	Error_message     string    `json:"Error_message" yaml:"Error_message" xml:"Error_message"`
	Created_at        time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Status            string    `json:"Status" yaml:"Status" xml:"Status"`
}

type Temp_inventory struct {
	Updated_at      time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	Warehouse_id    any       `json:"Warehouse_id" yaml:"Warehouse_id" xml:"Warehouse_id"`
	Reorder_point   float64   `json:"Reorder_point" yaml:"Reorder_point" xml:"Reorder_point"`
	Status          any       `json:"Status" yaml:"Status" xml:"Status"`
	Quantity        float64   `json:"Quantity" yaml:"Quantity" xml:"Quantity"`
	Id              any       `json:"Id" yaml:"Id" xml:"Id"`
	Maximum_level   float64   `json:"Maximum_level" yaml:"Maximum_level" xml:"Maximum_level"`
	Last_count_date time.Time `json:"Last_count_date" yaml:"Last_count_date" xml:"Last_count_date"`
	Last_sync_at    time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	Minimum_level   float64   `json:"Minimum_level" yaml:"Minimum_level" xml:"Minimum_level"`
	Created_at      time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Expiration_date time.Time `json:"Expiration_date" yaml:"Expiration_date" xml:"Expiration_date"`
	Product_id      any       `json:"Product_id" yaml:"Product_id" xml:"Product_id"`
	Location_code   string    `json:"Location_code" yaml:"Location_code" xml:"Location_code"`
	Lot_control     string    `json:"Lot_control" yaml:"Lot_control" xml:"Lot_control"`
}

type User_preferences struct {
	Updated_at       time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	Preference_value string    `json:"Preference_value" yaml:"Preference_value" xml:"Preference_value"`
	Id               any       `json:"Id" yaml:"Id" xml:"Id"`
	Preference_key   string    `json:"Preference_key" yaml:"Preference_key" xml:"Preference_key"`
	User_id          string    `json:"User_id" yaml:"User_id" xml:"User_id"`
	Created_at       time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
}

type Users struct {
	Name       string    `json:"Name" yaml:"Name" xml:"Name"`
	Email      string    `json:"Email" yaml:"Email" xml:"Email"`
	Role_id    int       `json:"Role_id" yaml:"Role_id" xml:"Role_id"`
	Created_at time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Username   string    `json:"Username" yaml:"Username" xml:"Username"`
	Updated_at time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	Active     bool      `json:"Active" yaml:"Active" xml:"Active"`
	Password   string    `json:"Password" yaml:"Password" xml:"Password"`
	Id         any       `json:"Id" yaml:"Id" xml:"Id"`
}

type Warehouses struct {
	Last_sync_at time.Time `json:"Last_sync_at" yaml:"Last_sync_at" xml:"Last_sync_at"`
	Address      string    `json:"Address" yaml:"Address" xml:"Address"`
	Created_at   time.Time `json:"Created_at" yaml:"Created_at" xml:"Created_at"`
	Updated_at   time.Time `json:"Updated_at" yaml:"Updated_at" xml:"Updated_at"`
	Is_active    bool      `json:"Is_active" yaml:"Is_active" xml:"Is_active"`
	External_id  string    `json:"External_id" yaml:"External_id" xml:"External_id"`
	Code         string    `json:"Code" yaml:"Code" xml:"Code"`
	Postal_code  string    `json:"Postal_code" yaml:"Postal_code" xml:"Postal_code"`
	State        string    `json:"State" yaml:"State" xml:"State"`
	Name         string    `json:"Name" yaml:"Name" xml:"Name"`
	Id           any       `json:"Id" yaml:"Id" xml:"Id"`
	City         string    `json:"City" yaml:"City" xml:"City"`
	Country      string    `json:"Country" yaml:"Country" xml:"Country"`
}
