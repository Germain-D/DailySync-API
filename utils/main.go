package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Ajouter des claims (informations supplémentaires)
	claims["authorized"] = true
	claims["user"] = "testUser"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Expiration dans 24h

	// Signer le token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {
	token, err := GenerateJWT()
	if err != nil {
		fmt.Println("Erreur lors de la génération du token:", err)
		return
	}
	fmt.Println("Token JWT généré:", token)

}
