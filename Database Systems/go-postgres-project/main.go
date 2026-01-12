package main

import (
	"encoding/json"
	"fmt"
	"go-postgres/config"
	"go-postgres/model"
	"net/http"

	_ "github.com/lib/pq"
) 

func main(){
config.Connect()

http.HandleFunc("/users", userHandler)

defer http.ListenAndServe(":8081", nil) 
fmt.Println("http server running , 8081 port")
}



func userHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodGet:
		getUsers(w)
	}
}



func getUsers(w http.ResponseWriter){
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if  err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)

}