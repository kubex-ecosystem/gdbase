package services

import (
	"fmt"

	gbm "github.com/rafa-mori/gdbase"

	// gl "github.com/rafa-mori/gdbase/logger"
	"sync"

	glb "github.com/rafa-mori/gdbase/internal/globals"
	gl "github.com/rafa-mori/gdbase/logger"
	t "github.com/rafa-mori/gdbase/types"
	l "github.com/rafa-mori/logz"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBService struct {
	Logger    l.Logger
	reference gbm.Reference
	mutexes   gbm.Mutexes

	db   *gorm.DB
	pool *sync.Pool

	// properties are used to store database settings and configurations
	properties map[string]any
}

func NewDatabaseService(config *t.DBConfig, logger l.Logger) (*DBService, error) {
	if logger == nil {
		logger = l.GetLogger("GDBase")
	}

	if config == nil {
		return nil, fmt.Errorf("❌ Configuração do banco de dados não pode ser nula")
	}
	if len(config.Databases) == 0 {
		return nil, fmt.Errorf("❌ Configuração de banco de dados não pode ser vazia")
	}

	//driver = db.Driver // Pro futuro.. rs
	var dbHost, dbPort, dbUser, dbPass, dbName, dsn string
	for _, dbConfig := range config.Databases {
		if dsn == "" {
			dsn = dbConfig.ConnectionString
		}
		if dsn == "" {
			dbHost = dbConfig.Host
			dbPort = dbConfig.Port.(string)
			dbUser = dbConfig.Username
			if dbConfig.Type != "postgresql" {
				dbPass = dbConfig.Password
				if dbPass == "" {
					dbPassKey, dbPassErr := glb.GetOrGenPasswordKeyringPass("pgpass")
					if dbPassErr != nil {
						gl.Log("error", fmt.Sprintf("❌ Erro ao recuperar senha do banco de dados: %v", dbPassErr))
						continue
					}
					dbConfig.Password = string(dbPassKey)
					dbPass = dbConfig.Password
				}

				dbName = dbConfig.Name
				dsn = fmt.Sprintf(
					"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo",
					dbHost, dbPort, dbUser, dbPass, dbName,
				)
			} else {
				// dbPass = dbConfig.Password
				dbName = dbConfig.Name
				dsn = fmt.Sprintf(
					"host=%s port=%s user=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo",
					dbHost, dbPort, dbUser, dbName,
				)
			}
			break
		}
	}

	dbService := &DBService{
		Logger:     logger,
		reference:  gbm.NewReference("DBService"),
		mutexes:    gbm.NewMutexesType(),
		properties: make(map[string]any),
		pool:       &sync.Pool{},
	}

	dbService.properties["config"] = gbm.NewProperty[*t.DBConfig]("config", &config, true, nil)

	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		return nil, fmt.Errorf("❌ Erro ao conectar ao banco de dados: %v", err)
	} else {
		dbService.db = db
	}

	return dbService, nil
}

func (d *DBService) Initialize() error {
	if d.db != nil {
		return nil
	}
	envT := d.properties["config"].(gbm.Property[gbm.Environment])
	if envT == nil {
		return fmt.Errorf("❌ Erro ao recuperar o ambiente")
	}
	env := envT.GetValue()
	if env == nil {
		return fmt.Errorf("❌ Erro ao recuperar o ambiente")
	}

	dbType := env.Getenv("DB_TYPE")
	dbHost := env.Getenv("DB_HOST")
	dbPort := env.Getenv("DB_PORT")
	dbUser := env.Getenv("DB_USER")
	dbPass := env.Getenv("DB_PASS")
	dbName := env.Getenv("DB_NAME")
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)
	db, err := ConnectDatabase(dbType, dsn)
	if err != nil {
		return fmt.Errorf("❌ Erro ao conectar ao banco de dados: %v", err)
	}
	d.db = db
	return nil
}

func (d *DBService) GetDB() (*gorm.DB, error) {
	if d.db == nil {
		return nil, fmt.Errorf("❌ Banco de dados não inicializado")
	}
	return d.db, nil
}

func (d *DBService) CloseDBConnection() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return fmt.Errorf("❌ Erro ao obter conexão SQL: %v", err)
	}
	return sqlDB.Close()
}

func (d *DBService) CheckDatabaseHealth() error {
	if err := d.db.Raw("SELECT 1").Error; err != nil {
		return fmt.Errorf("❌ Banco de dados offline: %v", err)
	}
	return nil
}

func (d *DBService) IsConnected() error {
	if d.db == nil {
		return fmt.Errorf("❌ Banco de dados não inicializado")
	}
	if err := d.db.Raw("SELECT 1").Error; err != nil {
		return fmt.Errorf("❌ Banco de dados offline: %v", err)
	}
	return nil
}

func (d *DBService) Reconnect() error {
	if d.db != nil {
		sqlDB, err := d.db.DB()
		if err != nil {
			return fmt.Errorf("❌ Erro ao obter conexão SQL: %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			return fmt.Errorf("❌ Erro ao fechar conexão SQL: %v", err)
		}
	}

	db, err := d.GetDB()
	if err != nil {
		return fmt.Errorf("❌ Erro ao obter banco de dados: %v", err)
	}

	d.db = db
	return nil
}

func (d *DBService) GetHost() (string, error) {
	if d.db == nil {
		return "", fmt.Errorf("❌ Banco de dados não inicializado")
	}
	host, ok := d.properties["host"].(gbm.Property[string])
	if !ok {
		return "", fmt.Errorf("❌ Erro ao obter host do banco de dados")
	}
	vl := host.GetValue()
	if vl == "" {
		return "", fmt.Errorf("❌ Host do banco de dados não encontrado")
	}
	return vl, nil
}

func (d *DBService) GetConfig() *t.DBConfig {
	if d.db == nil {
		return nil
	}
	config, ok := d.properties["config"].(gbm.Property[*t.DBConfig])
	if !ok {
		return nil
	}
	return config.GetValue()
}

func ConnectDatabase(dbType, dsn string) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch dbType {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(dsn)
	default:
		return nil, fmt.Errorf("banco de dados não suportado: %s", dbType)
	}

	return gorm.Open(dialector, &gorm.Config{})
}
