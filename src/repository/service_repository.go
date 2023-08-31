package repository

import (
	"database/sql"
	"log"

	"github.com/VictorCometti/api-atendimento/src/models"
)

type Repository struct {
	DB *sql.DB
}

func (repository *Repository) FindServiceByIdentifier(serviceIdentifier int) (services []models.Service, erro error) {
	query := ("SELECT * FROM api_schema.services WHERE service_identifier = ?")

	rows, err := repository.DB.Query(query, serviceIdentifier)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var service models.Service
		if erro = rows.Scan(
			&service.ServiceIdentifier,
			&service.ServiceDate,
			&service.Exams,
		); erro != nil {
			log.Fatalf("Erro ao tentar scanear as linhas da query. Erro: %v", erro)
		}
		services = append(services, service)
	}

	return
}

func NewRepository() (repository *Repository) {
	return
}
