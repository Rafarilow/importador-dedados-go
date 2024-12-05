package importer

import (
	"database/sql"
	"fmt"

	"importadorv2/internal/excel"
)

// importa contatos do Excel para o banco.
func ImportContacts(filePath string, db *sql.DB) error {
	contacts, err := excel.Read(filePath)
	if err != nil {
		return err
	}

	for _, contact := range contacts {
		_, err := db.Exec(
			`INSERT INTO contacts (nome, email, telefone) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE nome=VALUES(nome), telefone=VALUES(telefone)`,
			contact.Nome, contact.Email, contact.Telefone,
		)
		if err != nil {
			return fmt.Errorf("falha ao inserir contato %v: %v", contact, err)
		}
	}

	fmt.Println("Contatos importados com sucesso, você é foda!")
	return nil
}
