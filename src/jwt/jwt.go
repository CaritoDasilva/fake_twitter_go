package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fake_twitter/src/models"
)

func GenerateJWT(t models.User) (string, error) {
	miClave := []byte("EstaEsMiClaveSecreta")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"birthDate": t.BirthDate,
		"bio":       t.Bio,
		"location":  t.Location,
		"website":   t.Website,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
