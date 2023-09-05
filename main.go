package main

import (
	api "answer-heck/api"
	db "answer-heck/db"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
		q, _  := url.ParseQuery(req.URL.RawQuery)
		var user db.User
		if q.Has("username") {
			user = api.GetUserFromUsername(q.Get("username"))

			userJson, err := json.Marshal(user)
			if err != nil {
				log.Fatal(err)
			}
			w.Write([]byte(string(userJson)))
		} else if q.Has("id") {
			userid, err := strconv.Atoi(q.Get("id"))
			if err != nil {
				return
			}
			user = api.GetUserFromId(userid)

			userJson, err := json.Marshal(user)
			if err != nil {
				log.Fatal(err)
			}
			w.Write([]byte(string(userJson)))
		} else {
			cookie, err := req.Cookie("session")
			if err != nil {
				if err != http.ErrNoCookie {
					log.Fatal(err)
				} else {
					return
				}
			}
			user = api.GetUserFromSession(cookie.Value)

			userJson, err := json.Marshal(user)
			if err != nil {
				log.Fatal(err)
			}
			w.Write([]byte(string(userJson)))
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
			session := api.PostSession(username, password)
			if len(session) == 0 {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.Write([]byte(session))
			}
		}
	}
}

// don't sue me meta!!!
func api_threads(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		posts := api.GetThreads()
		postsJson, err := json.Marshal(posts)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(string(postsJson)))
	}
}

func api_post(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		url := req.FormValue("url")
		title := req.FormValue("title")
		body := req.FormValue("body")

		cookie, err := req.Cookie("session")
		if err != nil {
			if err != http.ErrNoCookie {
				log.Fatal(err)
			} else {
				return
			}
		}

		api.PostPost(cookie.Value, url, title, body)
	} else if req.Method == http.MethodGet {
		q, _  := url.ParseQuery(req.URL.RawQuery)
		var post db.Post
		if q.Has("id") {
			postid, err := strconv.Atoi(q.Get("id"))
			if err != nil {
				return
			}
			post = api.GetPostFromId(postid)

			postJson, err := json.Marshal(post)
			if err != nil {
				log.Fatal(err)
			}
			w.Write([]byte(string(postJson)))
		}
	}
}

func main() {
	db.Connect();

	http.Handle("/", http.FileServer(http.Dir("./client-web/build")))
	http.HandleFunc(".", serve)
	http.HandleFunc("/login", serve)
	http.HandleFunc("/register", serve)
	http.HandleFunc("/submit", serve)
	http.HandleFunc("/post", serve)

	http.HandleFunc("/api/user", api_user)
	http.HandleFunc("/api/user/", api_user)
	http.HandleFunc("/api/session/", api_session)
	http.HandleFunc("/api/threads/", api_threads)
	http.HandleFunc("/api/post/", api_post)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}