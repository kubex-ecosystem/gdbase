package types

import (
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	glb "github.com/rafa-mori/gdbase/internal/globals"
	ci "github.com/rafa-mori/gdbase/internal/interfaces"
	gl "github.com/rafa-mori/gdbase/internal/module/logger"
	crp "github.com/rafa-mori/gdbase/internal/security/crypto"
	krs "github.com/rafa-mori/gdbase/internal/security/external"
	t "github.com/rafa-mori/gdbase/internal/types"
	l "github.com/rafa-mori/logz"
	"gorm.io/gorm"
)

type JSONB map[string]any

// Serializer manual para o GORM
func (m JSONB) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *JSONB) Scan(vl any) error {
	if vl == nil {
		*m = JSONB{}
		return nil
	}
	return json.Unmarshal(vl.([]byte), m)
}

type DBService = IDBService

type JWT struct {
	Reference             *t.Reference `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	FilePath              string       `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	RefreshSecret         string       `gorm:"omitempty" json:"refresh_secret" yaml:"refresh_secret" xml:"refresh_secret" toml:"refresh_secret" mapstructure:"refresh_secret"`
	PrivateKey            string       `gorm:"omitempty" json:"private_key" yaml:"private_key" xml:"private_key" toml:"private_key" mapstructure:"private_key"`
	PublicKey             string       `gorm:"omitempty" json:"public_key" yaml:"public_key" xml:"public_key" toml:"public_key" mapstructure:"public_key"`
	ExpiresIn             int          `gorm:"omitempty" json:"expires_in" yaml:"expires_in" xml:"expires_in" toml:"expires_in" mapstructure:"expires_in"`
	IDExpirationSecs      int          `gorm:"omitempty" json:"id_expiration_secs" yaml:"id_expiration_secs" xml:"id_expiration_secs" toml:"id_expiration_secs" mapstructure:"id_expiration_secs"`
	RefreshExpirationSecs int          `gorm:"omitempty" json:"refresh_expiration_secs" yaml:"refresh_expiration_secs" xml:"refresh_expiration_secs" toml:"refresh_expiration_secs" mapstructure:"refresh_expiration_secs"`
	Mapper                ci.IMapper[JWT]
}
type Database struct {
	Reference        *t.Reference `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	Enabled          bool         `gorm:"default:true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`
	FilePath         string       `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	Type             string       `gorm:"not null" json:"type" yaml:"type" xml:"type" toml:"type" mapstructure:"type"`
	Driver           string       `gorm:"not null" json:"driver" yaml:"driver" xml:"driver" toml:"driver" mapstructure:"driver"`
	ConnectionString string       `gorm:"omitempty" json:"connection_string" yaml:"connection_string" xml:"connection_string" toml:"connection_string" mapstructure:"connection_string"`
	Dsn              string       `gorm:"omitempty" json:"dsn" yaml:"dsn" xml:"dsn" toml:"dsn" mapstructure:"dsn"`
	Path             string       `gorm:"omitempty" json:"path" yaml:"path" xml:"path" toml:"path" mapstructure:"path"`
	Host             string       `gorm:"omitempty" json:"host" yaml:"host" xml:"host" toml:"host" mapstructure:"host"`
	Port             interface{}  `gorm:"omitempty" json:"port" yaml:"port" xml:"port" toml:"port" mapstructure:"port"`
	Username         string       `gorm:"omitempty" json:"username" yaml:"username" xml:"username" toml:"username" mapstructure:"username"`
	Password         string       `gorm:"omitempty" json:"password" yaml:"password" xml:"password" toml:"password" mapstructure:"password"`
	Name             string       `gorm:"omitempty" json:"name" yaml:"name" xml:"name" toml:"name" mapstructure:"name"`
	Volume           string       `gorm:"omitempty" json:"volume" yaml:"volume" xml:"volume" toml:"volume" mapstructure:"volume"`
	Mapper           ci.IMapper[Database]
}
type Redis struct {
	Reference *t.Reference `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	FilePath  string       `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	Enabled   bool         `gorm:"default:true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`
	Addr      string       `gorm:"omitempty" json:"addr" yaml:"addr" xml:"addr" toml:"addr" mapstructure:"addr"`
	Port      interface{}  `gorm:"omitempty" json:"port" yaml:"port" xml:"port" toml:"port" mapstructure:"port"`
	Username  string       `gorm:"omitempty" json:"username" yaml:"username" xml:"username" toml:"username" mapstructure:"username"`
	Password  string       `gorm:"omitempty" json:"password" yaml:"password" xml:"password" toml:"password" mapstructure:"password"`
	DB        interface{}  `gorm:"omitempty" json:"db" yaml:"db" xml:"db" toml:"db" mapstructure:"db"`
	Volume    string       `gorm:"omitempty" json:"volume" yaml:"volume" xml:"volume" toml:"volume" mapstructure:"volume"`
	Mapper    ci.IMapper[Redis]
}
type RabbitMQ struct {
	Reference      *t.Reference `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	FilePath       string       `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	Enabled        bool         `gorm:"default:true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`
	Username       string       `gorm:"omitempty" json:"username" yaml:"username" xml:"username" toml:"username" mapstructure:"username"`
	Password       string       `gorm:"omitempty" json:"password" yaml:"password" xml:"password" toml:"password" mapstructure:"password"`
	Port           interface{}  `gorm:"omitempty" json:"port" yaml:"port" xml:"port" toml:"port" mapstructure:"port"`
	Host           string       `gorm:"omitempty" json:"host" yaml:"host" xml:"host" toml:"host" mapstructure:"host"`
	Volume         string       `gorm:"omitempty" json:"volume" yaml:"volume" xml:"volume" toml:"volume" mapstructure:"volume"`
	ErlangCookie   string       `gorm:"omitempty" json:"erlang_cookie" yaml:"erlang_cookie" xml:"erlang_cookie" toml:"erlang_cookie" mapstructure:"erlang_cookie"`
	ManagementUser string       `gorm:"omitempty" json:"management_user" yaml:"management_user" xml:"management_user" toml:"management_user" mapstructure:"management_user"`
	ManagementPass string       `gorm:"omitempty" json:"management_pass" yaml:"management_pass" xml:"management_pass" toml:"management_pass" mapstructure:"management_pass"`
	ManagementHost string       `gorm:"omitempty" json:"management_host" yaml:"management_host" xml:"management_host" toml:"management_host" mapstructure:"management_host"`
	ManagementPort string       `gorm:"omitempty" json:"management_port" yaml:"management_port" xml:"management_port" toml:"management_port" mapstructure:"management_port"`
	Mapper         ci.IMapper[RabbitMQ]
}
type Messagery struct {
	RabbitMQ *RabbitMQ `json:"rabbitmq" yaml:"rabbitmq" xml:"rabbitmq" toml:"rabbitmq" mapstructure:"rabbitmq"`
	Redis    *Redis    `json:"redis" yaml:"redis" xml:"redis" toml:"redis" mapstructure:"redis"`
	//*Kafka    `json:"kafka" yaml:"kafka" xml:"kafka" toml:"kafka" mapstructure:"kafka"`
	//*Nats     `json:"nats" yaml:"nats" xml:"nats" toml:"nats" mapstructure:"nats"`
	//*ActiveMQ `json:"activemq" yaml:"activemq" xml:"activemq" toml:"activemq" mapstructure:"activemq"`
	//*AMQP     `json:"amqp" yaml:"amqp" xml:"amqp" toml:"amqp" mapstructure:"amqp"`
	Mapper ci.IMapper[Messagery]
}
type MongoDB struct {
	Reference *t.Reference `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`
	FilePath  string       `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`
	Enabled   bool         `json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`
	Host      string       `json:"host" yaml:"host" xml:"host" toml:"host" mapstructure:"host"`
	Port      interface{}  `json:"port" yaml:"port" xml:"port" toml:"port" mapstructure:"port"`
	Username  string       `json:"username" yaml:"username" xml:"username" toml:"username" mapstructure:"username"`
	Password  string       `json:"password" yaml:"password" xml:"password" toml:"password" mapstructure:"password"`
	Mapper    ci.IMapper[MongoDB]
}
type DBConfig struct {
	// FilePath is used to configure the file path of the database
	FilePath string `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`

	// Logger is used to configure the logger
	Logger l.Logger `json:"logger" yaml:"logger" xml:"logger" toml:"logger" mapstructure:"logger"`

	// Mutexes is used to configure the mutexes, not serialized
	*t.Mutexes

	// JWT is used to configure the JWT token settings
	JWT JWT `json:"jwt" yaml:"jwt" xml:"jwt" toml:"jwt" mapstructure:"jwt"`

	// Reference is used to configure the reference of the database
	*t.Reference `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`

	// Enabled is used to enable or disable the database
	Enabled bool `json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`

	// MongoDB is used to configure the MongoDB database
	MongoDB MongoDB `json:"mongodb" yaml:"mongodb" xml:"mongodb" toml:"mongodb" mapstructure:"mongodb"`

	// Databases is used to configure the databases (Postgres, MySQL, SQLite, SQLServer, Oracle)
	Databases map[string]*Database `json:"databases" yaml:"databases" xml:"databases" toml:"databases" mapstructure:"databases"`

	// Messagery is used to configure the messagery database
	Messagery *Messagery

	Mapper ci.IMapper[*DBConfig]
}

