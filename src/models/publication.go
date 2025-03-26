package models

import (
	"errors"
	"strings"
	"time"
)

// Representa uma publicação feita por um usuário; É necessário que uma publicação tenha um usuário
type Publication struct {
	ID        int64     `json:"ID,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	UserID    uint64    `json:"userID,omitempty"`
	UserNick  string    `json:"userNick,omitempty"`
	Likes     uint64    `json:"likes"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Chama os métodos de validação e formatação para a publicação recebida
func (pub *Publication) Prepare() error {
	if err := pub.validate(); err != nil {
		return err
	}

	pub.format()
	return nil
}

func (pub *Publication) validate() error {
	if pub.Title == "" {
		return errors.New("Title is required")
	}

	if pub.Content == "" {
		return errors.New("Content is required")
	}

	return nil
}

func (pub *Publication) format() {
	pub.Title = strings.TrimSpace(pub.Title)
	pub.Content = strings.TrimSpace(pub.Content)
}
