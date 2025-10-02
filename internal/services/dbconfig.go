package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	gl "github.com/kubex-ecosystem/gdbase/internal/module/logger"
	crp "github.com/kubex-ecosystem/gdbase/internal/security/crypto"
	krs "github.com/kubex-ecosystem/gdbase/internal/security/external"

	ti "github.com/kubex-ecosystem/gdbase/internal/types"
	l "github.com/kubex-ecosystem/logz"
	"gorm.io/gorm"
)

const (
	KeyringService            = "kubex"
	DefaultGoBEKeyPath        = "$HOME/.kubex/gobe/gobe-key.pem"
	DefaultGoBECertPath       = "$HOME/.kubex/gobe/gobe-cert.pem"
	DefaultGodoBaseConfigPath = "$HOME/.kubex/gdbase/config/config.json"
)

const (
	DefaultVolumesDir     = "$HOME/.kubex/volumes"
	DefaultRedisVolume    = "$HOME/.kubex/volumes/redis"
	DefaultPostgresVolume = "$HOME/.kubex/volumes/postgresql"
	DefaultMongoVolume    = "$HOME/.kubex/volumes/mongo"
	DefaultRabbitMQVolume = "$HOME/.kubex/volumes/rabbitmq"
)

type IDBConfig interface {
	GetDBName() string
	GetDBType() string
	GetEnvironment() string
	GetPostgresConfig() *ti.Database
	GetMySQLConfig() *ti.Database
	GetSQLiteConfig() *ti.Database
	GetMongoDBConfig() *ti.MongoDB
	GetRedisConfig() *ti.Redis
	GetRabbitMQConfig() *ti.RabbitMQ
	IsAutoMigrate() bool
	IsDebug() bool
	GetLogger() any
	GetConfig(context.Context) *DBConfig
	GetConfigMap(context.Context) map[string]any
}

