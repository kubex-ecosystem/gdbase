package clients

import (
	"time"

	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

// ClientStatus represents the status of a client
type ClientStatus string

const (
	Active   ClientStatus = "ACTIVE"
	Inactive ClientStatus = "INACTIVE"
	Pending  ClientStatus = "PENDING"
	Blocked  ClientStatus = "BLOCKED"
	Archived ClientStatus = "ARCHIVED"
)

// ClientType represents the type of a client
type ClientType string

const (
	Individual ClientType = "individual"
	Company    ClientType = "company"
)

// IClientContact interface for ClientContact
// IClientContact defines methods for manipulating ClientContact fields
//
//go:generate mockgen -destination=../../mocks/mock_client_contact.go -package=mocks . IClientContact
type IClientContact interface {
	GetPhone() *string
	SetPhone(phone *string)
	GetMobilePhone() *string
	SetMobilePhone(mobilePhone *string)
	GetEmail() *string
	SetEmail(email *string)
	GetContactName() *string
	SetContactName(contactName *string)
}

// ClientContact represents the contact information of a client
type ClientContact struct {
	Phone       *string `json:"phone,omitempty" xml:"phone,omitempty" yaml:"phone,omitempty" gorm:"column:phone"`
	MobilePhone *string `json:"mobilePhone,omitempty" xml:"mobilePhone,omitempty" yaml:"mobilePhone,omitempty" gorm:"column:mobile_phone"`
	Email       *string `json:"email,omitempty" xml:"email,omitempty" yaml:"email,omitempty" gorm:"column:email"`
	ContactName *string `json:"contactName,omitempty" xml:"contactName,omitempty" yaml:"contactName,omitempty" gorm:"column:contact_name"`
}

func (c *ClientContact) GetPhone() *string        { return c.Phone }
func (c *ClientContact) SetPhone(phone *string)   { c.Phone = phone }
func (c *ClientContact) GetMobilePhone() *string  { return c.MobilePhone }
func (c *ClientContact) SetMobilePhone(m *string) { c.MobilePhone = m }
func (c *ClientContact) GetEmail() *string        { return c.Email }
func (c *ClientContact) SetEmail(email *string)   { c.Email = email }
func (c *ClientContact) GetContactName() *string  { return c.ContactName }
func (c *ClientContact) SetContactName(n *string) { c.ContactName = n }

// IArchiveInfo interface for ArchiveInfo
type IArchiveInfo interface {
	GetArchivedAt() time.Time
	SetArchivedAt(t time.Time)
	GetArchivedBy() string
	SetArchivedBy(by string)
	GetReason() string
	SetReason(reason string)
}

// ArchiveInfo represents archival information of a client
type ArchiveInfo struct {
	ArchivedAt time.Time `json:"archivedAt" xml:"archivedAt" yaml:"archivedAt" gorm:"column:archived_at"`
	ArchivedBy string    `json:"archivedBy" xml:"archivedBy" yaml:"archivedBy" gorm:"column:archived_by"`
	Reason     string    `json:"reason" xml:"reason" yaml:"reason" gorm:"column:reason"`
}

func (a *ArchiveInfo) GetArchivedAt() time.Time  { return a.ArchivedAt }
func (a *ArchiveInfo) SetArchivedAt(t time.Time) { a.ArchivedAt = t }
func (a *ArchiveInfo) GetArchivedBy() string     { return a.ArchivedBy }
func (a *ArchiveInfo) SetArchivedBy(by string)   { a.ArchivedBy = by }
func (a *ArchiveInfo) GetReason() string         { return a.Reason }
func (a *ArchiveInfo) SetReason(reason string)   { a.Reason = reason }

// IClientDetailed interface for ClientDetailed
type IClientDetailed interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetCode() *string
	SetCode(code *string)
	GetTradingName() *string
	SetTradingName(name *string)
	GetDocumentType() ClientType
	SetDocumentType(t ClientType)
	GetContact() IClientContact
	SetContact(contact IClientContact)
	GetMainAddress() t.IAddress
	SetMainAddress(addr t.IAddress)
	GetAddresses() []t.IAddress
	SetAddresses(addrs []t.IAddress)
	GetStatus() ClientStatus
	SetStatus(status ClientStatus)
	GetCreditLimit() *float64
	SetCreditLimit(limit *float64)
	GetPaymentTerms() *string
	SetPaymentTerms(terms *string)
	GetNotes() *string
	SetNotes(notes *string)
	GetTotalOrders() *int
	SetTotalOrders(total *int)
	GetTotalSpent() *float64
	SetTotalSpent(spent *float64)
	GetLastOrderDate() *time.Time
	SetLastOrderDate(date *time.Time)
	GetArchiveInfo() IArchiveInfo
	SetArchiveInfo(info IArchiveInfo)
}

