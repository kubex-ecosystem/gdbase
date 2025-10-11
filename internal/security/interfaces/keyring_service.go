// Package interfaces provides the interface for keyring services
package interfaces

type IKeyringService interface {
	StorePassword(password string) error
	RetrievePassword() (string, error)
}
