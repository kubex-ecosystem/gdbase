package services

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	// gl "github.com/kubex-ecosystem/gdbase/logger"
	"sync"

	crp "github.com/kubex-ecosystem/gdbase/internal/security/crypto"
	krs "github.com/kubex-ecosystem/gdbase/internal/security/external"

	glb "github.com/kubex-ecosystem/gdbase/internal/globals"
	ci "github.com/kubex-ecosystem/gdbase/internal/interfaces"
	gl "github.com/kubex-ecosystem/gdbase/internal/module/logger"
	ti "github.com/kubex-ecosystem/gdbase/internal/types"
	l "github.com/kubex-ecosystem/logz"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DBService struct {
	Logger    l.Logger
	reference ci.IReference
	mutexes   ci.IMutexes

	db   *gorm.DB
	pool *sync.Pool

	// properties are used to store database settings and configurations
	properties map[string]any
}

func newDatabaseService(ctx context.Context, config *DBConfig, logger l.Logger) (*DBService, error) {
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
			dbConfig.ConnectionString = dsn
			break
		}
	}

	dbService := &DBService{
		Logger:     logger,
		reference:  ti.NewReference("DBService"),
		mutexes:    ti.NewMutexesType(),
		properties: make(map[string]any),
		pool:       &sync.Pool{},
	}

	dbService.properties["config"] = ti.NewProperty("config", &config, true, nil)

	if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		return nil, fmt.Errorf("❌ Erro ao conectar ao banco de dados: %v", err)
	} else {
		dbService.db = db
	}

	return dbService, nil
}

func NewDatabaseService(ctx context.Context, config *DBConfig, logger l.Logger) (ci.IDBService, error) {
	return newDatabaseService(ctx, config, logger)
}

func (d *DBService) Initialize(ctx context.Context) error {
	if d.db != nil {
		return nil
	}
	cfgT := d.properties["config"].(*ti.Property[*DBConfig])
	if cfgT == nil {
		return fmt.Errorf("❌ Erro ao recuperar a configuração do banco de dados")
	}
	config := cfgT.GetValue()
	if config == nil {
		return fmt.Errorf("❌ Erro ao recuperar a configuração do banco de dados")
	}

	for _, dbConfig := range config.Databases {
		if dbConfig.Type == "" || !dbConfig.Enabled {
			continue
		}
		dsn := GetConnectionString(dbConfig)
		if dsn == "" {
			continue
		}
		if db, err := connectDatabase(context.Background(), dbConfig.Type, dsn); err != nil {
			return fmt.Errorf("❌ Erro ao conectar ao banco de dados: %v", err)
		} else {
			d.db = db
			break
		}
	}

	if d.db == nil {
		return fmt.Errorf("❌ Erro ao conectar ao banco de dados: configuração inválida")
	}

	return nil
}

func (d *DBService) InitializeFromEnv(ctx context.Context, env ci.IEnvironment) error {
	if d.db != nil {
		return nil
	}
	if env == nil {
		return fmt.Errorf("❌ Serviço de ambiente não pode ser nulo")
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
	databaseConfig := &ti.Database{
		Type:             dbType,
		Host:             dbHost,
		Port:             dbPort,
		Username:         dbUser,
		Password:         dbPass,
		Name:             dbName,
		ConnectionString: dsn,
		Enabled:          true,
	}
	dbConfig := d.properties["config"].(*ti.Property[*DBConfig]).GetValue()

	if dbConfig != nil {
		if _, exists := dbConfig.Databases[databaseConfig.Name]; exists {
			gl.Log("info", fmt.Sprintf("Configuração do banco de dados '%s' já existe, pulando criação", databaseConfig.Name))
			return nil
		}
	}
	if dbConfig == nil {
		dbConfig := &DBConfig{
			Databases: map[string]*ti.Database{
				databaseConfig.Name: databaseConfig,
			},
		}
		d.properties["config"] = ti.NewProperty("config", &dbConfig, true, nil)
	}
	// Aguarda o banco de dados ficar pronto e conecta
	// Timeout de 1 minuto para aguardar o banco de dados ficar pronto
	db, conn, err := waitAndConnect(context.Background(), databaseConfig, 1*time.Minute)
	if err != nil {
		return fmt.Errorf("❌ Erro ao conectar ao banco de dados: %v", err)
	}
	defer conn.Close()
	d.db = db
	return nil
}

func (d *DBService) GetDB(ctx context.Context) (*gorm.DB, error) {
	if d.db == nil {
		return nil, fmt.Errorf("❌ Banco de dados não inicializado")
	}
	return d.db, nil
}

func (d *DBService) CloseDBConnection(ctx context.Context) error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return fmt.Errorf("❌ Erro ao obter conexão SQL: %v", err)
	}
	return sqlDB.Close()
}

