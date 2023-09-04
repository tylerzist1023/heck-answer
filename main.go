package main

import (
	api "answer-heck/api"
	db "answer-heck/db"
	"fmt"
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

func api_session(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")

		if len(username) == 0 || len(password) == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			cookie := api.PostSession(username, password)
			// fmt.Printf("%s\n", cookie)
			if len(cookie) == 0 {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.Write([]byte(cookie))
			}
		}
	} else if req.Method == http.MethodGet {
		cookieObject, err := req.Cookie("session")
		if err != nil {
			if err != http.ErrNoCookie {
				log.Fatal(err)
			} else {
				return
			}
		}
		session := api.GetSession(cookieObject.Value)
		fmt.Printf("username: %s\n", session.Username)
		w.Write([]byte(session.Username))
	}
}

func main() {
	db.Connect();

	http.Handle("/", http.FileServer(http.Dir("./client-web/build")))
	http.HandleFunc(".", serve)
	http.HandleFunc("/login", serve)
	http.HandleFunc("/register", serve)

	http.HandleFunc("/api/user/", api_user)
	http.HandleFunc("/api/session/", api_session)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}