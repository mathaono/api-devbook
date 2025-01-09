package security

import "golang.org/x/crypto/bcrypt"

// Recebe uma string e codifica ela com um Hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compara a senha com o Hash e valida se possuem o mesmo valor
func PasswordValidate(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
