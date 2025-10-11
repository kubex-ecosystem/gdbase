// Package types implements various data types and structures used across the application,
// including database configurations and JSONB handling.
package types

import (
	"database/sql/driver"
	"encoding/json"
)

// JSONBData é uma interface que estende IJSONB e adiciona métodos para manipulação de dados JSONB.
type JSONBData interface {
	Scan(vl any) error
	Value() (driver.Value, error)
	ToMap() map[string]any
	ToInterface() any
	IsNil() bool
	IsEmpty() bool
	Get(key string) any
	Set(key string, value any)
	Delete(key string)
	Has(key string) bool
	Keys() []string
	Values() []any
	Len() int
	Clear()
}

// JSONBImpl é uma implementação concreta da interface JSONBData, representando um objeto JSONBImpl.
// Internamente, é um map[string]any que pode ser manipulado diretamente. Ela também implementa
// os métodos necessários para serialização e desserialização com o GORM (interface IJSONB).
type JSONBImpl map[string]any

type JSONB = JSONBImpl

// NewJSONBImpl cria uma nova instância de JSONB inicializada como um map vazio.
func NewJSONBImpl() *JSONBImpl {
	m := make(map[string]any)
	j := JSONBImpl(m)
	return &j
}

// NewJSONBData cria uma nova instância de JSONBData.
func NewJSONBData() JSONBData { return NewJSONBImpl() }

// Value is a Serializer manual para o GORM
func (m JSONBImpl) Value() (driver.Value, error) {
	return json.Marshal(m)
}

// Scan is a Deserializer manual para o GORM
func (m *JSONBImpl) Scan(vl any) error {
	if vl == nil {
		*m = JSONBImpl{}
		return nil
	}
	return json.Unmarshal(vl.([]byte), m)
}

// Métodos adicionais para manipulação de dados JSONB

func (m JSONBImpl) ToMap() map[string]any {
	if m == nil {
		return make(map[string]any)
	}
	return map[string]any(m)
}
func (m JSONBImpl) ToInterface() any {
	if m == nil {
		return map[string]any{}
	}
	return any(m)
}
func (m JSONBImpl) IsNil() bool {
	return m == nil
}
func (m JSONBImpl) IsEmpty() bool {
	return len(m) == 0
}
func (m JSONBImpl) Get(key string) any {
	if m == nil {
		return nil
	}
	return m[key]
}
func (m *JSONBImpl) Set(key string, value any) {
	if *m == nil {
		*m = make(map[string]any) // Cria o map no valor apontado
	}
	(*m)[key] = value // Acessa e modifica o map subjacente
}
func (m JSONBImpl) Delete(key string) {
	if m == nil {
		return
	}
	delete(m, key)
}
func (m JSONBImpl) Has(key string) bool {
	if m == nil {
		return false
	}
	_, ok := m[key]
	return ok
}
func (m JSONBImpl) Keys() []string {
	if m == nil {
		return []string{}
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
func (m JSONBImpl) Values() []any {
	if m == nil {
		return []any{}
	}
	values := make([]any, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
func (m JSONBImpl) Len() int {
	if m == nil {
		return 0
	}
	return len(m)
}
func (m *JSONBImpl) Clear() {
	if *m == nil {
		return
	}
	*m = make(map[string]any) // Zera o map na instância original
}

func JSONBFromMap(m map[string]any) JSONBImpl {
	return JSONBImpl(m)
}
func JSONBFromInterface(v any) JSONBImpl {
	if v == nil {
		return JSONBImpl{}
	}
	return JSONBImpl(v.(map[string]any))
}
