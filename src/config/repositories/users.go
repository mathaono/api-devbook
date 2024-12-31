package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Representa um repositório de usuários
type users struct {
	db *sql.DB
}

// Criação de um repositório de usuários
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Insere um usuário no banco de dados
func (repoUser users) Create(user models.User) (int64, error) {
	statement, err := repoUser.db.Prepare(
		"INSERT INTO users (name, nickname, email, password) VALUES ($1, $2, $3, $4);")
	if err != nil {
		fmt.Printf("Erro ao criar novo usuário na tabela: %s", err)
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return id, nil
}
