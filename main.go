package main

import (
	mux2 "github.com/gorilla/mux"
	"gorm/handlers"
	"log"
	"net/http"
)

func main() {

	mux := mux2.NewRouter()

	mux.HandleFunc("/", handlers.Root).Methods("GET")
	//mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	//mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	//mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	//mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	//mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", mux))

}
