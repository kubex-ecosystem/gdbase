// Package products contains models and interfaces related to products
package products

import (
	"time"

	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

// PriceTable represents a product's price table
type PriceTable struct {
	ID    string   `json:"id" xml:"id" yaml:"id" gorm:"column:id"`
	Name  string   `json:"name" xml:"name" yaml:"name" gorm:"column:name"`
	Price *t.Money `json:"price" xml:"price" yaml:"price" gorm:"column:price"`
}

// IPriceTable interface for abstraction
type IPriceTable interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetName() string
	SetName(name string)
	GetPrice() *t.Money
	SetPrice(price *t.Money)
}

func (p *PriceTable) TableName() string       { return "price_tables" }
func (p *PriceTable) GetID() string           { return p.ID }
func (p *PriceTable) SetID(id string)         { p.ID = id }
func (p *PriceTable) GetName() string         { return p.Name }
func (p *PriceTable) SetName(name string)     { p.Name = name }
func (p *PriceTable) GetPrice() *t.Money      { return p.Price }
func (p *PriceTable) SetPrice(price *t.Money) { p.Price = price }

// ProductCategory represents a product category
type ProductCategory struct {
	ID       string  `json:"id" xml:"id" yaml:"id" gorm:"column:id"`
	Name     string  `json:"name" xml:"name" yaml:"name" gorm:"column:name"`
	ParentID *string `json:"parentId,omitempty" xml:"parentId,omitempty" yaml:"parentId,omitempty" gorm:"column:parent_id"`
}

// IProductCategory interface for abstraction
//
//go:generate mockgen -destination=../../mocks/mock_product_category.go -package=mocks . IProductCategory
type IProductCategory interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetName() string
	SetName(name string)
	GetParentID() *string
	SetParentID(parentID *string)
}

func (c *ProductCategory) TableName() string       { return "product_categories" }
func (c *ProductCategory) GetID() string           { return c.ID }
func (c *ProductCategory) SetID(id string)         { c.ID = id }
func (c *ProductCategory) GetName() string         { return c.Name }
func (c *ProductCategory) SetName(name string)     { c.Name = name }
func (c *ProductCategory) GetParentID() *string    { return c.ParentID }
func (c *ProductCategory) SetParentID(pid *string) { c.ParentID = pid }

// Product represents a product
type Product struct {
	ID                    string           `json:"id" xml:"id" yaml:"id" gorm:"column:id;primaryKey"`
	Code                  string           `json:"code" xml:"code" yaml:"code" gorm:"column:code"`
	SKU                   string           `json:"sku" xml:"sku" yaml:"sku" gorm:"column:sku"`
	EAN                   *string          `json:"ean,omitempty" xml:"ean,omitempty" yaml:"ean,omitempty" gorm:"column:ean"`
	Name                  string           `json:"name" xml:"name" yaml:"name" gorm:"column:name"`
	Description           string           `json:"description" xml:"description" yaml:"description" gorm:"column:description"`
	ImageURL              *string          `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" yaml:"imageUrl,omitempty" gorm:"column:image_url"`
	Unit                  string           `json:"unit" xml:"unit" yaml:"unit" gorm:"column:unit"`
	Weight                *float64         `json:"weight,omitempty" xml:"weight,omitempty" yaml:"weight,omitempty" gorm:"column:weight"`
	Dimensions            *Dimensions      `json:"dimensions,omitempty" xml:"dimensions,omitempty" yaml:"dimensions,omitempty" gorm:"embedded;embeddedPrefix:dimensions_"`
	CategoryID            string           `json:"categoryId" xml:"categoryId" yaml:"categoryId" gorm:"column:category_id"`
	Category              *ProductCategory `json:"category,omitempty" xml:"category,omitempty" yaml:"category,omitempty" gorm:"foreignKey:CategoryID"`
	PriceTables           []PriceTable     `json:"priceTables" xml:"priceTables" yaml:"priceTables" gorm:"-"`
	Stock                 Stock            `json:"stock" xml:"stock" yaml:"stock" gorm:"embedded;embeddedPrefix:stock_"`
	MinOrderQuantity      *int             `json:"minOrderQuantity,omitempty" xml:"minOrderQuantity,omitempty" yaml:"minOrderQuantity,omitempty" gorm:"column:min_order_quantity"`
	MultipleOrderQuantity *int             `json:"multipleOrderQuantity,omitempty" xml:"multipleOrderQuantity,omitempty" yaml:"multipleOrderQuantity,omitempty" gorm:"column:multiple_order_quantity"`
	IsActive              bool             `json:"isActive" xml:"isActive" yaml:"isActive" gorm:"column:is_active"`
	CreatedAt             time.Time        `json:"createdAt" xml:"createdAt" yaml:"createdAt" gorm:"column:created_at"`
	UpdatedAt             time.Time        `json:"updatedAt" xml:"updatedAt" yaml:"updatedAt" gorm:"column:updated_at"`
}

