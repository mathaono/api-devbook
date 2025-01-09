package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Representa a estrutura de um usuário
type User struct {
	ID        int       `json:"ID,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Chama os métodos privados de validação e formatação de campos
func (user *User) Prepare(step string) error {
	err := user.validate(step)
	if err != nil {
		return err
	}

	err = user.format(step)
	if err != nil {
		return err
	}
	return nil
}

// Validação dos campos
func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name is required")
	} else if user.Nickname == "" {
		return errors.New("nickname is required")
	} else if user.Email == "" {
		return errors.New("email is required")
	} else if step == "create" && user.Password == "" {
		return errors.New("password is required")
	} else {
		err := checkmail.ValidateFormat(user.Email)
		if err != nil {
			return errors.New("email with invalid format")
		}
		return nil
	}
}

// Formatação dos campos string
func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)

	if step == "create" {
		passwordHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
	}

	return nil
}
