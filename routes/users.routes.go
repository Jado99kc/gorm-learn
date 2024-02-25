package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jado99kc/gorm-learn/db"
	"github.com/jado99kc/gorm-learn/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)

	json.NewEncoder(w).Encode(users)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getting user"))

}
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)

}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting user"))

}
