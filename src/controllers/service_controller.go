package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/VictorCometti/api-atendimento/src/db"
	"github.com/VictorCometti/api-atendimento/src/repository"
)

func FindServiceByIdentifier(w http.ResponseWriter, r *http.Request) {
	serviceIdentifierString := (r.URL.Query().Get("identifier"))

	serviceIdentifierInt, erro := strconv.Atoi(serviceIdentifierString)

	if erro != nil {
		log.Fatalf("Erro ao tentar converter para int: Erro: %v", erro)
	}

	conn, erro := db.GetConnection()

	if erro != nil {
		log.Fatalf("Erro ao tentar abrir a conexão: Erro: %v", erro)
	}

	defer conn.Close()

	repositorio := repository.NewServiceRepository(conn)

	services, err := repositorio.FindServiceByIdentifier(serviceIdentifierInt)

	fmt.Printf("Service: %v", services)
	if err != nil {
		log.Fatalf("Erro ao tentar buscar um serviço pelo repositorio. Erro: %v", erro)
	}

	json, erro := json.Marshal(services)

	if erro != nil {
		log.Fatalf("Erro: %v", erro)
	}

	w.Write(json)
}

func UpdateServiceByUserIdentifier(w http.ResponseWriter, r *http.Request) {

}
