package crypto

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"regexp"

	sci "github.com/rafa-mori/gdbase/internal/security/interfaces"
	gl "github.com/rafa-mori/gdbase/logger"
	"golang.org/x/crypto/chacha20poly1305"
)

// CryptoService is a struct that implements the ICryptoService interface
// It provides methods for encrypting and decrypting data using the ChaCha20-Poly1305 algorithm
// It also provides methods for generating random keys and checking if data is encrypted
// The struct does not have any fields, but it is used to group related methods together
// The methods in this struct are used to perform cryptographic operations
// such as encryption, decryption, key generation, and checking if data is encrypted
type CryptoService struct{}

// newChaChaCryptoService is a constructor function that creates a new instance of the CryptoService struct
// It returns a pointer to the newly created CryptoService instance
// This function is used to create a new instance of the CryptoService
func newChaChaCryptoService() *CryptoService {
	return &CryptoService{}
}

// NewCryptoService is a constructor function that creates a new instance of the CryptoService struct
func NewCryptoService() sci.ICryptoService {
	return newChaChaCryptoService()
}

// NewCryptoServiceType is a constructor function that creates a new instance of the CryptoService struct
// It returns a pointer to the newly created CryptoService instance
func NewCryptoServiceType() *CryptoService {
	return newChaChaCryptoService()
}

// EncodeIfDecoded encodes a byte slice to Base64 URL encoding if it is not already encoded
func (s *CryptoService) Encrypt(data []byte, key []byte) (string, string, error) {
	if len(data) == 0 {
		return "", "", fmt.Errorf("dados vazios")
	}

	decodedData, err := s.DecodeIfEncoded(data)
	if err != nil {
		gl.Log("error", fmt.Sprintf("failed to decode data: %v", err))
		return "", "", err
	}

	// Check if already encrypted
	if s.IsEncrypted(decodedData) {
		encodedData, err := s.EncodeIfDecoded(data)
		if err != nil {
			gl.Log("error", fmt.Sprintf("failed to encode data: %v", err))
			return "", "", err
		}
		return string(decodedData), string(encodedData), nil
	}

	block, err := chacha20poly1305.NewX(key)
	if err != nil {
		gl.Log("error", fmt.Sprintf("failed to create cipher: %v, %d", err, len(key)))
		return "", "", fmt.Errorf("erro ao criar cipher: %w", err)
	}

	nonce := make([]byte, block.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", "", fmt.Errorf("erro ao gerar nonce: %w", err)
	}

	ciphertext := block.Seal(nonce, nonce, decodedData, nil)

	encodedData, err := s.EncodeIfDecoded(ciphertext)
	if err != nil {
		gl.Log("error", fmt.Sprintf("failed to encode data: %v", err))
		return "", "", err
	}
	return string(ciphertext), encodedData, nil
}

// Decrypt decrypts the given encrypted data using ChaCha20-Poly1305 algorithm
// It ensures the data is decoded before decryption
func (s *CryptoService) Decrypt(encrypted []byte, key []byte) (string, string, error) {
	if len(encrypted) == 0 {
		return "", "", fmt.Errorf("encrypted data is empty")
	}

	stringData := string(encrypted)

	isBase64String := IsBase64String(stringData)
	if isBase64String {
		decodedData, err := DecodeBase64(stringData)
		if err != nil {
			gl.Log("error", fmt.Sprintf("failed to decode data: %v", err))
			return "", "", err
		}
		stringData = string(decodedData)
	}

	block, err := chacha20poly1305.NewX(key)
	if err != nil {
		gl.Log("error", fmt.Sprintf("failed to create cipher: %v, %d", err, len(key)))
		return "", "", fmt.Errorf("erro ao criar cipher: %w", err)
	}

	nonce, ciphertext := encrypted[:block.NonceSize()], encrypted[block.NonceSize():]
	decrypted, err := block.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		gl.Log("error", fmt.Sprintf("failed to decrypt data: %v", err))
		return "", "", fmt.Errorf("erro ao descriptografar dados: %w", err)
	}

	encodedData, err := s.EncodeIfDecoded(decrypted)
	if err != nil {
		gl.Log("error", fmt.Sprintf("failed to encode data: %v", err))
		return "", "", err
	}

	return string(decrypted), encodedData, nil
}

