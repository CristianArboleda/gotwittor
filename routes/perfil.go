package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CristianArboleda/gotwittor/db"
)

func GetPerfil(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(rw, "The id param is required", http.StatusBadRequest)
		return
	}
	perfil, err := db.FindUserById(ID)
	if err != nil {
		http.Error(rw, "Error trying to get the record: "+err.Error(), http.StatusBadRequest)
		return
	}
	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(perfil)
}