// IProduct interface for abstraction
type IProduct interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetCode() string
	SetCode(code string)
	GetSKU() string
	SetSKU(sku string)
	GetEAN() *string
	SetEAN(ean *string)
	GetName() string
	SetName(name string)
	GetDescription() string
	SetDescription(desc string)
	GetImageURL() *string
	SetImageURL(url *string)
	GetUnit() string
	SetUnit(unit string)
	GetWeight() *float64
	SetWeight(weight *float64)
	GetDimensions() IDimensions
	SetDimensions(dim IDimensions)
	GetCategoryID() string
	SetCategoryID(id string)
	GetCategory() IProductCategory
	SetCategory(cat IProductCategory)
	GetPriceTables() []IPriceTable
	SetPriceTables(tables []IPriceTable)
	GetStock() IStock
	SetStock(stock IStock)
	GetMinOrderQuantity() *int
	SetMinOrderQuantity(qty *int)
	GetMultipleOrderQuantity() *int
	SetMultipleOrderQuantity(qty *int)
	GetIsActive() bool
	SetIsActive(active bool)
	GetCreatedAt() time.Time
	SetCreatedAt(t time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(t time.Time)
}

func (p *Product) TableName() string          { return "products" }
func (p *Product) GetID() string              { return p.ID }
func (p *Product) SetID(id string)            { p.ID = id }
func (p *Product) GetCode() string            { return p.Code }
func (p *Product) SetCode(code string)        { p.Code = code }
func (p *Product) GetSKU() string             { return p.SKU }
func (p *Product) SetSKU(sku string)          { p.SKU = sku }
func (p *Product) GetEAN() *string            { return p.EAN }
func (p *Product) SetEAN(ean *string)         { p.EAN = ean }
func (p *Product) GetName() string            { return p.Name }
func (p *Product) SetName(name string)        { p.Name = name }
func (p *Product) GetDescription() string     { return p.Description }
func (p *Product) SetDescription(desc string) { p.Description = desc }
func (p *Product) GetImageURL() *string       { return p.ImageURL }
func (p *Product) SetImageURL(url *string)    { p.ImageURL = url }
func (p *Product) GetUnit() string            { return p.Unit }
func (p *Product) SetUnit(unit string)        { p.Unit = unit }
func (p *Product) GetWeight() *float64        { return p.Weight }
func (p *Product) SetWeight(weight *float64)  { p.Weight = weight }
func (p *Product) GetDimensions() IDimensions { return p.Dimensions }
func (p *Product) SetDimensions(dim IDimensions) {
	if v, ok := dim.(*Dimensions); ok {
		p.Dimensions = v
	}
}
func (p *Product) GetCategoryID() string         { return p.CategoryID }
func (p *Product) SetCategoryID(id string)       { p.CategoryID = id }
func (p *Product) GetCategory() IProductCategory { return p.Category }
func (p *Product) SetCategory(cat IProductCategory) {
	if v, ok := cat.(*ProductCategory); ok {
		p.Category = v
	}
}
func (p *Product) GetPriceTables() []IPriceTable {
	tables := make([]IPriceTable, len(p.PriceTables))
	for i := range p.PriceTables {
		tables[i] = &p.PriceTables[i]
	}
	return tables
}
func (p *Product) SetPriceTables(tables []IPriceTable) {
	p.PriceTables = make([]PriceTable, len(tables))
	for i, t := range tables {
		if v, ok := t.(*PriceTable); ok {
			p.PriceTables[i] = *v
		}
	}
}
func (p *Product) GetStock() IStock { return &p.Stock }
func (p *Product) SetStock(stock IStock) {
	if v, ok := stock.(*Stock); ok {
		p.Stock = *v
	}
}
func (p *Product) GetMinOrderQuantity() *int         { return p.MinOrderQuantity }
func (p *Product) SetMinOrderQuantity(qty *int)      { p.MinOrderQuantity = qty }
func (p *Product) GetMultipleOrderQuantity() *int    { return p.MultipleOrderQuantity }
func (p *Product) SetMultipleOrderQuantity(qty *int) { p.MultipleOrderQuantity = qty }
func (p *Product) GetIsActive() bool                 { return p.IsActive }
func (p *Product) SetIsActive(active bool)           { p.IsActive = active }
func (p *Product) GetCreatedAt() time.Time           { return p.CreatedAt }
func (p *Product) SetCreatedAt(t time.Time)          { p.CreatedAt = t }
func (p *Product) GetUpdatedAt() time.Time           { return p.UpdatedAt }
func (p *Product) SetUpdatedAt(t time.Time)          { p.UpdatedAt = t }

