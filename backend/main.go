package main

import (
	"backend/database"
	"backend/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Inicializando o banco de dados
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
	}

	// Configurando rotas
	r := routes.SetupRoutes(db)

	// Configurando o middleware para permitir CORS
	allowedOrigins := map[string]bool{
		"http://localhost:3000": true, // Permitir origem do frontend
	}

	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if allowedOrigins[origin] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				if r.Method == "OPTIONS" {
					w.WriteHeader(http.StatusOK)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}

	// Adicionando o middleware CORS às rotas
	http.Handle("/", corsHandler(r))

	// Iniciando o servidor
	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
