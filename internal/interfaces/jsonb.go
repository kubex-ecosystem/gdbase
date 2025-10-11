package interfaces

import "database/sql/driver"

// IJSONB é uma interface que define os métodos necessários para tipos que podem ser
// serializados e desserializados como JSONB em um banco de dados. Ela representa o contrato
// mínimo original usado pelo GORM para lidar com o tipo JSONB do PostgreSQL.
type IJSONB interface {
	Value() (driver.Value, error)
	Scan(vl any) error
}
