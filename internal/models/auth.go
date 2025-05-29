package models

// User represents an authenticated user
// Equivalent to the User interface in TypeScript
import "time"

type User struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Role        UserRole   `json:"role"`
	Permissions []string   `json:"permissions,omitempty"`
	PhotoURL    *string    `json:"photoUrl,omitempty"`
	Region      *string    `json:"region,omitempty"`
	LastLogin   *time.Time `json:"lastLogin,omitempty"`
}

// UserRole represents the roles a user can have
type UserRole string

const (
	Admin      UserRole = "admin"
	Backoffice UserRole = "backoffice"
	Supervisor UserRole = "supervisor"
	Seller     UserRole = "seller"
)

// Permission represents the permissions a user can have
type Permission string

const (
	// System Administration
	ManageSystemSettings Permission = "manage_system_settings"
	ManageUsers          Permission = "manage_users"
	ManageRoles          Permission = "manage_roles"
	ViewSystemLogs       Permission = "view_system_logs"

	// Master Data Management
	ManageAllProducts Permission = "manage_all_products"
	ManagePrices      Permission = "manage_prices"
	ManageAllClients  Permission = "manage_all_clients"
	ManageCategories  Permission = "manage_categories"

	// Orders and Sales
	ManageAllOrders Permission = "manage_all_orders"
	CreateOrders    Permission = "create_orders"
	ApproveOrders   Permission = "approve_orders"
	CancelOrders    Permission = "cancel_orders"
	EditOrders      Permission = "edit_orders"

	// Commissions
	ManageCommissionRules Permission = "manage_commission_rules"
	ProcessCommissions    Permission = "process_commissions"
	ViewAllCommissions    Permission = "view_all_commissions"

	// Reports and Analysis
	ViewAllReports Permission = "view_all_reports"
	ExportReports  Permission = "export_reports"
)

// UserRolePermissions maps roles to their permissions
var UserRolePermissions = map[UserRole][]Permission{
	Admin: {
		ManageSystemSettings,
		ManageUsers,
		ManageRoles,
		ViewSystemLogs,
		ManageAllProducts,
		ManagePrices,
		ManageAllClients,
		ManageCategories,
		ManageAllOrders,
		CreateOrders,
		ApproveOrders,
		CancelOrders,
		EditOrders,
		ManageCommissionRules,
		ProcessCommissions,
		ViewAllCommissions,
		ViewAllReports,
		ExportReports,
	},
	Backoffice: {
		ManageAllProducts,
		ManagePrices,
		ManageAllClients,
		ManageCategories,
		ManageAllOrders,
		CreateOrders,
		ApproveOrders,
		CancelOrders,
		EditOrders,
		ProcessCommissions,
		ViewAllCommissions,
		ViewAllReports,
		ExportReports,
	},
	Supervisor: {
		ViewAllReports,
		CreateOrders,
	},
	Seller: {
		CreateOrders,
	},
}
