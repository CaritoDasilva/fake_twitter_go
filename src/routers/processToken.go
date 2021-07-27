package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fake_twitter/src/bd"
	"github.com/fake_twitter/src/models"
)

var Email string
var IDUser string

//ProcessToken trae los valores del token
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myPass := []byte("EstaEsMiClaveSecreta")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return myPass, nil
	})
	if err == nil {
		_, found, _ := bd.CheckUserExist(claims.Email)
		if found {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}

	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
