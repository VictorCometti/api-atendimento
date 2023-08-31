package repository

import (
	"database/sql"
	"log"

	"github.com/VictorCometti/api-atendimento/src/models"
)

type repository struct {
	DB *sql.DB
}

func NewServiceRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (repository *repository) FindServiceByIdentifier(serviceIdentifier int) (exams []models.ExamDTO, erro error) {
	query := "SELECT service.service_date, service.service_identifier, exam.patient_name, exam.exam_description, exam.requesting_physician FROM atendimento_schema.service LEFT JOIN atendimento_schema.exam ON service.id = exam.service_id WHERE service.service_identifier = $1"

	rows, erro := repository.DB.Query(query, serviceIdentifier)

	if erro != nil {
		log.Fatalf("Erro ao tentar executar a query de buscar de serviço.")
	}

	defer rows.Close()

	for rows.Next() {
		var exam models.ExamDTO

		if erro = rows.Scan(
			&exam.ServiceDate,
			&exam.ServiceIdentifier,
			&exam.PatientName,
			&exam.ExamDescription,
			&exam.RequestingPhysician,
		); erro != nil {
			log.Fatalf("Erro ao tentar scanear as linhas do banco de dados. Erro: %v", erro)
		}
		exams = append(exams, exam)
	}

	return
}
