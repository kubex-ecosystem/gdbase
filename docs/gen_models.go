package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	gl "github.com/kubex-ecosystem/gdbase/internal/module/logger"
	is "github.com/kubex-ecosystem/gdbase/internal/services"
	t "github.com/kubex-ecosystem/gdbase/types"
	l "github.com/kubex-ecosystem/logz"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Column struct {
	Name string
	Type string
}

// Main GenerateModels generates user models from database
func Main() {
	// Initialize database
	_, dbSQL, err := initDB()
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Error initializing database: %v", err))
		return
	}
	defer dbSQL.Close()

	// Query to get table structure
	rows, err := dbSQL.Query(`
        SELECT table_name, column_name, data_type
        FROM information_schema.columns
        WHERE table_schema = 'public';
    `)

	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Error executing query: %v", err))
		return
	}

	defer rows.Close()

	tables := make(map[string][]Column)

	for rows.Next() {
		var tableName, columnName, dataType string
		if err := rows.Scan(&tableName, &columnName, &dataType); err != nil {
			gl.Log("fatal", fmt.Sprintf("Error scanning row: %v", err))
			return
		}
		tables[tableName] = append(tables[tableName], Column{Name: titleCase(columnName), Type: mapSQLType(dataType)})
	}

	// Generate Go code from structure
	generateGoModels(tables)
}

// Function to title case field names
func titleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}

func initDB() (*gorm.DB, *sql.DB, error) {
	dbConfig := t.NewDBConfigWithFilePath("GoBE-DB", "/home/user/.kubex/gdbase/config/config.json")
	if dbConfig == nil {
		gl.Log("fatal", "Error loading database configuration")
		return nil, nil, fmt.Errorf("error loading database configuration")
	}
	// Initialize database
	// Create database service
	dbService, err := is.NewDatabaseService(dbConfig, l.GetLogger("gen_models"))
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Error creating database service: %v", err))
		return nil, nil, err
	}
	// Initialize database service
	err = dbService.Initialize()
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Error initializing database service: %v", err))
		return nil, nil, err
	}
	// Get database configuration
	db, err := dbService.GetDB()
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Error getting database: %v", err))
		return nil, nil, err
	}
	// Database connection
	dbSQL, err := db.DB()
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Error getting database connection: %v", err))
		return nil, nil, err
	}
	//defer dbSql.Close()

	if err := dbSQL.Ping(); err != nil {
		gl.Log("fatal", fmt.Sprintf("Error connecting to database: %v", err))
		return nil, nil, err
	}
	fmt.Println("Database connection established successfully!")

	return db, dbSQL, nil
}

// Generate Go structs dynamically
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
		gl.Log("fatal", fmt.Sprintf("Error creating file: %v", err))
		return
	}
	defer file.Close()

	tmpl, err := template.New("models").
		Funcs(template.FuncMap{"title": titleCase}).
		Parse(modelTemplate)
	// Option("missingkey=zero") to avoid missing key error
	if err != nil {
		gl.Log("fatal", fmt.Sprintf("Error creating template: %v", err))
		return
	}

	writer := io.Writer(file)
	if err = tmpl.Execute(writer, tables); err != nil {
		gl.Log("fatal", fmt.Sprintf("Error executing template: %v", err))
		return
	}
	file.Sync()
	fmt.Println("models.go file generated successfully!")
}

// Maps SQL types to Go types
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
