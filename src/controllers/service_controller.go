package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/VictorCometti/api-atendimento/src/db"
	"github.com/VictorCometti/api-atendimento/src/repository"
	"github.com/VictorCometti/api-atendimento/src/response"
)

func FindServiceByIdentifier(w http.ResponseWriter, r *http.Request) {
	serviceIdentifierString := r.URL.Query().Get("identifier")
	serviceIdentifierInt, erro := strconv.Atoi(serviceIdentifierString)
	if erro != nil {
		log.Fatalf("Erro: %v", erro)
	}
	conn, erro := db.GetConnection()
	if erro != nil {
		log.Fatalf("Erro: %v", erro)
	}
	defer conn.Close()
	repository := repository.NewRepository()
	services, err := repository.FindServiceByIdentifier(serviceIdentifierInt)
	if err != nil {
		log.Fatalf("Erro: %v", erro)
	}

	json, erro := json.Marshal(services)
	if erro != nil {
		log.Fatalf("Erro: %v", erro)
	}
	response.JSON(w, http.StatusOK, json)
}

func UpdateServiceByUserIdentifier(w http.ResponseWriter, r *http.Request) {

}
