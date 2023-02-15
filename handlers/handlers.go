package handlers

import (
	"encoding/json"
	mux2 "github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}

	db.Database.Find(&users)
	sendData(rw, http.StatusOK, users)

}

func GetUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, http.StatusOK, user)
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&user)
		sendData(rw, 200, user)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var userId int64
	if old_user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		userId = old_user.Id

		user := models.User{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&user); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			user.Id = userId
			db.Database.Save(&user)
			sendData(rw, 200, user)
		}
	}

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		db.Database.Delete(&user)
		sendData(rw, 200, user)
	}

}

func getUserById(r *http.Request) (models.User, *gorm.DB) {
	vars := mux2.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.User{}

	if err := db.Database.First(&user, userId); err.Error != nil {
		return user, err
	}

	return user, nil
}
