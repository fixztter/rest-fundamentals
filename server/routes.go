package server

import (
	"fmt"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetUsers(w, r)
		case http.MethodPost:
			AddUser(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not Allowed\n")
			return
		}
	})
}
