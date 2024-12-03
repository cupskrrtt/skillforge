package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("views/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error parsing template: %v", err)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
		}
	}).Methods("GET")

	r.HandleFunc("api/auth/login", func(w http.ResponseWriter, r *http.Request) {
		var user UserLogin
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			log.Printf("Error decoding JSON: %v", err)
			return
		}

		fmt.Fprintf(w, "email: %s, password: %s", user.Email, user.Password)
	}).Methods("POST")

	log.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
