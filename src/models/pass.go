package models

// Representa o formato da alteração de senha
type Pass struct {
	NewPass     string `json:"newPass"`
	CurrentPass string `json:"currentPass"`
}
