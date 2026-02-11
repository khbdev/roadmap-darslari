package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UsersHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Userslarga salom :)")
}

func UserSwtich(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		UsersHandler(w)
	}
}

var port string = "8080"

func main() {
	http.HandleFunc("/users", UserSwtich)

	defer http.ListenAndServe(port, nil)
	fmt.Printf("Server running %s", port)
}
