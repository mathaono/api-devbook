package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// String de conexão com o postgreSQL
	// Porta que a API estará rodando
	ConnectStringDB = ""
	Port            = 0
)

// Carrega as variáveis de ambiente
func Loading() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Erro na variável de ambiente: ", err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	ConnectStringDB = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
}
