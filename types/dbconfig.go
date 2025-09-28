package types

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	glb "github.com/kubex-ecosystem/gdbase/internal/globals"
	gl "github.com/kubex-ecosystem/gdbase/internal/module/logger"
	crp "github.com/kubex-ecosystem/gdbase/internal/security/crypto"
	krs "github.com/kubex-ecosystem/gdbase/internal/security/external"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
	l "github.com/kubex-ecosystem/logz"
	"gorm.io/gorm"
)

type DBConfig struct {
	// Name is used to configure the name of the database
	Name string `json:"name" yaml:"name" xml:"name" toml:"name" mapstructure:"name"`

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

	Mapper *t.Mapper[*DBConfig] `json:"-" yaml:"-" xml:"-" toml:"-" mapstructure:"-"`
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
						Username:       "gobe",
						Password:       rabbitPass,
						Port:           "5672",
						ManagementPort: "15672",
						Vhost:          "gobe",
					},
				},
			}
			dbConfig = dbConfigDefault
			mapper = t.NewMapper(&dbConfig, filePath)
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
		dbConfig.Mapper = t.NewMapperType(&dbConfig, dbConfig.FilePath)
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
