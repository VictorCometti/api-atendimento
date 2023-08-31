package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func getStringConnection() (connectionString string, erro error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("src/config")

	if erro = viper.ReadInConfig(); erro != nil {
		log.Fatalf("Erro ao tentar ler o árquivo de conexão: Erro: %v", erro)
	}

	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	hostname := viper.GetString("database.hostname")
	port := viper.GetString("database.port")
	database := viper.GetString("database.database_name")

	connectionString = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"
	return

}
func GetConnection() (db *sql.DB, erro error) {
	connectionString, erro := getStringConnection()
	if erro != nil {
		log.Fatalf("Erro ao tentar pegar a string de conexão. Erro: %v", erro)
		return
	}
	db, erro = sql.Open("postgres", connectionString)
	if erro != nil {
		log.Fatalf("Erro ao tentar abrir a conexão com o banco de dados. Erro: %v", erro)
		return
	}
	if erro = db.Ping(); erro != nil {
		log.Fatalf("Erro na hora de consultar o ping da conexão. Erro: %v", erro)
		return
	}
	return

}
