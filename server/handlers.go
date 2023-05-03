package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not Allowed\n")
			return
		}
		fmt.Fprintf(w, "Hi there!\n")
	})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usrs)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	usr := &User{}
	err := json.NewDecoder(r.Body).Decode(usr)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	usrs = append(usrs, usr)
	fmt.Fprintf(w, "User added successfully\n")
}
