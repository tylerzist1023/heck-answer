package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func serve(w http.ResponseWriter, req *http.Request) {
	data, err := os.ReadFile("./client-web/build/index.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}

func api(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%v", req.Method)
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")

		if len(username) == 0 || len(password) == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Write([]byte(username))
			w.Write([]byte(password))
		}		
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./client-web/build")))
	http.HandleFunc(".", serve)
	http.HandleFunc("/login", serve)
	http.HandleFunc("/register", serve)

	http.HandleFunc("/api", api)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}