// Dimensions represents the dimensions of a product
type Dimensions struct {
	Length float64 `json:"length" xml:"length" yaml:"length" gorm:"column:length"`
	Width  float64 `json:"width" xml:"width" yaml:"width" gorm:"column:width"`
	Height float64 `json:"height" xml:"height" yaml:"height" gorm:"column:height"`
}

// IDimensions interface for abstraction
type IDimensions interface {
	GetLength() float64
	SetLength(length float64)
	GetWidth() float64
	SetWidth(width float64)
	GetHeight() float64
	SetHeight(height float64)
}

func (d *Dimensions) GetLength() float64       { return d.Length }
func (d *Dimensions) SetLength(length float64) { d.Length = length }
func (d *Dimensions) GetWidth() float64        { return d.Width }
func (d *Dimensions) SetWidth(width float64)   { d.Width = width }
func (d *Dimensions) GetHeight() float64       { return d.Height }
func (d *Dimensions) SetHeight(height float64) { d.Height = height }

// Stock represents the stock information of a product
type Stock struct {
	Available int `json:"available" xml:"available" yaml:"available" gorm:"column:available"`
	Reserved  int `json:"reserved" xml:"reserved" yaml:"reserved" gorm:"column:reserved"`
	Virtual   int `json:"virtual" xml:"virtual" yaml:"virtual" gorm:"column:virtual"`
}

// IStock interface for abstraction
type IStock interface {
	GetAvailable() int
	SetAvailable(avail int)
	GetReserved() int
	SetReserved(res int)
	GetVirtual() int
	SetVirtual(virt int)
}

func (s *Stock) GetAvailable() int      { return s.Available }
func (s *Stock) SetAvailable(avail int) { s.Available = avail }
func (s *Stock) GetReserved() int       { return s.Reserved }
func (s *Stock) SetReserved(res int)    { s.Reserved = res }
func (s *Stock) GetVirtual() int        { return s.Virtual }
func (s *Stock) SetVirtual(virt int)    { s.Virtual = virt }

// ProductFilters represents filters for querying products
type ProductFilters struct {
	Search     *string `json:"search,omitempty" xml:"search,omitempty" yaml:"search,omitempty" gorm:"-"`
	CategoryID *string `json:"categoryId,omitempty" xml:"categoryId,omitempty" yaml:"categoryId,omitempty" gorm:"-"`
	InStock    *bool   `json:"inStock,omitempty" xml:"inStock,omitempty" yaml:"inStock,omitempty" gorm:"-"`
}
