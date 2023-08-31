package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func getStringConnection() (connectionString string, erro error) {
	viper.SetConfigFile("config")
	viper.AddConfigPath("src/config")
	viper.ReadInConfig()

	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	databaseName := viper.GetString("database.databaseName")

	connectionString = "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + databaseName + "sslmode=false"
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
