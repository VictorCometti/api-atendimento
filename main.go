package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VictorCometti/api-atendimento/src/router"
)

func main() {
	fmt.Println("Subindo a api")
	router := router.GetRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
