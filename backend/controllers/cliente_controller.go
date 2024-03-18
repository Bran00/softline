package controllers

import (
	"backend/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateCliente(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cliente models.Cliente
		if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Verificar se há campos vazios
		if cliente.Nome == "" || cliente.Fantasia == "" || cliente.Documento == "" || cliente.Endereco == "" {
			http.Error(w, "Todos os campos devem ser preenchidos", http.StatusBadRequest)
			return
		}

		// Insira a lógica para adicionar o cliente ao banco de dados
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer tx.Rollback() // Rollback se ocorrer um erro

		result, err := tx.Exec("INSERT INTO clientes (nome, fantasia, documento, endereco) VALUES (?, ?, ?, ?)",
			cliente.Nome, cliente.Fantasia, cliente.Documento, cliente.Endereco)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		clienteID, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cliente.ID = int(clienteID)

		if err := tx.Commit(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Cliente criado com sucesso",
			"cliente": cliente,
		})
	}
}

func GetClientes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM clientes")
		if err != nil {
			http.Error(w, "Erro ao buscar clientes", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var clientes []models.Cliente

		for rows.Next() {
			var c models.Cliente
			err := rows.Scan(&c.ID, &c.Nome, &c.Fantasia, &c.Documento, &c.Endereco)
			if err != nil {
				http.Error(w, "Erro ao ler dados do cliente", http.StatusInternalServerError)
				return
			}

			clientes = append(clientes, c)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clientes)
	}
}

func GetClienteByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		clienteID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "ID do cliente inválido", http.StatusBadRequest)
			return
		}

		var cliente models.Cliente
		err = db.QueryRow("SELECT id, nome, fantasia, documento, endereco FROM clientes WHERE id = ?", clienteID).Scan(&cliente.ID, &cliente.Nome, &cliente.Fantasia, &cliente.Documento, &cliente.Endereco)
		if err != nil {
			http.Error(w, "Erro ao buscar cliente pelo ID", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cliente)
	}
}

func UpdateCliente(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		clienteID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var cliente models.Cliente
		if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Verificar se há campos vazios
		if cliente.Nome == "" || cliente.Fantasia == "" || cliente.Documento == "" || cliente.Endereco == "" {
			http.Error(w, "Todos os campos devem ser preenchidos", http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("UPDATE clientes SET nome=?, fantasia=?, documento=?, endereco=? WHERE id=?")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(cliente.Nome, cliente.Fantasia, cliente.Documento, cliente.Endereco, clienteID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cliente.ID = clienteID

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Cliente atualizado com sucesso",
			"cliente": cliente,
		})
	}
}

// DeleteCliente exclui um cliente existente do banco de dados
func DeleteCliente(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		clienteID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("DELETE FROM clientes WHERE id=?")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(clienteID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Responder com sucesso
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Produto deletado com sucesso"})
	}
}
