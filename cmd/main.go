package main

import (
	"flag"
	"fmt"
	"os"

	"importadorv2/internal/db"
	"importadorv2/internal/importer"
)

func main() {
	filePath := flag.String("file", "", "Path to the Excel file")
	host := flag.String("host", "localhost", "Database host")
	user := flag.String("user", "user", "Database user")
	password := flag.String("password", "password", "Database password")
	dbName := flag.String("dbname", "contacts_db", "Database name")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Erro: planilha não encontrada")
		os.Exit(1)
	}

	// Conecta ao banco de dados
	conn, err := db.Connect(*host, *user, *password, *dbName)
	if err != nil {
		fmt.Printf("Falha ao conectar ao banco: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Importa contatos
	if err := importer.ImportContacts(*filePath, conn); err != nil {
		fmt.Printf("Falha ao importar os contatos: %v\n", err)
		os.Exit(1)
	}
}
