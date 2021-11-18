package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/CristianArboleda/gotwittor/db"
	"github.com/CristianArboleda/gotwittor/jwt"
	"github.com/CristianArboleda/gotwittor/models"
)

/*Login: method that check the login params */
func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")

	var us models.User

	err := json.NewDecoder(r.Body).Decode(&us)

	if err != nil {
		http.Error(rw, "Error in the Login body request: "+err.Error(), 400)
		return
	}

	if len(us.Email) == 0 {
		http.Error(rw, "The email field is required", 400)
		return
	}

	userDB, exist := db.CheckLogin(us.Email, us.Password)

	if !exist {
		http.Error(rw, "Invalid user and password", 400)
		return
	}

	jwtKey, errorJWT := jwt.BuildJWT(userDB)

	if errorJWT != nil {
		http.Error(rw, "Error building the JWT:"+errorJWT.Error(), 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(resp)

	//Cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(rw, &http.Cookie{
		Name: "token", Value: jwtKey, Expires: expirationTime,
	})

}
