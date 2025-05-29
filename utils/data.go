package utils

import (
	"bytes"
	"compress/gzip"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"hash"
	"io"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

// Contains verifica se um elemento está presente em um slice, mapa ou string
// slice: coleção onde o elemento será procurado
// element: elemento a ser procurado
// Retorna true se o elemento estiver presente, caso contrário, false
func Contains(slice interface{}, element interface{}) bool {
	if element == nil || element == "" {
		return false
	}

	interfaceType := reflect.TypeOf(slice).Kind()
	elemToCompare := element.(string)

	switch interfaceType {
	case reflect.String:
		s := slice.(string)
		return strings.Contains(s, elemToCompare)
	case reflect.Int:
		s := slice.(int)
		return s == element
	case reflect.Slice:
		s := reflect.ValueOf(slice)
		for i := 0; i < s.Len(); i++ {
			if s.Index(i).String() == element {
				return true
			}
		}
		return false
	case reflect.Map:
		s := reflect.ValueOf(slice)
		for _, val := range s.MapKeys() {
			if val.String() == element {
				return true
			}
		}
		return false
	case reflect.Array:
		s := reflect.ValueOf(slice)
		for i := 0; i < s.Len(); i++ {
			if s.Index(i).String() == element {
				return true
			}
		}
		return false
	case reflect.Struct:
		s := reflect.ValueOf(slice)
		for i := 0; i < s.NumField(); i++ {
			if s.Field(i).String() == element {
				return true
			}
		}
		return false
	case reflect.Ptr:
		s := reflect.ValueOf(slice).Elem()
		for i := 0; i < s.Len(); i++ {
			if s.Index(i).String() == element {
				return true
			}
		}
		return false
	default:
		return false
	}
}

// ContainsPattern verifica se o nome do arquivo contém um dos padrões fornecidos
// filename: nome do arquivo a ser verificado
// patterns: lista de padrões a serem procurados
// Retorna true se o nome do arquivo contiver um dos padrões, caso contrário, false
func ContainsPattern(filename string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.Contains(filename, pattern) {
			return true
		}
	}
	return false
}

