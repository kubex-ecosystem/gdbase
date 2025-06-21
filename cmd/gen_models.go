package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	is "github.com/rafa-mori/gdbase/internal/services"
	gl "github.com/rafa-mori/gdbase/logger"
	t "github.com/rafa-mori/gdbase/types"
	l "github.com/rafa-mori/logz"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Column struct {
	Name string
	Type string
}

// GenUser gera os modelos de usuário a partir do banco de dados
func Main() {
	// Inicializa o banco de dados
	_, dbSql, err := initDB()
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Erro ao inicializar o banco de dados: %v", err))
		return
	}
	defer dbSql.Close()

	// Consulta para obter estrutura das tabelas
	rows, err := dbSql.Query(`
        SELECT table_name, column_name, data_type
        FROM information_schema.columns
        WHERE table_schema = 'public';
    `)

	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Erro ao executar consulta: %v", err))
		return
	}

	defer rows.Close()

	tables := make(map[string][]Column)

	for rows.Next() {
		var tableName, columnName, dataType string
		if err := rows.Scan(&tableName, &columnName, &dataType); err != nil {
			gl.Log("fatal", fmt.Sprintf("Erro ao escanear linha: %v", err))
			return
		}
		tables[tableName] = append(tables[tableName], Column{Name: titleCase(columnName), Type: mapSQLType(dataType)})
	}

	// Gerar código Go a partir da estrutura
	generateGoModels(tables)
}

// Função para colocar título no nome dos campos
func titleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}

func initDB() (*gorm.DB, *sql.DB, error) {
	dbConfig := t.NewDBConfigWithFilePath("GoBE-DB", "/home/user/.kubex/gdbase/config/config.json")
	if dbConfig == nil {
		gl.Log("fatal", "Erro ao carregar configuração do banco de dados")
		return nil, nil, fmt.Errorf("erro ao carregar configuração do banco de dados")
	}
	// Inicializa o banco de dados
	// Criação do serviço de banco de dados
	dbService, err := is.NewDatabaseService(dbConfig, l.GetLogger("gen_models"))
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Erro ao criar serviço de banco de dados: %v", err))
		return nil, nil, err
	}
	// Inicialização do serviço de banco de dados
	err = dbService.Initialize()
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Erro ao inicializar serviço de banco de dados: %v", err))
		return nil, nil, err
	}
	// Configuração do banco de dados
	db, err := dbService.GetDB()
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Erro ao obter banco de dados: %v", err))
		return nil, nil, err
	}
	// Conexão com o banco de dados
	dbSql, err := db.DB()
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Erro ao obter conexão com o banco de dados: %v", err))
		return nil, nil, err
	}
	//defer dbSql.Close()

	if err := dbSql.Ping(); err != nil {
		gl.Log("fatal", fmt.Sprintf("Erro ao conectar ao banco de dados: %v", err))
		return nil, nil, err
	}
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

	return db, dbSql, nil
}

// Gera structs Go dinamicamente
func generateGoModels(tables map[string][]Column) {
	modelTemplate := `package main

{{range $table, $columns := .}}
type {{$table | title}} struct {
	{{range $columns}}
		{{.Name}} {{.Type}} ` + "`" + `json:"{{.Name}}" yaml:"{{.Name}}" xml:"{{.Name}}"` + "`" + `{{end}}
}
{{end}}
`

	file, err := os.Create("models.go")
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Erro ao criar arquivo: %v", err))
		return
	}
	defer file.Close()

	tmpl, err := template.New("models").
		Funcs(template.FuncMap{"title": titleCase}).
		Parse(modelTemplate)
	// Option("missingkey=zero") para evitar erro de chave ausente
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Erro ao criar template: %v", err))
		return
	}

	writer := io.Writer(file)
	if err = tmpl.Execute(writer, tables); err != nil {
		gl.Log("fatal", fmt.Sprintf("Erro ao executar template: %v", err))
		return
	}
	file.Sync()
	fmt.Println("Arquivo models.go gerado com sucesso!")
}

// Mapeia tipos SQL para Go
func mapSQLType(sqlType string) string {
	switch strings.ToLower(sqlType) {
	case "integer":
		return "int"
	case "numeric":
		return "float64"
	case "text", "varchar", "character varying":
		return "string"
	case "timestamp", "timestamp without time zone":
		return "time.Time"
	case "boolean":
		return "bool"
	default:
		return "any"
	}
}
