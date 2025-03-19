package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Retorna um response JSON para a request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("❌ [ERROR] - Failed to convert response: %v", err)
			log.Fatal(err)
		}
	}

	log.Printf("Status: %d - Info: %+v", statusCode, data)
}

// Retorna uma struct de erro no formato JSON
func Erro(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Erro       string `json:"erro"`
		HttpStatus int    `json:"httpStatus"`
	}{
		Erro:       err.Error(),
		HttpStatus: statusCode,
	})
}