// ClientDetailed represents a detailed client structure
type ClientDetailed struct {
	ID            string        `json:"id" xml:"id" yaml:"id" gorm:"column:id;primaryKey"`
	Code          *string       `json:"code,omitempty" xml:"code,omitempty" yaml:"code,omitempty" gorm:"column:code"`
	TradingName   *string       `json:"tradingName,omitempty" xml:"tradingName,omitempty" yaml:"tradingName,omitempty" gorm:"column:trading_name"`
	DocumentType  ClientType    `json:"documentType" xml:"documentType" yaml:"documentType" gorm:"column:document_type"`
	Contact       ClientContact `json:"contact" xml:"contact" yaml:"contact" gorm:"embedded;embeddedPrefix:contact_"`
	MainAddress   t.Address     `json:"mainAddress" xml:"mainAddress" yaml:"mainAddress" gorm:"embedded;embeddedPrefix:main_address_"`
	Addresses     []t.Address   `json:"addresses" xml:"addresses" yaml:"addresses" gorm:"-"`
	Status        ClientStatus  `json:"status" xml:"status" yaml:"status" gorm:"column:status"`
	CreditLimit   *float64      `json:"creditLimit,omitempty" xml:"creditLimit,omitempty" yaml:"creditLimit,omitempty" gorm:"column:credit_limit"`
	PaymentTerms  *string       `json:"paymentTerms,omitempty" xml:"paymentTerms,omitempty" yaml:"paymentTerms,omitempty" gorm:"column:payment_terms"`
	Notes         *string       `json:"notes,omitempty" xml:"notes,omitempty" yaml:"notes,omitempty" gorm:"column:notes"`
	TotalOrders   *int          `json:"totalOrders,omitempty" xml:"totalOrders,omitempty" yaml:"totalOrders,omitempty" gorm:"column:total_orders"`
	TotalSpent    *float64      `json:"totalSpent,omitempty" xml:"totalSpent,omitempty" yaml:"totalSpent,omitempty" gorm:"column:total_spent"`
	LastOrderDate *time.Time    `json:"lastOrderDate,omitempty" xml:"lastOrderDate,omitempty" yaml:"lastOrderDate,omitempty" gorm:"column:last_order_date"`
	ArchiveInfo   *ArchiveInfo  `json:"archiveInfo,omitempty" xml:"archiveInfo,omitempty" yaml:"archiveInfo,omitempty" gorm:"embedded;embeddedPrefix:archive_"`
	CreatedAt     time.Time     `json:"createdAt" xml:"createdAt" yaml:"createdAt" gorm:"column:created_at"`
	UpdatedAt     time.Time     `json:"updatedAt" xml:"updatedAt" yaml:"updatedAt" gorm:"column:updated_at"`
	LastSync      time.Time     `json:"lastSync" xml:"lastSync" yaml:"lastSync" gorm:"column:last_sync"`
}

func (c *ClientDetailed) TableName() string            { return "clients" }
func (c *ClientDetailed) GetID() string                { return c.ID }
func (c *ClientDetailed) SetID(id string)              { c.ID = id }
func (c *ClientDetailed) GetCode() *string             { return c.Code }
func (c *ClientDetailed) SetCode(code *string)         { c.Code = code }
func (c *ClientDetailed) GetTradingName() *string      { return c.TradingName }
func (c *ClientDetailed) SetTradingName(name *string)  { c.TradingName = name }
func (c *ClientDetailed) GetDocumentType() ClientType  { return c.DocumentType }
func (c *ClientDetailed) SetDocumentType(t ClientType) { c.DocumentType = t }
func (c *ClientDetailed) GetContact() IClientContact   { return &c.Contact }
func (c *ClientDetailed) SetContact(contact IClientContact) {
	if v, ok := contact.(*ClientContact); ok {
		c.Contact = *v
	}
}
func (c *ClientDetailed) GetMainAddress() t.IAddress { return &c.MainAddress }
func (c *ClientDetailed) SetMainAddress(addr t.IAddress) {
	if v, ok := addr.(*t.Address); ok {
		c.MainAddress = *v
	}
}
func (c *ClientDetailed) GetAddresses() []t.IAddress {
	addrs := make([]t.IAddress, len(c.Addresses))
	for i := range c.Addresses {
		addrs[i] = &c.Addresses[i]
	}
	return addrs
}
func (c *ClientDetailed) SetAddresses(addrs []t.IAddress) {
	c.Addresses = make([]t.Address, len(addrs))
	for i, a := range addrs {
		if v, ok := a.(*t.Address); ok {
			c.Addresses[i] = *v
		}
	}
}
func (c *ClientDetailed) GetStatus() ClientStatus          { return c.Status }
func (c *ClientDetailed) SetStatus(status ClientStatus)    { c.Status = status }
func (c *ClientDetailed) GetCreditLimit() *float64         { return c.CreditLimit }
func (c *ClientDetailed) SetCreditLimit(limit *float64)    { c.CreditLimit = limit }
func (c *ClientDetailed) GetPaymentTerms() *string         { return c.PaymentTerms }
func (c *ClientDetailed) SetPaymentTerms(terms *string)    { c.PaymentTerms = terms }
func (c *ClientDetailed) GetNotes() *string                { return c.Notes }
func (c *ClientDetailed) SetNotes(notes *string)           { c.Notes = notes }
func (c *ClientDetailed) GetTotalOrders() *int             { return c.TotalOrders }
func (c *ClientDetailed) SetTotalOrders(total *int)        { c.TotalOrders = total }
func (c *ClientDetailed) GetTotalSpent() *float64          { return c.TotalSpent }
func (c *ClientDetailed) SetTotalSpent(spent *float64)     { c.TotalSpent = spent }
func (c *ClientDetailed) GetLastOrderDate() *time.Time     { return c.LastOrderDate }
func (c *ClientDetailed) SetLastOrderDate(date *time.Time) { c.LastOrderDate = date }
func (c *ClientDetailed) GetArchiveInfo() IArchiveInfo     { return c.ArchiveInfo }
func (c *ClientDetailed) SetArchiveInfo(info IArchiveInfo) {
	if v, ok := info.(*ArchiveInfo); ok {
		c.ArchiveInfo = v
	}
}