type DBConfig struct {
	// Name is used to configure the name of the database
	Name string `json:"name" yaml:"name" xml:"name" toml:"name" mapstructure:"name"`

	// FilePath is used to configure the file path of the database
	FilePath string `json:"file_path" yaml:"file_path" xml:"file_path" toml:"file_path" mapstructure:"file_path"`

	// Logger is used to configure the logger
	Logger l.Logger `json:"logger" yaml:"logger" xml:"logger" toml:"logger" mapstructure:"logger"`

	// Mutexes is used to configure the mutexes, not serialized
	*ti.Mutexes

	// Properties is used to configure the properties of the database, not serialized
	properties map[string]interface{} `json:"-" yaml:"-" xml:"-" toml:"-" mapstructure:"-"`

	// Debug is used to configure the debug mode
	Debug bool `json:"debug" yaml:"debug" xml:"debug" toml:"debug" mapstructure:"debug"`

	// AutoMigrate is used to configure the auto migration of the database
	AutoMigrate bool `json:"auto_migrate" yaml:"auto_migrate" xml:"auto_migrate" toml:"auto_migrate" mapstructure:"auto_migrate"`

	// JWT is used to configure the JWT token settings
	JWT ti.JWT `json:"jwt" yaml:"jwt" xml:"jwt" toml:"jwt" mapstructure:"jwt"`

	// Reference is used to configure the reference of the database
	*ti.Reference `json:"reference" yaml:"reference" xml:"reference" toml:"reference" mapstructure:"reference,squash"`

	// Enabled is used to enable or disable the database
	Enabled bool `json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled" mapstructure:"enabled"`

	// MongoDB is used to configure the MongoDB database
	MongoDB ti.MongoDB `json:"mongodb" yaml:"mongodb" xml:"mongodb" toml:"mongodb" mapstructure:"mongodb"`

	// Databases is used to configure the databases (Postgres, MySQL, SQLite, SQLServer, Oracle)
	Databases map[string]*ti.Database `json:"databases" yaml:"databases" xml:"databases" toml:"databases" mapstructure:"databases"`

	// Messagery is used to configure the messagery database
	Messagery *ti.Messagery

	Mapper *ti.Mapper[*DBConfig] `json:"-" yaml:"-" xml:"-" toml:"-" mapstructure:"-"`
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
		filePath = os.ExpandEnv(DefaultGodoBaseConfigPath)
	}

	dbConfig := &DBConfig{FilePath: filePath}
	mapper := ti.NewMapper[*DBConfig](&dbConfig, filePath)
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
				Databases: map[string]*ti.Database{
					"postgresql": {
						Enabled:          true,
						Reference:        ti.NewReference(name).GetReference(),
						Type:             "postgresql",
						Driver:           "postgres",
						ConnectionString: fmt.Sprintf("postgres://kubex_adm:%s@localhost:5432/kubex_db", pgPass),
						Dsn:              fmt.Sprintf("postgres://kubex_adm:%s@localhost:5432/kubex_db", pgPass),
						Path:             os.ExpandEnv(DefaultPostgresVolume),
						Host:             "localhost",
						Port:             "5432",
						Username:         "kubex_adm",
						Password:         pgPass,
						Volume:           os.ExpandEnv(DefaultPostgresVolume),
						Name:             "kubex_db",
					},
				},
				Messagery: &ti.Messagery{
					Redis: &ti.Redis{
						Reference: ti.NewReference(name + "_Redis").GetReference(),
						FilePath:  filePath,
						Enabled:   true,
						Addr:      "localhost",
						Port:      "6379",
						Username:  "default",
						Password:  redisPass,
						DB:        0,
					},
					RabbitMQ: &ti.RabbitMQ{
						Reference:      ti.NewReference(name + "_RabbitMQ").GetReference(),
						FilePath:       filePath,
						Enabled:        true,
						Username:       "gobe",
						Password:       rabbitPass,
						Port:           "5672",
						ManagementPort: "15672",
						Vhost:          "gobe",
					},
				},
			}
			dbConfig = dbConfigDefault
			mapper = ti.NewMapper(&dbConfig, filePath)
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
		dbConfig.Databases = map[string]*ti.Database{}
	}
	if dbConfig.Messagery == nil {
		dbConfig.Messagery = &ti.Messagery{}
	}

	return dbConfig
}
func NewDBConfigWithArgs(ctx context.Context, name, filePath string, enabled bool, logger l.Logger, debug bool) *DBConfig {
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
		dbConfig.Mutexes = ti.NewMutexesType()
	}
	if dbConfig.FilePath == "" {
		dbConfig.FilePath = os.ExpandEnv(DefaultGodoBaseConfigPath)
	}
	willWrite := false
	if dbConfig.Mapper == nil {
		dbConfig.Mapper = ti.NewMapperType(&dbConfig, dbConfig.FilePath)
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
func NewDBConfigFromFile(ctx context.Context, dbConfigFilePath string, autoMigrate bool, logger l.Logger, debug bool) (*DBConfig, error) {
	var dbConfig *DBConfig
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return dbConfig, nil
}

func getPasswordFromKeyring(name string) (string, error) {
	krPass, pgPassErr := krs.NewKeyringService(KeyringService, fmt.Sprintf("gdbase-%s", name)).RetrievePassword()
	if pgPassErr != nil && pgPassErr.Error() != "keyring: item not found" {
		krPassKey, krPassKeyErr := crp.NewCryptoServiceType().GenerateKey()
		if krPassKeyErr != nil {
			gl.Log("error", fmt.Sprintf("Error generating key: %v", krPassKeyErr))
			return "", krPassKeyErr
		}
		storeErr := krs.NewKeyringService(KeyringService, fmt.Sprintf("gdbase-%s", name)).StorePassword(string(krPassKey))
		if storeErr != nil {
			gl.Log("error", fmt.Sprintf("Error storing key: %v", storeErr))
			return "", storeErr
		}
		krPass = string(krPassKey)
	}
	return base64.URLEncoding.EncodeToString([]byte(krPass)), nil
}

func GetConnectionString(dbConfig *ti.Database) string {
	if dbConfig.ConnectionString != "" {
		return dbConfig.ConnectionString
	}
	if dbConfig.Host != "" && dbConfig.Port != nil && dbConfig.Username != "" && dbConfig.Name != "" {
		dbPass := dbConfig.Password
		if dbPass == "" {
			dbPassKey, dbPassErr := getPasswordFromKeyring("pgpass")
			if dbPassErr != nil {
				gl.Log("error", fmt.Sprintf("❌ Erro ao recuperar senha do banco de dados: %v", dbPassErr))
			} else {
				dbConfig.Password = string(dbPassKey)
				dbPass = dbConfig.Password
			}
		}
		return fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo",
			// "host=%s port=%s user=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo",
			dbConfig.Host, dbConfig.Port.(string), dbConfig.Username, dbPass, dbConfig.Name,
			// dbConfig.Host, dbConfig.Port.(string), dbConfig.Username /* dbPass, */, dbConfig.Name,
		)
	}
	return ""
}