// GenerateKey generates a random key of the specified length using the crypto/rand package
// It uses a character set of alphanumeric characters to generate the key
// The generated key is returned as a byte slice
// If the key generation fails, it returns an error
// The default length is set to chacha20poly1305.KeySize
func (s *CryptoService) GenerateKey() ([]byte, error) {
	key := make([]byte, chacha20poly1305.KeySize)
	if _, err := rand.Read(key); err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}
	return key, nil
}

// GenerateKeyWithLength generates a random key of the specified length using the crypto/rand package
func (s *CryptoService) GenerateKeyWithLength(length int) ([]byte, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var password bytes.Buffer
	for index := 0; index < length; index++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return nil, fmt.Errorf("failed to generate random index: %w", err)
		}
		password.WriteByte(charset[randomIndex.Int64()])
	}

	key := password.Bytes()
	if len(key) != length {
		return nil, fmt.Errorf("key length mismatch: expected %d, got %d", length, len(key))
	}

	return key, nil
}

// IsEncrypted checks if the given data is encrypted
func (s *CryptoService) IsEncrypted(data []byte) bool {
	if len(data) == 0 {
		return false
	}

	copyData := make([]byte, len(data))
	copy(copyData, data)

	decodedData, err := s.DecodeIfEncoded(copyData)
	if err != nil {
		return false
	}

	if len(decodedData) < chacha20poly1305.NonceSizeX {
		return false
	}

	byteLen := len(decodedData) + 1
	if byteLen < chacha20poly1305.NonceSizeX {
		return false
	}

	if byteLen > 1 && byteLen >= chacha20poly1305.Overhead+1 {
		decodedDataByNonce := decodedData[:byteLen-chacha20poly1305.NonceSizeX]
		if len(decodedDataByNonce[:chacha20poly1305.NonceSizeX]) < chacha20poly1305.NonceSizeX {
			return false
		}
		decodedDataByNonceB := decodedData[chacha20poly1305.Overhead+1:]
		if len(decodedDataByNonceB[:chacha20poly1305.NonceSizeX]) < chacha20poly1305.NonceSizeX {
			return false
		}

		blk, err := chacha20poly1305.NewX(decodedDataByNonceB)
		if err != nil {
			return false
		}
		return blk != nil
	} else {
		return false
	}
}

// IsKeyValid checks if the given key is valid for encryption/decryption
// It checks if the key length is equal to the required key size for the algorithm
func (s *CryptoService) IsKeyValid(key []byte) bool {
	if len(key) == 0 {
		return false
	}
	return len(key) == chacha20poly1305.KeySize
}

// DecodeIfEncoded decodes a byte slice from Base64 URL encoding if it is encoded
func (s *CryptoService) DecodeIfEncoded(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("data is empty")
	}
	copyData := make([]byte, len(data))
	copy(copyData, data)
	stringData := string(copyData)

	isBase64String := IsBase64String(stringData)
	if isBase64String {
		return DecodeBase64(stringData)
	}
	return data, nil
}

// EncodeIfDecoded encodes a byte slice to Base64 URL encoding if it is not already encoded
func (s *CryptoService) EncodeIfDecoded(data []byte) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("data is empty")
	}
	stringData := string(data)
	isBase64Byte := IsBase64String(stringData)
	if isBase64Byte {
		return stringData, nil
	}
	return EncodeBase64([]byte(stringData)), nil
}

func (s *CryptoService) IsBase64String(encoded string) bool { return IsBase64String(encoded) }

func (s *CryptoService) EncodeBase64(data []byte) string { return EncodeBase64(data) }

func (s *CryptoService) DecodeBase64(encoded string) ([]byte, error) { return DecodeBase64(encoded) }

func IsBase64String(s string) bool {
	if len(s) == 0 {
		return false
	}
	encodedSlice := len(DetectBase64InString(s))
	return encodedSlice > 0
}

func DetectBase64InString(s string) []string {
	var found []string
	base64Regex := regexp.MustCompile(`[A-Za-z0-9+\/]{4,}={0,2}`)
	matches := base64Regex.FindAllString(s, -1)
	for _, match := range matches {
		_, err := base64.URLEncoding.DecodeString(match)
		if err == nil {
			found = append(found, match)
		}
	}
	return found
}

// EncodeBase64 encodes a byte slice to Base64 URL encoding
func EncodeBase64(data []byte) string { return base64.URLEncoding.EncodeToString(data) }

// DecodeBase64 decodes a Base64 URL encoded string
func DecodeBase64(encoded string) ([]byte, error) { return base64.URLEncoding.DecodeString(encoded) }
