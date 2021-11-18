package routes

import (
	"github.com/CristianArboleda/gotwittor/db"
	"github.com/CristianArboleda/gotwittor/models"
	"io"
	"net/http"
	"os"
	"strings"
)

var bannersPath string = "uploads/banners/"
var avatarsPath string = "uploads/avatars/"

// AddAvatar : Upload the avatar image and update the user
func AddAvatar(rw http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(rw, "The avatar is required: "+err.Error(), http.StatusBadRequest)
		return
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	var filePath string = avatarsPath + UserID + "." + extension

	localFile, errorOpenFile := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if errorOpenFile != nil {
		http.Error(rw, "Error creating file in the system: "+errorOpenFile.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(localFile, file)
	if err != nil {
		http.Error(rw, "Error writing file in the system: "+err.Error(), http.StatusBadRequest)
		return
	}
	var us models.User
	var status bool
	us.Avatar = UserID + "." + extension
	status, err = db.UpdateUser(us, UserID)
	if err != nil || !status {
		http.Error(rw, "Error updating the avatar in the db: "+err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}

func GetAvatar(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "The id param is required", http.StatusBadRequest)
		return
	}

	profile, err := db.FindUserById(ID)
	if err != nil {
		http.Error(rw, "Error trying to get the user: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, errOpen := os.Open(avatarsPath + profile.Avatar)
	if errOpen != nil {
		http.Error(rw, "Avatar not found : "+errOpen.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(rw, file)
	if err != nil {
		http.Error(rw, "Error return the avatar : "+err.Error(), http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

// AddBanner : Upload the banner image and update the user
func AddBanner(rw http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(rw, "The banner is required: "+err.Error(), http.StatusBadRequest)
		return
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	var filePath string = bannersPath + UserID + "." + extension

	localFile, errorOpenFile := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if errorOpenFile != nil {
		http.Error(rw, "Error creating file in the system: "+errorOpenFile.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(localFile, file)
	if err != nil {
		http.Error(rw, "Error writing file in the system: "+err.Error(), http.StatusBadRequest)
		return
	}
	var us models.User
	var status bool
	us.Banner = UserID + "." + extension
	status, err = db.UpdateUser(us, UserID)
	if err != nil || !status {
		http.Error(rw, "Error updating the banner in the db: "+err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}

func GetBanner(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "The id param is required", http.StatusBadRequest)
		return
	}

	profile, err := db.FindUserById(ID)
	if err != nil {
		http.Error(rw, "Error trying to get the user: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, errOpen := os.Open(bannersPath + profile.Banner)
	if errOpen != nil {
		http.Error(rw, "Banner not found : "+errOpen.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(rw, file)
	if err != nil {
		http.Error(rw, "Error return the banner : "+err.Error(), http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}
