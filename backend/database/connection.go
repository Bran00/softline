package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// InitDB inicializa a conexão com o banco de dados
func InitDB() (*sql.DB, error) {
	// Conectando ao banco de dados
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	// Verifique se a conexão com o banco de dados é bem-sucedida
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CloseDB fecha a conexão com o banco de dados
func CloseDB(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}
