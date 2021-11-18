package routes

import (
	"encoding/json"
	"github.com/CristianArboleda/gotwittor/db"
	"github.com/CristianArboleda/gotwittor/models"
	"net/http"
	"time"
)

// SaveRelation : Save a relation between two users
func SaveRelation(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "The id param is required", http.StatusBadRequest)
		return
	}

	record, _ := db.FindRelationByUserAndRelationUser(UserID, ID)
	if len(record.UserID) > 0 {
		http.Error(rw, "The relation already exists", http.StatusBadRequest)
		return
	}
	var relation models.Relation
	relation.UserID = UserID
	relation.RelationUserID = ID
	relation.CreateDate = time.Now()

	status, errorSaving := db.SaveRelation(relation)
	if errorSaving != nil || !status {
		http.Error(rw, "Error saving the relation: "+errorSaving.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

// DeleteRelation : Delete a relation between two users
func DeleteRelation(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "The id param is required", http.StatusBadRequest)
		return
	}

	relation, _ := db.FindRelationByUserAndRelationUser(UserID, ID)
	if len(relation.UserID) < 1 {
		http.Error(rw, "The relation does not exist", http.StatusBadRequest)
		return
	}
	status, errorSaving := db.DeleteRelation(relation)
	if errorSaving != nil || !status {
		http.Error(rw, "Error deleting the relation: "+errorSaving.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

// GetRelation : get the relation between two users
func GetRelation(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "The id param is required", http.StatusBadRequest)
		return
	}
	relation, err := db.FindRelationByUserAndRelationUser(UserID, ID)
	if err != nil {
		http.Error(rw, "Error trying to get the relation: "+err.Error(), http.StatusBadRequest)
		return
	}
	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(relation)
}
