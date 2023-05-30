package internal

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = "31f6817b-f731-4d09-8ee2-337c2a5e2a19"

func GenerateToken() (string, error) {
	// Crear una nueva reclamación (claims) con la información deseada
	claims := jwt.MapClaims{
		"sub": "usuario1",
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expira en 24 horas
	}
	// Crear el token JWT firmado con la clave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) bool {
	// Parsear el token y validar la firma
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inválido: %v", token.Header["alg"])
		}

		// Devolver la clave secreta utilizada para firmar el token
		return []byte(secretKey), nil
	})

	if err != nil {
		return false
	}

	// Verificar si el token es válido
	if token.Valid {
		return true
	}

	return false
}
