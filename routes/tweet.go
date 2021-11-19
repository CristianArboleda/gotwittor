package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/CristianArboleda/gotwittor/db"
	"github.com/CristianArboleda/gotwittor/models"
)

// SaveTweet : save a new Tweet
func SaveTweet(rw http.ResponseWriter, r *http.Request) {
	var tw models.Tweet
	err := json.NewDecoder(r.Body).Decode(&tw)
	if err != nil {
		http.Error(rw, "Error in the body request: "+err.Error(), 400)
		return
	}
	if len(tw.Message) == 0 {
		http.Error(rw, "The message field is required", 400)
		return
	}

	tw.UserID = UserID
	tw.CreateDate = time.Now()

	_, status, err := db.SaveTweet(tw)

	if err != nil {
		http.Error(rw, "Error saving the tweet: "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(rw, "Error saving the tweet", 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

// DeleteTweet : delete a Tweet
func DeleteTweet(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "The id param is required", http.StatusBadRequest)
		return
	}
	err := db.DeleteTweet(ID, UserID)
	if err != nil {
		http.Error(rw, "Error trying to delete the record: "+err.Error(), http.StatusBadRequest)
		return
	}
	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}

// GetTweets : get all tweets of a user
func GetTweets(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "The id param is required", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(rw, "The page param is required", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(rw, "The page param must be a numeric value", http.StatusBadRequest)
		return
	}
	pag := int64(page)

	results, status := db.FindTweetsByUserID(UserID, pag)
	if !status {
		http.Error(rw, "Error trying to get the record: "+err.Error(), http.StatusBadRequest)
		return
	}
	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(results)
}

// GetFollowersTweets : get all followers tweets
func GetFollowersTweets(rw http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(rw, "The page param is required", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(rw, "The page param must be a numeric value", http.StatusBadRequest)
		return
	}
	pag := int64(page)

	results, status := db.FindFollowersTweets(UserID, pag)
	if !status {
		http.Error(rw, "Error trying to get the follower tweets", http.StatusBadRequest)
		return
	}

	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(results)
}
