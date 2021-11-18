package jwt

import (
	"time"

	"github.com/CristianArboleda/gotwittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Pass []byte = []byte("This_is_the_super_pass")

// BuildJWT : function that build the JWT
func BuildJWT(us models.User) (string, error) {

	payload := jwt.MapClaims{
		"email":     us.Email,
		"name":      us.Name,
		"lastName":  us.LastName,
		"birthDate": us.BirthDate,
		"_id":       us.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	jwtKey, err := token.SignedString(Pass)
	if err != nil {
		return jwtKey, err
	}

	return jwtKey, nil
}