func (d *DBService) CheckDatabaseHealth(ctx context.Context) error {
	if err := d.db.Raw("SELECT 1").Error; err != nil {
		return fmt.Errorf("❌ Banco de dados offline: %v", err)
	}
	return nil
}

func (d *DBService) GetConnection(ctx context.Context, database string, timeout time.Duration) (*sql.Conn, error) {
	if d.db == nil {
		return nil, fmt.Errorf("❌ Banco de dados não inicializado")
	}
	// Pega a configuração do banco de dados
	cfgT := d.properties["config"].(*ti.Property[*DBConfig])
	if cfgT == nil {
		return nil, fmt.Errorf("❌ Erro ao recuperar a configuração do banco de dados")
	}
	config := cfgT.GetValue()
	if config == nil {
		return nil, fmt.Errorf("❌ Erro ao recuperar a configuração do banco de dados")
	}
	var dbConfig *ti.Database
	for _, dbConf := range config.Databases {
		if dbConf.Name == database && dbConf.Enabled {
			dbConfig = dbConf
			break
		}
	}
	if dbConfig == nil {
		return nil, fmt.Errorf("❌ Configuração do banco de dados '%s' não encontrada ou desabilitada", database)
	}

	// Aguarda o banco de dados ficar pronto e retorna a conexão
	// Timeout de 1 minuto para aguardar o banco de dados ficar pronto
	if timeout <= 0 {
		timeout = 1 * time.Minute
	}
	_, conn, err := waitAndConnect(context.Background(), dbConfig, timeout)
	if err != nil {
		return nil, fmt.Errorf("❌ Erro ao conectar ao banco de dados '%s': %v", database, err)
	}
	return conn, nil
}

func (d *DBService) IsConnected(ctx context.Context) error {
	if d.db == nil {
		return fmt.Errorf("❌ Banco de dados não inicializado")
	}
	if err := d.db.Raw("SELECT 1").Error; err != nil {
		return fmt.Errorf("❌ Banco de dados offline: %v", err)
	}
	return nil
}

func (d *DBService) Reconnect(ctx context.Context) error {
	if d.db != nil {
		sqlDB, err := d.db.DB()
		if err != nil {
			return fmt.Errorf("❌ Erro ao obter conexão SQL: %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			return fmt.Errorf("❌ Erro ao fechar conexão SQL: %v", err)
		}
	}

	db, err := d.GetDB(ctx)
	if err != nil {
		return fmt.Errorf("❌ Erro ao obter banco de dados: %v", err)
	}

	d.db = db
	return nil
}

func (d *DBService) GetName(ctx context.Context) (string, error) {
	if d.db == nil {
		return "", fmt.Errorf("❌ Banco de dados não inicializado")
	}
	name, ok := d.properties["name"].(*ti.Property[string])
	if !ok {
		return "", fmt.Errorf("❌ Erro ao obter nome do banco de dados")
	}
	vl := name.GetValue()
	if vl == "" {
		return "", fmt.Errorf("❌ Nome do banco de dados não encontrado")
	}
	return vl, nil
}

func (d *DBService) GetHost(ctx context.Context) (string, error) {
	if d.db == nil {
		return "", fmt.Errorf("❌ Banco de dados não inicializado")
	}
	host, ok := d.properties["host"].(*ti.Property[string])
	if !ok {
		return "", fmt.Errorf("❌ Erro ao obter host do banco de dados")
	}
	vl := host.GetValue()
	if vl == "" {
		return "", fmt.Errorf("❌ Host do banco de dados não encontrado")
	}
	return vl, nil
}

func (d *DBService) GetConfig(ctx context.Context) map[string]any {
	if d.db == nil {
		return nil
	}
	cfgT := d.properties["config"].(*ti.Property[*DBConfig])
	if cfgT == nil {
		return nil
	}
	config := cfgT.GetValue()
	return config.GetConfig(ctx)
}

func (d *DBService) RunMigrations(ctx context.Context, files map[string]string) (int, int, error) {
	if d.db == nil {
		return 0, 0, fmt.Errorf("❌ Banco de dados não inicializado")
	}
	sqlDB, err := d.db.DB()
	if err != nil {
		return 0, 0, fmt.Errorf("❌ Erro ao obter conexão SQL: %v", err)
	}
	conn, err := sqlDB.Conn(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("❌ Erro ao obter conexão SQL: %v", err)
	}
	defer conn.Close()

	return 0, 0, nil

}

func connectDatabase(_ context.Context, dbType, dsn string) (*gorm.DB, error) {
	// var dialector *sql.DB
	var dialector *sql.DB
	var err error
	// Abre a conexão SQL padrão
	switch dbType {
	case "mysql":
		dialector, err = sql.Open("mysql", dsn)
	case "postgres", "postgresql":
		dialector, err = sql.Open("postgres", dsn)
	case "sqlite":
		dialector, err = sql.Open("sqlite", dsn)
	case "mariadb":
		dialector, err = sql.Open("mariadb", dsn)
	case "sqlserver":
		dialector, err = sql.Open("sqlserver", dsn) // Implementar quando necessário
	case "oracle":
		// dialector = oracle.Open(dsn) // Implementar quando necessário
		return nil, fmt.Errorf("banco de dados Oracle não suportado no momento")
	case "mongodb":
		return nil, fmt.Errorf("banco de dados MongoDB não suportado pelo GORM")
	case "redis":
		return nil, fmt.Errorf("banco de dados Redis não suportado pelo GORM")
	case "rabbitmq":
		return nil, fmt.Errorf("RabbitMQ não é um banco de dados suportado pelo GORM")
	default:
		return nil, fmt.Errorf("banco de dados não suportado: %s", dbType)
	}
	if err != nil {
		return nil, fmt.Errorf("❌ Erro ao abrir conexão SQL: %v", err)
	}

	var gormDialector gorm.Dialector
	switch dbType {
	case "mysql":
		gormDialector = mysql.New(mysql.Config{
			Conn:                      dialector,
			SkipInitializeWithVersion: false,
		})
	case "postgres", "postgresql":
		gormDialector = postgres.New(postgres.Config{
			Conn:                 dialector,
			PreferSimpleProtocol: true, // Recomendado para evitar problemas com tipos complexos
		})
	case "sqlite":
		gormDialector = sqlite.New(sqlite.Config{
			Conn: dialector,
		})
	case "mariadb":
		gormDialector = mysql.New(mysql.Config{
			Conn:                      dialector,
			SkipInitializeWithVersion: false,
		})
	case "sqlserver":
		gormDialector = sqlserver.New(sqlserver.Config{
			Conn: dialector,
		})
	default:
		return nil, fmt.Errorf("banco de dados não suportado: %s", dbType)
	}

	db, err := gorm.Open(gormDialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ Erro ao conectar ao banco de dados: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("❌ Erro ao obter conexão SQL: %v", err)
	}

	// Configurações opcionais de pool de conexões
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(0)

	// Testa a conexão
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("❌ Erro ao pingar o banco de dados: %v", err)
	}

	return db, nil
}

