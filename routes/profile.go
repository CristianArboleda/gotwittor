package routes

import (
	"encoding/json"
	"github.com/CristianArboleda/gotwittor/db"
	"github.com/CristianArboleda/gotwittor/models"
	"net/http"
	"strconv"
)

// GetProfile : get profile info
func GetProfile(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "The id param is required", http.StatusBadRequest)
		return
	}
	profile, err := db.FindUserById(ID)
	if err != nil {
		http.Error(rw, "Error trying to get the record: "+err.Error(), http.StatusBadRequest)
		return
	}
	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(profile)
}

// UpdateProfile : Update profile info
func UpdateProfile(rw http.ResponseWriter, r *http.Request) {
	var us models.User
	err := json.NewDecoder(r.Body).Decode(&us)
	if err != nil {
		http.Error(rw, "Error in the body request: "+err.Error(), http.StatusBadRequest)
		return
	}

	//The logged user only can modify his profile
	status, errUpdate := db.UpdateUser(us, UserID)
	if errUpdate != nil {
		http.Error(rw, "Error updating the record : "+errUpdate.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(rw, "Error updating the record", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

// GetProfilesByFilters : Get a list of profiles by filters
func GetProfilesByFilters(rw http.ResponseWriter, r *http.Request) {
	relationType := r.URL.Query().Get("type")
	if len(relationType) < 1 {
		http.Error(rw, "The type param is required", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(rw, "The page param is required", http.StatusBadRequest)
		return
	}
	pag, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(rw, "The page param must be a numeric value", http.StatusBadRequest)
		return
	}
	page := int64(pag)
	search := r.URL.Query().Get("search")

	result, status := db.FindUsersByFilters(UserID, page, search, relationType)
	if !status {
		http.Error(rw, "Error searching the records", http.StatusBadRequest)
		return
	}
	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(result)
}
