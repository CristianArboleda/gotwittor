package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CristianArboleda/gotwittor/db"
	"github.com/CristianArboleda/gotwittor/models"
)

// Register : function to register to new user
func Register(rw http.ResponseWriter, r *http.Request) {
	var us models.User
	err := json.NewDecoder(r.Body).Decode(&us)
	if err != nil {
		http.Error(rw, "Error in the body request: "+err.Error(), 400)
		return
	}

	if len(us.Email) == 0 {
		http.Error(rw, "The email field is required", 400)
		return
	}
	if len(us.Password) < 6 {
		http.Error(rw, "The password must be a minimun of 6 character", 400)
		return
	}

	_, exist, _ := db.FindUserByEmail(us.Email)
	if exist {
		http.Error(rw, "Alrready exist a registered user with this email.", 400)
		return
	}

	_, status, errorSaving := db.SaveUser(us)
	if errorSaving != nil {
		http.Error(rw, "Error saving the user: "+errorSaving.Error(), 400)
		return
	}
	if status == false {
		http.Error(rw, "Error saving the user: invalid status", 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
