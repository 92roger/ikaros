// tokenservice.go

package services

import (
	"encoding/json"
	"net/http"
	"os"
)

// GerarToken retorna o token e a sigla como JSON
func GerarToken(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("ACCESS_TOKEN")
	sigla := os.Getenv("ACCESS_SIGLA")
	// Retorna o token e a sigla como JSON
	json.NewEncoder(w).Encode(map[string]string{"token": token, "sigla": sigla})
}
