package authentication

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Retorna um Token assinado com as permissões do usuário
func CreateToken(userID uint64) (string, error) {
	// Permissões dentro do Token
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	// chave secret
	return token.SignedString([]byte(config.SecretKey))
}
