package routes

import (
	"backend/controllers"
	"database/sql"

	"github.com/gorilla/mux"
)

// SetupRoutes configura as rotas da aplicação
func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Rotas para clientes
	r.HandleFunc("/clientes", controllers.GetClientes(db)).Methods("GET")
	r.HandleFunc("/cliente", controllers.CreateCliente(db)).Methods("POST")
	r.HandleFunc("/cliente/{id}", controllers.GetClienteByID(db)).Methods("GET")
	r.HandleFunc("/cliente/{id}", controllers.UpdateCliente(db)).Methods("PUT")
	r.HandleFunc("/cliente/{id}", controllers.DeleteCliente(db)).Methods("DELETE")

	// Rotas para produtos
	r.HandleFunc("/produtos", controllers.GetProdutos(db)).Methods("GET")
	r.HandleFunc("/produto", controllers.CreateProduto(db)).Methods("POST")
	r.HandleFunc("/produto/{id}", controllers.GetProdutoByID(db)).Methods("GET")
	r.HandleFunc("/produto/{id}", controllers.UpdateProduto(db)).Methods("PUT")
	r.HandleFunc("/produto/{id}", controllers.DeleteProduto(db)).Methods("DELETE")

	return r
}
