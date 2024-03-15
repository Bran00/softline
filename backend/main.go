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

	// Iniciando o servidor
	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
