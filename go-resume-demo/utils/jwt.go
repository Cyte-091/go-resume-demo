package utils

import "github.com/golang-jwt/jwt/v5"

var jwtSecret = []byte("go-resume-demo-secret")

type Claims struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	claims := Claims{UserName: username}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) { return jwtSecret, nil })
	if err != nil {
		return nil, err
	}
	return token.Claims.(*Claims), nil
}
