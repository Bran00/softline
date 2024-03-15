package controllers

import (
	"backend/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetProdutos retorna todos os produtos do banco de dados
func GetProdutos(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Consulta ao banco de dados
		rows, err := db.Query("SELECT * FROM produtos")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Criando slice para armazenar os produtos
		var produtos []models.Produto

		// Iterando sobre os resultados da consulta
		for rows.Next() {
			var p models.Produto
			err := rows.Scan(&p.ID, &p.Nome, &p.Descricao, &p.CodigoBarras, &p.ValorVenda, &p.PesoBruto, &p.PesoLiquido)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			produtos = append(produtos, p)
		}

		// Verificando se ocorreu algum erro durante a iteração
		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convertendo para JSON e escrevendo na resposta HTTP
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(produtos)
	}
}

// GetProdutoByID retorna um único produto pelo seu ID
func GetProdutoByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extrair o ID do produto dos parâmetros da URL
		params := mux.Vars(r)
		produtoID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Consultar o banco de dados para obter o produto pelo ID
		var p models.Produto
		err = db.QueryRow("SELECT * FROM produtos WHERE id = ?", produtoID).Scan(&p.ID, &p.Nome, &p.Descricao, &p.CodigoBarras, &p.ValorVenda, &p.PesoBruto, &p.PesoLiquido)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Escrever o produto encontrado na resposta HTTP
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	}
}

// CreateProduto cria um novo produto no banco de dados
func CreateProduto(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var p models.Produto
        err := json.NewDecoder(r.Body).Decode(&p)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Verificar se todos os campos obrigatórios estão presentes
        if p.Nome == "" || p.Descricao == "" || p.CodigoBarras == "" || p.ValorVenda == 0 || p.PesoBruto == 0 || p.PesoLiquido == 0 {
            http.Error(w, "Todos os campos devem ser preenchidos", http.StatusBadRequest)
            return
        }

        // Inserir o produto no banco de dados
        _, err = db.Exec("INSERT INTO produtos (nome, descricao, codigo_barras, valor_venda, peso_bruto, peso_liquido) VALUES (?, ?, ?, ?, ?, ?)",
            p.Nome, p.Descricao, p.CodigoBarras, p.ValorVenda, p.PesoBruto, p.PesoLiquido)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]string{"message": "Produto criado com sucesso"})
    }
}

func UpdateProduto(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extrair o ID do produto dos parâmetros da URL
        params := mux.Vars(r)
        produtoID, err := strconv.Atoi(params["id"])
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        var p models.Produto
        err = json.NewDecoder(r.Body).Decode(&p)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Verificar se todos os campos obrigatórios estão presentes
        if p.Nome == "" || p.Descricao == "" || p.CodigoBarras == "" || p.ValorVenda == 0 || p.PesoBruto == 0 || p.PesoLiquido == 0 {
            http.Error(w, "Todos os campos devem ser preenchidos", http.StatusBadRequest)
            return
        }

        // Atualizar o produto no banco de dados
        _, err = db.Exec("UPDATE produtos SET nome = ?, descricao = ?, codigo_barras = ?, valor_venda = ?, peso_bruto = ?, peso_liquido = ? WHERE id = ?",
            p.Nome, p.Descricao, p.CodigoBarras, p.ValorVenda, p.PesoBruto, p.PesoLiquido, produtoID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Produto atualizado com sucesso"})
    }
}

// DeleteProduto deleta um produto existente no banco de dados
func DeleteProduto(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extrair o ID do produto dos parâmetros da URL
        params := mux.Vars(r)
        produtoID, err := strconv.Atoi(params["id"])
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Verificar se o ID do produto é válido
        if produtoID <= 0 {
            http.Error(w, "ID do produto inválido", http.StatusBadRequest)
            return
        }

        // Deletar o produto do banco de dados
        _, err = db.Exec("DELETE FROM produtos WHERE id = ?", produtoID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Responder com sucesso
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Produto deletado com sucesso"})
    }
}
