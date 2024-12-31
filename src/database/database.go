package database

import (
	"api/src/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Certifique-se de importar o driver PostgreSQL
)

// Abre e retorna conexão com a base de dados
func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.ConnectStringDB)
	if err != nil {
		log.Fatal("Erro de conexão com o banco de dados: ", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