func (d *DBConfig) GetDBName() string {
	if d == nil {
		return ""
	}
	for _, db := range d.Databases {
		if db.Enabled {
			return db.Name
		}
	}
	return ""
}
func (d *DBConfig) GetEnvironment() string {
	if d == nil {
		return ""
	}
	// config, ok := d.properties["config"].(*ti.Property[*DBConfig])
	// if !ok {
	// 	return ""
	// }
	// dbConfig := config.GetValue()
	// if dbConfig == nil {
	// 	return ""
	// }
	return os.Getenv("GO_ENV")
}
func (d *DBConfig) GetDBType() string {
	if d == nil {
		return ""
	}
	for _, db := range d.Databases {
		if db.Enabled {
			return db.Type
		}
	}
	return ""
}
func (d *DBConfig) GetPostgresConfig() *ti.Database {
	if d == nil {
		return nil
	}
	postgres, ok := d.Databases["postgresql"]
	if !ok {
		return nil
	}
	return postgres
}
func (d *DBConfig) GetMySQLConfig() *ti.Database {
	if d == nil {
		return nil
	}
	mysql, ok := d.Databases["mysql"]
	if !ok {
		return nil
	}
	return mysql
}

func (d *DBConfig) GetSQLiteConfig() *ti.Database {
	if d == nil {
		return nil
	}
	sqlite, ok := d.Databases["sqlite"]
	if !ok {
		return nil
	}
	return sqlite
}
func (d *DBConfig) GetMongoDBConfig() *ti.MongoDB {
	if d == nil {
		return nil
	}
	return &d.MongoDB
}
func (d *DBConfig) GetRedisConfig() *ti.Redis {
	if d == nil {
		return nil
	}
	return d.Messagery.Redis
}
func (d *DBConfig) GetRabbitMQConfig() *ti.RabbitMQ {
	if d == nil {
		return nil
	}
	return d.Messagery.RabbitMQ
}
func (d *DBConfig) IsAutoMigrate() bool {
	if d == nil {
		return false
	}

	return d.Enabled
}
func (d *DBConfig) IsDebug() bool {
	if d == nil {
		return false
	}
	config, ok := d.properties["config"].(*ti.Property[*DBConfig])
	if !ok {
		return false
	}
	dbConfig := config.GetValue()
	if dbConfig == nil {
		return false
	}
	return dbConfig.Debug
}
func (d *DBConfig) GetLogger() interface{} {
	if d == nil {
		return nil
	}
	return d.Logger
}

func (d *DBConfig) GetConfig(ctx context.Context) *DBConfig {
	if d == nil {
		return nil
	}
	return d
}

func (d *DBConfig) GetConfigMap(ctx context.Context) map[string]any {
	// Quero percorrer toda a struct de DBConfig independente da profundidade e das propriedades
	// Criar um map disso e retornar
	if d == nil {
		return nil
	}
	arrForConfig := make(map[string]any)

	for key, val := range d.properties {
		arrForConfig[key] = val
	}

	v := reflect.ValueOf(*d)
	reflect.VisibleFields(v.Type())
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i).Interface()
		if field.Type.Kind() == reflect.Ptr && !v.Field(i).IsNil() {
			elem := v.Field(i).Elem()
			if elem.Kind() == reflect.Struct {
				nestedConfig := make(map[string]any)
				for j := 0; j < elem.NumField(); j++ {
					nestedField := elem.Type().Field(j)
					nestedValue := elem.Field(j).Interface()
					nestedConfig[nestedField.Name] = nestedValue
				}
				arrForConfig[field.Name] = nestedConfig
			} else {
				arrForConfig[field.Name] = value
			}
		} else if field.Type.Kind() == reflect.Struct {
			nestedConfig := make(map[string]any)
			elem := v.Field(i)
			for j := 0; j < elem.NumField(); j++ {
				nestedField := elem.Type().Field(j)
				nestedValue := elem.Field(j).Interface()
				nestedConfig[nestedField.Name] = nestedValue
			}
			arrForConfig[field.Name] = nestedConfig
		} else {
			arrForConfig[field.Name] = value
		}
	}

	return arrForConfig
}
