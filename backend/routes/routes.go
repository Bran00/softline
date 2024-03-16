package routes

import (
	"backend/controllers"
	"backend/middleware"
	"database/sql"

	"github.com/gorilla/mux"
)

// SetupRoutes configura as rotas da aplicação
func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Rota para login
	r.HandleFunc("/login", controllers.LoginController).Methods("POST")

	// Middleware JWT para rotas de clientes
	clientesRoute := r.PathPrefix("/clientes").Subrouter()
	clientesRoute.Use(middleware.JWTMiddleware)
	clientesRoute.HandleFunc("", controllers.GetClientes(db)).Methods("GET")
	clientesRoute.HandleFunc("", controllers.CreateCliente(db)).Methods("POST")
	clientesRoute.HandleFunc("/{id}", controllers.GetClienteByID(db)).Methods("GET")
	clientesRoute.HandleFunc("/{id}", controllers.UpdateCliente(db)).Methods("PUT")
	clientesRoute.HandleFunc("/{id}", controllers.DeleteCliente(db)).Methods("DELETE")

	// Middleware JWT para rotas de produtos
	produtosRoute := r.PathPrefix("/produtos").Subrouter()
	produtosRoute.Use(middleware.JWTMiddleware)
	produtosRoute.HandleFunc("", controllers.GetProdutos(db)).Methods("GET")
	produtosRoute.HandleFunc("", controllers.CreateProduto(db)).Methods("POST")
	produtosRoute.HandleFunc("/{id}", controllers.GetProdutoByID(db)).Methods("GET")
	produtosRoute.HandleFunc("/{id}", controllers.UpdateProduto(db)).Methods("PUT")
	produtosRoute.HandleFunc("/{id}", controllers.DeleteProduto(db)).Methods("DELETE")

	return r
}