// EncryptData encripta os dados usando uma chave fornecida
// data: dados a serem encriptados
// key: chave usada para encriptar os dados
// Retorna os dados encriptados como string e um erro, se houver
func EncryptData(data, key string) (string, error) {
	block, err := aes.NewCipher([]byte(createHash(key)))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

// DecryptData descriptografa os dados usando uma chave fornecida
// data: dados a serem descriptografados
// key: chave usada para descriptografar os dados
// Retorna os dados descriptografados como string e um erro, se houver
func DecryptData(data, key string) (string, error) {
	block, err := aes.NewCipher([]byte(createHash(key)))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	dataBytes, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}

	nonce, ciphertext := dataBytes[:nonceSize], dataBytes[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// HashData gera um hash SHA-256 dos dados fornecidos
// data: dados a serem hasheados
// Retorna o hash dos dados como string
func HashData(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func NewHash() hash.Hash {
	return sha256.New()
}

// ValidateHash valida se os dados correspondem ao hash fornecido
// data: dados a serem validados
// hash: hash a ser comparado
// Retorna true se os dados corresponderem ao hash, caso contrário, false
func ValidateHash(data, hash string) bool {
	return HashData(data) == hash
}

// createHash cria um hash SHA-256 de uma chave fornecida
// key: chave a ser hasheada
// Retorna o hash da chave como string
func createHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	return hex.EncodeToString(hash.Sum(nil))
}

// CompressData comprime dados usando gzip
// data: dados a serem comprimidos
// Retorna os dados comprimidos como string e um erro, se houver
func CompressData(data string) (string, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	_, err := writer.Write([]byte(data))
	if err != nil {
		return "", err
	}
	writer.Close()
	return base64.URLEncoding.EncodeToString(buf.Bytes()), nil
}

// DecompressData descomprime dados previamente comprimidos
// data: dados a serem descomprimidos
// Retorna os dados descomprimidos como string e um erro, se houver
func DecompressData(data string) (string, error) {
	decoded, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	reader, err := gzip.NewReader(bytes.NewReader(decoded))
	if err != nil {
		return "", err
	}
	defer reader.Close()
	result, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func CompressFolder(folderPath string, outputPath string, compressType string) error {
	if compressType == "zip" {
		return compressFolderToZip(folderPath, outputPath)
	} else if compressType == "tar" {
		return compressFolderToTar(folderPath, outputPath)
	} else {
		return fmt.Errorf("tipo de compressão não suportado")
	}
}
func compressFolderToZip(folderPath string, outputPath string) error {
	compressCmd := fmt.Sprintf("zip -r %s %s -9", outputPath, folderPath)
	cmd := exec.Command("sh", "-c", compressCmd)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func compressFolderToTar(folderPath string, outputPath string) error {
	compressCmd := fmt.Sprintf("tar -czf %s %s", outputPath, folderPath)
	cmd := exec.Command("sh", "-c", compressCmd)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func DecompressFolder(folderPath string, outputPath string) error {
	detectedType, detectedTypeErr := detectCompressType(folderPath)
	if detectedTypeErr != nil {
		return detectedTypeErr
	}
	if detectedType == "zip" {
		return decompressFolderFromZip(folderPath, outputPath)
	} else if detectedType == "tar" {
		return decompressFolderFromTar(folderPath, outputPath)
	} else {
		return fmt.Errorf("tipo de compressão não suportado")
	}
}
func detectCompressType(folderPath string) (string, error) {
	cmdFile := fmt.Sprintf("file %s", folderPath)
	cmd := exec.Command("sh", "-c", cmdFile)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	if strings.Contains(string(output), "Zip archive data") {
		return "zip", nil
	} else if strings.Contains(string(output), "gzip compressed data") {
		return "tar", nil
	} else {
		return "", fmt.Errorf("tipo de compressão não suportado")
	}
}
func decompressFolderFromZip(folderPath string, outputPath string) error {
	decompressCmd := fmt.Sprintf("unzip %s -d %s", folderPath, outputPath)
	cmd := exec.Command("sh", "-c", decompressCmd)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func decompressFolderFromTar(folderPath string, outputPath string) error {
	decompressCmd := fmt.Sprintf("tar -xzf %s -C %s", folderPath, outputPath)
	cmd := exec.Command("sh", "-c", decompressCmd)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// EncodeData codifica dados em Base64
// data: dados a serem codificados
// Retorna os dados codificados como string
func EncodeData(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

// DecodeData decodifica dados de Base64
// data: dados a serem decodificados
// Retorna os dados decodificados como string e um erro, se houver
func DecodeData(data string) (string, error) {
	decoded, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// SignData assina dados digitalmente usando uma chave privada
// data: dados a serem assinados
// privateKey: chave privada usada para assinar os dados
// Retorna a assinatura dos dados como string e um erro, se houver
func SignData(data string, privateKey *rsa.PrivateKey) (string, error) {
	hashed := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(signature), nil
}

// VerifyData verifica a assinatura digital dos dados usando uma chave pública
// data: dados a serem verificados
// signature: assinatura a ser verificada
// publicKey: chave pública usada para verificar a assinatura
// Retorna um erro se a verificação falhar, caso contrário, nil
func VerifyData(data, signature string, publicKey *rsa.PublicKey) error {
	hashed := sha256.Sum256([]byte(data))
	decodedSig, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], decodedSig)
}

// LoadPrivateKey carrega uma chave privada de um arquivo
// path: caminho do arquivo contendo a chave privada
// Retorna a chave privada e um erro, se houver
func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("chave privada inválida")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// LoadPublicKey carrega uma chave pública de um arquivo
// path: caminho do arquivo contendo a chave pública
// Retorna a chave pública e um erro, se houver
func LoadPublicKey(path string) (any, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("chave pública inválida")
	}
	return x509.ParsePKIXPublicKey(block.Bytes)
}

// ConvertAnyDataToType converte qualquer tipo de dados para o tipo fornecido
// data: dados a serem convertidos
// targetType: tipo alvo para a conversão
// Retorna os dados convertidos e um erro, se houver
func ConvertAnyDataToType(data interface{}, targetType string) (interface{}, error) {
	if reflect.TypeOf(data).String() == targetType {
		return data, nil
	}

	switch targetType {
	case "string":
		if data == nil {
			return "", nil
		}
		return fmt.Sprintf("%v", data), nil
	case "int":
		if data == nil {
			return 0, nil
		}
		return strconv.Atoi(fmt.Sprintf("%v", data))
	case "float":
		if data == nil {
			return 0.0, nil
		}
		return strconv.ParseFloat(fmt.Sprintf("%v", data), 64)
	case "bool":
		if data == nil {
			return false, nil
		}
		return strconv.ParseBool(fmt.Sprintf("%v", data))
	case "[]byte":
		if data == nil {
			return []byte{}, nil
		}
		return []byte(fmt.Sprintf("%v", data)), nil
	default:
		return nil, fmt.Errorf("tipo de dados não suportado")
	}
}

// GetGoType retorna o tipo de dados de uma variável em Go
// data: variável a ser verificada
// Retorna o tipo de dados como string
func GetGoType(data interface{}) string {
	return reflect.TypeOf(data).String()
}

// DBTypeToGoType converte um tipo de dados de banco de dados para um tipo de dados em Go
// dbType: tipo de dados do banco de dados
// Retorna o tipo de dados correspondente em Go como string
func DBTypeToGoType(dbType string) string {
	switch dbType {
	case "NUMBER", "REAL", "DECIMAL", "NUMERIC", "FLOAT":
		return "float64"
	case "VARCHAR", "TEXT", "VARCHAR2":
		return "string"
	case "INT", "INTEGER":
		return "int"
	case "DATE", "DATETIME", "TIMESTAMP":
		return "time.Time"
	case "BOOLEAN", "BIT":
		return "bool"
	case "BLOB", "VARBINARY", "BYTEA":
		return "[]byte"
	case "CLOB":
		return "string"
	case "TINYINT":
		return "int8"
	default:
		return "interface{}"
	}
}

func LoadConfigFile(fileType, configPath string) (any, error) {
	switch fileType {
	case "ini":
		return LoadINIConfigFile(configPath)
	case "json":
		return LoadJSONConfigFile(configPath)
	case "yaml":
		return LoadYAMLConfigFile(configPath)
	case "xml":
		return LoadXMLConfigFile(configPath)
	case "toml":
		return LoadTOMLConfigFile(configPath)
	default:
		return nil, fmt.Errorf("tipo de arquivo de configuração não suportado")
	}
}

func LoadINIConfigFile(configPath string) (interface{}, error) {
	if configPath == "" {
		return nil, fmt.Errorf("caminho do arquivo de configuração não fornecido")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("arquivo de configuração não encontrado")
	}

	cfg, err := ini.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar arquivo de configuração: %v", err)
	}

	return cfg, nil
}

func LoadJSONConfigFile(configPath string) (interface{}, error) {
	if configPath == "" {
		return nil, fmt.Errorf("caminho do arquivo de configuração não fornecido")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("arquivo de configuração não encontrado")
	}
	config, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func LoadYAMLConfigFile(configPath string) (interface{}, error) {
	if configPath == "" {
		return nil, fmt.Errorf("caminho do arquivo de configuração não fornecido")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("arquivo de configuração não encontrado")
	}
	config, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func LoadXMLConfigFile(configPath string) (interface{}, error) {
	if configPath == "" {
		return nil, fmt.Errorf("caminho do arquivo de configuração não fornecido")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("arquivo de configuração não encontrado")
	}
	config, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func LoadTOMLConfigFile(configPath string) (interface{}, error) {
	if configPath == "" {
		return nil, fmt.Errorf("caminho do arquivo de configuração não fornecido")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("arquivo de configuração não encontrado")
	}
	config, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func IsBase64String(s string) bool {
	matched, _ := regexp.MatchString("^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$", s)
	return matched
}

func IsBase64ByteSlice(s []byte) bool {
	matched, _ := regexp.Match("^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$", s)
	return matched
}

func IsBase64ByteSliceString(s string) bool {
	matched, _ := regexp.Match("^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$", []byte(s))
	return matched
}
func IsBase64ByteSliceStringWithPadding(s string) bool {
	matched, _ := regexp.Match("^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$", []byte(s))
	return matched
}

func IsUrlEncodeString(s string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9%_.-]+$", s)
	return matched
}
func IsUrlEncodeByteSlice(s []byte) bool {
	matched, _ := regexp.Match("^[a-zA-Z0-9%_.-]+$", s)
	return matched
}

func IsBase62String(s string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9+_]+$", s)
	return matched
}
func IsBase62ByteSlice(s []byte) bool {
	matched, _ := regexp.Match("^[a-zA-Z0-9+_]+$", s)
	return matched
}
