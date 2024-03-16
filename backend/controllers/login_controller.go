package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Chave secreta padrão para assinar os tokens JWT
var secretKey = []byte(os.Getenv("SECRET_KEY"))

// JWTClaims define a estrutura das informações incluídas no token JWT
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// LoginController realiza a autenticação do administrador e retorna um token JWT
func LoginController(w http.ResponseWriter, r *http.Request) {
	// Obtenha as variáveis de ambiente para o usuário e a senha do administrador
	userAdmin := os.Getenv("ADMIN_USER")
	hashedPasswordAdmin := hashPassword(os.Getenv("PASSWORD"))

	// Verifique se as variáveis de ambiente foram definidas
	if userAdmin == "" || hashedPasswordAdmin == "" {
		http.Error(w, "Credenciais do administrador não configuradas", http.StatusInternalServerError)
		return
	}

	// Decodifique o corpo da solicitação JSON para obter os valores de username e password
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação", http.StatusBadRequest)
		return
	}

	// Compare o nome de usuário fornecido com o usuário do administrador
	if credentials.Username != userAdmin {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Compare a senha fornecida com a senha hashada do administrador
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswordAdmin), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Crie um token JWT com as informações do usuário
	tokenClaims := JWTClaims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
			Issuer:    "seu-servidor",                        // Quem emitiu o token
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	// Assine o token com a chave secreta e obtenha o token como uma string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	// Se as credenciais estiverem corretas, retorne uma mensagem de sucesso com o token JWT
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login bem-sucedido",
		"token":   tokenString,
	})
}

// hashPassword gera o hash de uma senha usando bcrypt
func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err) // Trate o erro de acordo com sua aplicação
	}
	return string(hashedPassword)
}
