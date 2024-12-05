package excel

import (
	"fmt"
	"regexp"

	"github.com/xuri/excelize/v2"
)

type Contact struct {
	Nome     string
	Email    string
	Telefone string
}

// lê um arquivo Excel e retorna uma lista de contatos.
func Read(filePath string) ([]Contact, error) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	contacts := []Contact{}
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	for i, row := range rows {
		if i == 0 {
			continue // Ignora o header
		}

		if len(row) < 3 {
			return nil, fmt.Errorf("row %d campo incompleto", i+1)
		}

		contact := Contact{
			Nome:     row[0],
			Email:    row[1],
			Telefone: row[2],
		}

		// Validações
		if !isValidEmail(contact.Email) {
			return nil, fmt.Errorf("e-mail invalido: %s", contact.Email)
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}

// valida se o e-mail é válido.
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