// waitAndConnect aguarda PostgreSQL estar pronto e retorna conexão
func waitAndConnect(ctx context.Context, cfg *ti.Database, maxWait time.Duration) (*gorm.DB, *sql.Conn, error) {
	maxAttempts := 5
	retryInterval := 5 * time.Second
	if maxWait > 0 {
		maxAttempts = int(maxWait / retryInterval)
		if maxAttempts < 1 {
			maxAttempts = 1
		}
	}

	gl.Log("debug", fmt.Sprintf("⏳ Aguardando PostgreSQL ficar pronto (até %d tentativas em %v)...", maxAttempts, maxAttempts*int(retryInterval.Seconds())))

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err := connectDatabase(ctx, "postgresql", cfg.ConnectionString)
		if err != nil {
			gl.Log("debug", fmt.Sprintf("Tentativa %d/%d: falha ao conectar: %v", attempt, maxAttempts, err))
			time.Sleep(retryInterval)
			continue
		}

		sqlDB, err := db.DB()
		if err != nil {
			gl.Log("debug", fmt.Sprintf("Tentativa %d/%d: falha ao obter DB: %v", attempt, maxAttempts, err))
			time.Sleep(retryInterval)
			continue
		}

		conn, err := sqlDB.Conn(ctx)
		if err != nil {
			gl.Log("debug", fmt.Sprintf("Tentativa %d/%d: falha ao obter conexão: %v", attempt, maxAttempts, err))
			time.Sleep(retryInterval)
			continue
		}
		defer conn.Close()

		pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		if err := conn.PingContext(pingCtx); err != nil {
			cancel()
			gl.Log("debug", fmt.Sprintf("Tentativa %d/%d: falha ao pingar conexão: %v", attempt, maxAttempts, err))
			time.Sleep(retryInterval)
			continue
		}
		cancel()

		gl.Log("info", fmt.Sprintf("✅ PostgreSQL pronto (tentativa %d/%d)", attempt, maxAttempts))
		return db, conn, nil
	}

	return nil, nil, fmt.Errorf("PostgreSQL não ficou pronto após %d tentativas em %v", maxAttempts, maxAttempts*int(retryInterval.Seconds()))
}

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
