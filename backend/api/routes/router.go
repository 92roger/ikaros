package routes

import (
	"ikaros/backend/api/middleware"
	"ikaros/backend/api/services"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	AccessToken = "ABCTOKEN123"
	AccessSigla = "admin"
)

// SetupRouter configura e retorna um roteador do Gorilla Mux
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Adiciona o middleware de autenticação
	r.Use(AuthMiddleware)

	// Adiciona o middleware de logging
	r.Use(LoggerMiddleware)
	// Endpoint de teste para gerar um token (não use em produção!)
	r.HandleFunc("/gerar-token", services.GerarToken).Methods("GET")
	// Aplica o middleware de limitação de taxa
	r.Use(middleware.RateLimitMiddleware)
	r.HandleFunc("/processar-csv", services.ProcessarCSVHandler).Methods("GET")

	// Atualize a rota para usar o novo handler
	r.HandleFunc("/dados-cadastrais", services.ProcessarDadosCadastraisHandler).Methods("GET")
	r.HandleFunc("/dados-cadastrais", services.InserirDadosCadastraisHandler).Methods("POST")
	r.HandleFunc("/dados-cadastrais", services.AtualizarDadosCadastraisHandler).Methods("PUT")
	r.HandleFunc("/dados-cadastrais", services.DeletarDadosCadastraisHandler).Methods("DELETE")

	return r
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.Header.Set("Authorization", AccessToken)
		r.Header.Set("X-Sigla", AccessSigla)

		token := r.Header.Get("Authorization")
		sigla := r.Header.Get("X-Sigla")

		accessToken := os.Getenv("ACCESS_TOKEN")
		accessSigla := os.Getenv("ACCESS_SIGLA")

		if token != accessToken || sigla != accessSigla {
			http.Error(w, "Acesso negado", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// LoggerMiddleware é um middleware que registra informações sobre cada requisição
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r) // Processa a requisição
		log.Printf("%s %s %s %v", r.RemoteAddr, r.Method, r.URL.Path, time.Since(start))
	})
}
