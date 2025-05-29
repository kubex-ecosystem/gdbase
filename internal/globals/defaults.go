package globals

import (
	"fmt"
	"os"

	crp "github.com/rafa-mori/gdbase/internal/security/crypto"
	krs "github.com/rafa-mori/gdbase/internal/security/external"
	gl "github.com/rafa-mori/gdbase/logger"
)

var (
	KubexKeyringName = "kubex"
	KubexKeyringKey  string
)

func init() {
	var err error
	if KubexKeyringKey == "" {
		KubexKeyringKey, err = GetOrGenPasswordKeyringPass(KubexKeyringName)
		if err != nil {
			gl.Log("fatal", fmt.Sprintf("Error initializing keyring: %v", err))
		}
	}
}

const (
	KeyringService            = "kubex"
	DefaultGoBEKeyPath        = "$HOME/.kubex/gobe/gobe-key.pem"
	DefaultGoBECertPath       = "$HOME/.kubex/gobe/gobe-cert.pem"
	DefaultGodoBaseConfigPath = "$HOME/.kubex/gdbase/config/config.json"
	DefaultVolumesDir         = "$HOME/.kubex/volumes"
	DefaultRedisVolume        = "$HOME/.kubex/volumes/redis"
	DefaultPostgresVolume     = "$HOME/.kubex/volumes/postgresql"
	DefaultMongoVolume        = "$HOME/.kubex/volumes/mongo"
	DefaultRabbitMQVolume     = "$HOME/.kubex/volumes/rabbitmq"
)

type GenericRepo interface {
	Create(u interface{}) (interface{}, error)
	FindOne(where ...interface{}) (interface{}, error)
	FindAll(where ...interface{}) ([]interface{}, error)
	Update(u interface{}) (interface{}, error)
	Delete(id uint) error
}
type Certificate struct {
}
type Docker struct{}
type FileSystem struct {
}
type Cache struct {
	Enabled          bool   `json:"enabled"`
	Setup            bool   `json:"setup"`
	CacheDir         string `json:"cache_dir"`
	SetupFlagPath    string `json:"setup_flag_path"`
	DepsFlagPath     string `json:"deps_flag_path"`
	ServicesFlagPath string `json:"services_flag_path"`
	VaultFlagPath    string `json:"vault_flag_path"`
}
type ValidationError struct {
	Field   string
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}
func (v *ValidationError) FieldError() map[string]string {
	return map[string]string{v.Field: v.Message}
}
func (v *ValidationError) FieldsError() map[string]string {
	return map[string]string{v.Field: v.Message}
}
func (v *ValidationError) ErrorOrNil() error {
	return v
}

var (
	ErrUsernameRequired = &ValidationError{Field: "username", Message: "Username is required"}
	ErrPasswordRequired = &ValidationError{Field: "password", Message: "Password is required"}
	ErrEmailRequired    = &ValidationError{Field: "email", Message: "Email is required"}
	ErrDBNotProvided    = &ValidationError{Field: "db", Message: "Database not provided"}
	ErrModelNotFound    = &ValidationError{Field: "model", Message: "Model not found"}
)

// GetOrGenPasswordKeyringPass retrieves the password from the keyring or generates a new one if it doesn't exist
// It uses the keyring service name to store and retrieve the password
// These methods aren't exposed to the outside world, only accessible through the package main logic
func GetOrGenPasswordKeyringPass(name string) (string, error) {
	// Try to retrieve the password from the keyring
	krPass, krPassErr := krs.NewKeyringService(KeyringService, fmt.Sprintf("gobe-%s", name)).RetrievePassword()
	if krPassErr != nil && krPassErr == os.ErrNotExist {
		gl.Log("debug", fmt.Sprintf("Key found for %s", name))
		// If the error is "keyring: item not found", generate a new key
		gl.Log("debug", fmt.Sprintf("Key not found, generating new key for %s", name))
		krPassKey, krPassKeyErr := crp.NewCryptoServiceType().GenerateKey()
		if krPassKeyErr != nil {
			gl.Log("error", fmt.Sprintf("Error generating key: %v", krPassKeyErr))
			return "", krPassKeyErr
		}
		krPass = string(krPassKey)

		// Store the password in the keyring and return the encoded password
		return storeKeyringPassword(name, []byte(krPass))
	} else if krPassErr != nil {
		gl.Log("error", fmt.Sprintf("Error retrieving key: %v", krPassErr))
		return "", krPassErr
	}

	if !crp.IsBase64String(krPass) {
		krPass = crp.NewCryptoService().EncodeBase64([]byte(krPass))
	}

	return krPass, nil
}

// storeKeyringPassword stores the password in the keyring
// It will check if data is encoded, if so, will decode, store and then
// encode again or encode for the first time, returning always a portable data for
// the caller/logic outside this package be able to use it better and safer
// This method is not exposed to the outside world, only accessible through the package main logic
func storeKeyringPassword(name string, pass []byte) (string, error) {
	cryptoService := crp.NewCryptoServiceType()
	// Will decode if encoded, but only if the password is not empty, not nil and not ENCODED
	copyPass := make([]byte, len(pass))
	copy(copyPass, pass)

	var decodedPass []byte
	if crp.IsBase64String(string(copyPass)) {
		var decodeErr error
		// Will decode if encoded, but only if the password is not empty, not nil and not ENCODED
		decodedPass, decodeErr = cryptoService.DecodeIfEncoded(copyPass)
		if decodeErr != nil {
			gl.Log("error", fmt.Sprintf("Error decoding password: %v", decodeErr))
			return "", decodeErr
		}
	} else {
		decodedPass = copyPass
	}

	// Store the password in the keyring decoded to avoid storing the encoded password
	// locally are much better for security keep binary static and encoded to handle with transport
	// integration and other utilities
	storeErr := krs.NewKeyringService(KeyringService, fmt.Sprintf("gobe-%s", name)).StorePassword(string(decodedPass))
	if storeErr != nil {
		gl.Log("error", fmt.Sprintf("Error storing key: %v", storeErr))
		return "", storeErr
	}

	// Handle with logging here for getOrGenPasswordKeyringPass output
	encodedPass, encodeErr := cryptoService.EncodeIfDecoded(decodedPass)
	if encodeErr != nil {
		gl.Log("error", fmt.Sprintf("Error encoding password: %v", encodeErr))
		return "", encodeErr
	}

	// Return the encoded password to be used by the caller/logic outside this package
	return encodedPass, nil
}
