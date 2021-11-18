package routes

import (
	"errors"
	"strings"

	"github.com/CristianArboleda/gotwittor/db"
	"github.com/CristianArboleda/gotwittor/jwt"
	"github.com/CristianArboleda/gotwittor/models"
	jwtgo "github.com/dgrijalva/jwt-go"
)

/*Email : email of current user */
var Email string

/*IDUsuario : ID of current user */
var IDUsuario string

/*ProccessToken : funtion that process the token */
func ProccessToken(token string) (*models.Claim, bool, string, error) {

	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, "", errors.New("invalit format token")
	}

	token = strings.TrimSpace(splitToken[1])

	tk, err := jwtgo.ParseWithClaims(token, claims, func(t *jwtgo.Token) (interface{}, error) {
		return jwt.JWTPass, nil
	})

	if err != nil {
		return claims, false, "", err
	}
	if !tk.Valid {
		return claims, false, "", errors.New("invalid token")
	}

	_, exist, _ := db.FindUserByEmail(claims.Email)

	if exist {
		Email = claims.Email
		IDUsuario = claims.ID.Hex()
	}

	return claims, exist, IDUsuario, nil
}
