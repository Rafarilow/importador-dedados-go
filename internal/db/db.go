package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" 
)

// Connect conecta ao banco e garante que a tabela exista
func Connect(host, user, password, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, dbName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao banco: %w", err)
	}

	if err := ensureTableExists(db); err != nil {
		return nil, fmt.Errorf("falha ao criar a tabela(ela já existe): %w", err)
	}

	return db, nil
}

// cria a tabela "contacts" se ela não existir.
func ensureTableExists(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS contacts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		nome VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		telefone VARCHAR(20) NOT NULL
	);`
	_, err := db.Exec(query)
	return err
}
