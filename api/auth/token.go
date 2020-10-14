package auth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CreateToken ...
func CreateToken(userid uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorised"] = true
	claims["user_id"] = userid
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
