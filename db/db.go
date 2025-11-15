package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB inicializa a conexão com o banco de dados PostgreSQL.
func InitDB() {
	// Variáveis de ambiente para conexão com o PostgreSQL
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if password == "" {
		password = "mysecretpassword"
	}
	if dbname == "" {
		dbname = "upskilling_db"
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}

	// Testa a conexão
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados PostgreSQL estabelecida com sucesso.")
}

// GetDB retorna a instância da conexão com o banco de dados.
func GetDB() *sql.DB {
	return db
}

// CloseDB fecha a conexão com o banco de dados.
func CloseDB() {
	if db != nil {
		db.Close()
		log.Println("Conexão com o banco de dados fechada.")
	}
}
