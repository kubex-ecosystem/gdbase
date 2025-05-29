package types

type Address struct {
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	PostalCode string  `json:"postalCode"`
	Country    string  `json:"country"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type AddressList struct {
	Addresses []Address `json:"addresses"`
}

type AddressListResponse struct {
	Addresses  []Address `json:"addresses"`
	Total      int       `json:"total"`
	TotalPages int       `json:"totalPages"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
}

type AddressFilterParams struct {
	Street          *string  `json:"street,omitempty"`
	City            *string  `json:"city,omitempty"`
	State           *string  `json:"state,omitempty"`
	PostalCode      *string  `json:"postalCode,omitempty"`
	Country         *string  `json:"country,omitempty"`
	Latitude        *float64 `json:"latitude,omitempty"`
	Longitude       *float64 `json:"longitude,omitempty"`
	SortBy          *string  `json:"sortBy,omitempty"`
	SortDirection   *string  `json:"sortDirection,omitempty"`
	IncludeArchived bool     `json:"includeArchived"`
}

type AddressSortField string
type AddressSortDirection string

const (
	StreetAddressSortField     AddressSortField     = "street"
	CityAddressSortField       AddressSortField     = "city"
	StateAddressSortField      AddressSortField     = "state"
	PostalCodeAddressSortField AddressSortField     = "postalCode"
	CountryAddressSortField    AddressSortField     = "country"
	LatitudeAddressSortField   AddressSortField     = "latitude"
	LongitudeAddressSortField  AddressSortField     = "longitude"
	AscAddressSortDirection    AddressSortDirection = "asc"
	DescAddressSortDirection   AddressSortDirection = "desc"
)

type AddressResponse struct {
	Address Address `json:"address"`
}

type CreateAddressDTO struct {
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	PostalCode string  `json:"postalCode"`
	Country    string  `json:"country"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type UpdateAddressDTO struct {
	Street     *string  `json:"street,omitempty"`
	City       *string  `json:"city,omitempty"`
	State      *string  `json:"state,omitempty"`
	PostalCode *string  `json:"postalCode,omitempty"`
	Country    *string  `json:"country,omitempty"`
	Latitude   *float64 `json:"latitude,omitempty"`
	Longitude  *float64 `json:"longitude,omitempty"`
}

type AddressStatus string
type AddressStatusEnum string

type AddressStatusEnumList struct {
	AddressStatusEnum []AddressStatusEnum `json:"addressStatusEnum"`
}

type AddressStatusEnumResponse struct {
	AddressStatusEnum AddressStatusEnum `json:"addressStatusEnum"`
}

type AddressStatusEnumListResponse struct {
	AddressStatusEnumList []AddressStatusEnum `json:"addressStatusEnumList"`
	Total                 int                 `json:"total"`
	TotalPages            int                 `json:"totalPages"`
	Page                  int                 `json:"page"`
	Limit                 int                 `json:"limit"`
}

// IAddress interface for abstraction and encapsulation
//
//go:generate mockgen -destination=../mocks/mock_address.go -package=mocks . IAddress
type IAddress interface {
	GetStreet() string
	SetStreet(street string)
	GetCity() string
	SetCity(city string)
	GetState() string
	SetState(state string)
	GetPostalCode() string
	SetPostalCode(postalCode string)
	GetCountry() string
	SetCountry(country string)
	GetLatitude() float64
	SetLatitude(lat float64)
	GetLongitude() float64
	SetLongitude(lon float64)
}

func (a *Address) GetStreet() string               { return a.Street }
func (a *Address) SetStreet(street string)         { a.Street = street }
func (a *Address) GetCity() string                 { return a.City }
func (a *Address) SetCity(city string)             { a.City = city }
func (a *Address) GetState() string                { return a.State }
func (a *Address) SetState(state string)           { a.State = state }
func (a *Address) GetPostalCode() string           { return a.PostalCode }
func (a *Address) SetPostalCode(postalCode string) { a.PostalCode = postalCode }
func (a *Address) GetCountry() string              { return a.Country }
func (a *Address) SetCountry(country string)       { a.Country = country }
func (a *Address) GetLatitude() float64            { return a.Latitude }
func (a *Address) SetLatitude(lat float64)         { a.Latitude = lat }
func (a *Address) GetLongitude() float64           { return a.Longitude }
func (a *Address) SetLongitude(lon float64)        { a.Longitude = lon }
