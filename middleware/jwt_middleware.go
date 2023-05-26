package middleware

import (
	"code/competence/constant"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id string, email string, password string) (string, error) {

	// payload
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["email"] = email
	claims["password"] = password

	// header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// return + signature
	return token.SignedString([]byte(constant.SECRET_JWT))

}

func GetIDFromToken(tokenString string) (string, error) {
	// Parsing dan verifikasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.SECRET_JWT), nil
	})
	if err != nil {
		return "", err
	}

	// Memeriksa apakah token valid dan memiliki claims yang sesuai
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(string)
		return id, nil
	}

	return "", fmt.Errorf("Invalid token")
}
