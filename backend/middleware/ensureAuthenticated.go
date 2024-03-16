package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

// JWTMiddleware é um middleware para verificar a validade do token JWT
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtenha o token JWT do cabeçalho Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "token JWT ausente", http.StatusUnauthorized)
			return
		}

		// Verifique se o cabeçalho Authorization está no formato "Bearer token"
		authComponents := strings.Split(authHeader, " ")
		if len(authComponents) != 2 || authComponents[0] != "Bearer" {
			http.Error(w, "formato de token JWT inválido", http.StatusUnauthorized)
			return
		}

		// Extrair o token JWT da parte "token" do cabeçalho
		tokenString := authComponents[1]

		// Verificar a validade do token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verificar o método de assinatura do token
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inválido: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil 
		})
		if err != nil || !token.Valid {
			http.Error(w, "token JWT inválido", http.StatusUnauthorized)
			return
		}

		// Se o token for válido, chame o próximo handler
		next.ServeHTTP(w, r)
	})
}