// ClientResponse represents a paginated response of clients
type ClientResponse struct {
	Data       []ClientDetailed `json:"data" xml:"data" yaml:"data" gorm:"-"`
	Total      int              `json:"total" xml:"total" yaml:"total" gorm:"-"`
	TotalPages int              `json:"totalPages" xml:"totalPages" yaml:"totalPages" gorm:"-"`
	Page       int              `json:"page" xml:"page" yaml:"page" gorm:"-"`
	Limit      int              `json:"limit" xml:"limit" yaml:"limit" gorm:"-"`
}

// ClientSortField represents the fields by which clients can be sorted
type ClientSortField string

const (
	Name        ClientSortField = "name"
	Document    ClientSortField = "document"
	CreatedAt   ClientSortField = "createdAt"
	UpdatedAt   ClientSortField = "updatedAt"
	City        ClientSortField = "city"
	CreditLimit ClientSortField = "creditLimit"
	Status      ClientSortField = "status"
)

// SortDirection represents the direction of sorting
type SortDirection string

const (
	Asc  SortDirection = "asc"
	Desc SortDirection = "desc"
)

// ClientFilterParams represents the parameters for filtering clients
type ClientFilterParams struct {
	Name            *string          `json:"name,omitempty"`
	Document        *string          `json:"document,omitempty"`
	Email           *string          `json:"email,omitempty"`
	Phone           *string          `json:"phone,omitempty"`
	Status          *ClientStatus    `json:"status,omitempty"`
	IncludeArchived bool             `json:"includeArchived"`
	SortBy          *ClientSortField `json:"sortBy,omitempty"`
	SortDirection   *SortDirection   `json:"sortDirection,omitempty"`
}

// ArchiveClientDTO represents the data for archiving a client
type ArchiveClientDTO struct {
	Reason     string `json:"reason"`
	ArchivedBy string `json:"archivedBy"`
}

// CreateClientDTO represents the data for creating a client
type CreateClientDTO struct {
	Code         *string       `json:"code,omitempty"`
	TradingName  *string       `json:"tradingName,omitempty"`
	DocumentType ClientType    `json:"documentType"`
	Contact      ClientContact `json:"contact"`
	MainAddress  t.Address     `json:"mainAddress"`
	Addresses    []t.Address   `json:"addresses"`
	Status       ClientStatus  `json:"status"`
	CreditLimit  *float64      `json:"creditLimit,omitempty"`
	PaymentTerms *string       `json:"paymentTerms,omitempty"`
	Notes        *string       `json:"notes,omitempty"`
}

// UpdateClientDTO represents the data for updating a client
type UpdateClientDTO struct {
	Code         *string        `json:"code,omitempty"`
	TradingName  *string        `json:"tradingName,omitempty"`
	DocumentType *ClientType    `json:"documentType,omitempty"`
	Contact      *ClientContact `json:"contact,omitempty"`
	MainAddress  *t.Address     `json:"mainAddress,omitempty"`
	Addresses    *[]t.Address   `json:"addresses,omitempty"`
	Status       *ClientStatus  `json:"status,omitempty"`
	CreditLimit  *float64       `json:"creditLimit,omitempty"`
	PaymentTerms *string        `json:"paymentTerms,omitempty"`
	Notes        *string        `json:"notes,omitempty"`
}

// PaginationParams represents pagination parameters
type PaginationParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// PaginatedClientResult represents a paginated result of clients
type PaginatedClientResult struct {
	Data       []ClientDetailed `json:"data"`
	Total      int              `json:"total"`
	Page       int              `json:"page"`
	Limit      int              `json:"limit"`
	TotalPages int              `json:"totalPages"`
}
