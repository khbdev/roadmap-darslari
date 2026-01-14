package main

import (
	"encoding/json"
	"fmt"
	"go-postgres/config"
	"go-postgres/model"
	"log"
	"net/http"
)

func main(){
	config.Connect()

	http.HandleFunc("/users", UserHandler)

 defer	http.ListenAndServe(":8081", nil)

 fmt.Println("Server Running :8081")

}


func UserHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		GetUSers(w)
	case http.MethodPost:
		PostUsers(w, r)
	}
}


func GetUSers(w http.ResponseWriter){
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if  err != nil {
		log.Fatal(err)
	}

	var users []model.User
	for rows.Next(){
		var u model.User
		if  err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}
  json.NewEncoder(w).Encode(users)
}

func PostUsers(w http.ResponseWriter, r *http.Request){
	var u model.User
if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
    http.Error(w, "invalid json", http.StatusBadRequest)
    return
}


	err := config.DB.QueryRow("INSERT INTO users(name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
	if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
   w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(u)
}