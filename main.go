package main

import (
	api "answer-heck/api"
	db "answer-heck/db"
	"log"
	"net/http"
	"os"
	"strings"
)

func serve(w http.ResponseWriter, req *http.Request) {
	data, err := os.ReadFile("./client-web/build/index.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}

func api_user(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")

		if len(username) == 0 || len(password) == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			api.PostUser(username, password)
			w.Write([]byte(username))
			w.Write([]byte(password))
		}		
	} else if req.Method == http.MethodGet {
		urlPart := strings.Split(req.URL.Path, "/")
		username := urlPart[3]

		user := api.GetUser(username)
		if len(user.Username) == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Write([]byte(user.Username))
		}
	}
}

func main() {
	db.Connect();

	http.Handle("/", http.FileServer(http.Dir("./client-web/build")))
	http.HandleFunc(".", serve)
	http.HandleFunc("/login", serve)
	http.HandleFunc("/register", serve)

	http.HandleFunc("/api/user/", api_user)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}