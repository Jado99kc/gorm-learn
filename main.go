package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jado99kc/gorm-learn/db"
	"github.com/jado99kc/gorm-learn/models"
	"github.com/jado99kc/gorm-learn/routes"
)

func main() {

	db.DbConnnection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	http.ListenAndServe(":3000", router)
}
