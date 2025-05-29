package interfaces

type User interface {
	GetID() string
	GetName() string
	GetUsername() string
	GetEmail() string
	GetRoleID() uint
	GetPhone() string
	GetDocument() string
	GetAddress() string
	GetCity() string
	GetState() string
	GetCountry() string
	GetZip() string
	GetBirth() string
	GetAvatar() string
	GetPicture() string
	GetActive() bool
	SetName(name string)
	SetUsername(username string)
	SetPassword(password string) error
	SetEmail(email string)
	SetRoleID(roleID uint)
	SetPhone(phone string)
	SetDocument(document string)
	SetAddress(address string)
	SetCity(city string)
	SetState(state string)
	SetCountry(country string)
	SetZip(zip string)
	SetBirth(birth string)
	SetAvatar(avatar string)
	SetPicture(picture string)
	SetActive(active bool)
	CheckPasswordHash(password string) bool
	Sanitize()
	Validate() error

	getUserObj() *User
}