func newDBConfig(name, filePath string, enabled bool, logger l.Logger, debug bool) *DBConfig {
	if logger == nil {
		logger = l.NewLogger("GodoBase")
	}

	if debug {
		gl.SetDebug(debug)
	}

	if name == "" {
		name = "default"
	}
	if filePath == "" {
		filePath = os.ExpandEnv(glb.DefaultGodoBaseConfigPath)
	}

	dbConfig := &DBConfig{FilePath: filePath}
	mapper := t.NewMapper[*DBConfig](&dbConfig, filePath)
	obj, err := mapper.DeserializeFromFile("json")
	if err != nil {
		gl.Log("warn", fmt.Sprintf("Error deserializing file: %v", err))
		if obj == nil {
			pgPass, pgPassErr := getPasswordFromKeyring(name)
			if pgPassErr != nil {
				gl.Log("error", fmt.Sprintf("Error getting password from keyring: %v", pgPassErr))
				return nil
			}
			redisPass, redisPassErr := getPasswordFromKeyring(name + "_Redis")
			if redisPassErr != nil {
				gl.Log("error", fmt.Sprintf("Error getting password from keyring: %v", redisPassErr))
				return nil
			}
			rabbitPass, rabbitPassErr := getPasswordFromKeyring(name + "_RabbitMQ")
			if rabbitPassErr != nil {
				gl.Log("error", fmt.Sprintf("Error getting password from keyring: %v", rabbitPassErr))
				return nil
			}

			dbConfigDefault := &DBConfig{
				Databases: map[string]*Database{
					"postgresql": {
						Enabled:          true,
						Reference:        t.NewReference(name).GetReference(),
						Type:             "postgresql",
						Driver:           "postgres",
						ConnectionString: fmt.Sprintf("postgres://kubex_adm:%s@localhost:5432/kubex_db", pgPass),
						Dsn:              fmt.Sprintf("postgres://kubex_adm:%s@localhost:5432/kubex_db", pgPass),
						Path:             os.ExpandEnv(glb.DefaultPostgresVolume),
						Host:             "localhost",
						Port:             "5432",
						Username:         "kubex_adm",
						Password:         pgPass,
						Volume:           os.ExpandEnv(glb.DefaultPostgresVolume),
						Name:             "kubex_db",
					},
				},
				Messagery: &Messagery{
					Redis: &Redis{
						Reference: t.NewReference(name + "_Redis").GetReference(),
						FilePath:  filePath,
						Enabled:   true,
						Addr:      "localhost",
						Port:      "6379",
						Username:  "default",
						Password:  redisPass,
						DB:        0,
					},
					RabbitMQ: &RabbitMQ{
						Reference:      t.NewReference(name + "_RabbitMQ").GetReference(),
						FilePath:       filePath,
						Enabled:        true,
						Username:       "guest",
						Password:       rabbitPass,
						Port:           "5672",
						ManagementPort: "15672",
					},
				},
			}
			dbConfig = dbConfigDefault
			mapper = t.NewMapper[*DBConfig](&dbConfig, filePath)
			mapper.SerializeToFile("json")
			if _, statErr := os.Stat(filepath.Dir(filePath)); os.IsNotExist(statErr) {
				gl.Log("fatal", fmt.Sprintf("Error creating directory: %v", statErr))
			}
			gl.Log("info", fmt.Sprintf("File %s created with default values", filePath))

			if data, dataErr := mapper.Serialize("json"); dataErr != nil {
				gl.Log("fatal", fmt.Sprintf("Error serializing file: %v", dataErr))
			} else {
				if err := os.WriteFile(filePath, data, 0644); err != nil {
					gl.Log("fatal", fmt.Sprintf("Error writing file: %v", err))
				}
			}
		}
	}
	if dbConfig.Logger == nil {
		dbConfig.Logger = logger
	}
	if dbConfig.Databases == nil {
		dbConfig.Databases = map[string]*Database{}
	}
	if dbConfig.Messagery == nil {
		dbConfig.Messagery = &Messagery{}
	}

	return dbConfig
}
func NewDBConfigWithArgs(name, filePath string, enabled bool, logger l.Logger, debug bool) *DBConfig {
	return newDBConfig(name, filePath, enabled, logger, debug)
}
func NewDBConfig(dbConfig *DBConfig) *DBConfig {
	if dbConfig.Logger == nil {
		dbConfig.Logger = l.NewLogger("GodoBase")
	}
	if dbConfig.Name == "" {
		dbConfig.Name = "default"
	}
	if dbConfig.Mutexes == nil {
		dbConfig.Mutexes = t.NewMutexesType()
	}
	if dbConfig.FilePath == "" {
		dbConfig.FilePath = os.ExpandEnv(glb.DefaultGodoBaseConfigPath)
	}
	willWrite := false
	if dbConfig.Mapper == nil {
		dbConfig.Mapper = t.NewMapper[*DBConfig](&dbConfig, dbConfig.FilePath)
		if _, statErr := os.Stat(dbConfig.FilePath); os.IsNotExist(statErr) {
			if err := os.MkdirAll(dbConfig.FilePath, os.ModePerm); err != nil {
				gl.Log("error", fmt.Sprintf("Error creating directory: %v", err))
			} else {
				willWrite = true
			}
		} else {
			_, err := dbConfig.Mapper.DeserializeFromFile("yaml")
			if err != nil {
				gl.Log("error", fmt.Sprintf("Error deserializing file: %v", err))
			}
		}
	}
	if willWrite {
		dbConfig.Mapper.SerializeToFile("json")
	}
	return dbConfig
}
func NewDBConfigWithDBConnection(db *gorm.DB) *DBConfig {
	return newDBConfig("default", "", true, nil, false)
}
func NewDBConfigWithFilePath(name, filePath string) *DBConfig {
	return newDBConfig(name, filePath, true, nil, false)
}
func getPasswordFromKeyring(name string) (string, error) {
	krPass, pgPassErr := krs.NewKeyringService(glb.KeyringService, fmt.Sprintf("gdbase-%s", name)).RetrievePassword()
	if pgPassErr != nil && pgPassErr.Error() != "keyring: item not found" {
		krPassKey, krPassKeyErr := crp.NewCryptoServiceType().GenerateKey()
		if krPassKeyErr != nil {
			gl.Log("error", fmt.Sprintf("Error generating key: %v", krPassKeyErr))
			return "", krPassKeyErr
		}
		storeErr := krs.NewKeyringService(glb.KeyringService, fmt.Sprintf("gdbase-%s", name)).StorePassword(string(krPassKey))
		if storeErr != nil {
			gl.Log("error", fmt.Sprintf("Error storing key: %v", storeErr))
			return "", storeErr
		}
		krPass = string(krPassKey)
	}
	return base64.URLEncoding.EncodeToString([]byte(krPass)), nil
}
