package main

import (
    api "answer-heck/api"
    db "answer-heck/db"
    "log"
    "net/http"
    "os"
    "strconv"
    "github.com/gorilla/mux"
)

func serve(w http.ResponseWriter, req *http.Request) {
    data, err := os.ReadFile("./client-web/build/index.html")
    if err != nil {
        log.Fatal(err)
    }

    w.Write(data)
}

/**
 * REST API blueprint
 *  GET  /api/user/{user_id}            returns user information
 *  POST /api/user                      creates new user with un/pw form
 *  POST /api/session                   creates new session with un/pw form and returns session key
 *  GET  /api/post/{post_id}            get post with id={post_id}
 *  POST /api/post                      creates new post
 *  POST /api/post/{post_id}            creates new post as a reply to {post_id}
 *  GET  /api/threads                   get posts with no parent
 *  GET  /api/post/{post_id}/children   get posts with parent={post_id}
 */

func GetUser(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)

    userId := vars["user_id"]
    userIdInteger, err := strconv.Atoi(userId)
    if err != nil {
        return
    }

    w.Write([]byte(api.GetUserFromId(userIdInteger)))
}

func GetUserFromSession(w http.ResponseWriter, req *http.Request) {
    sessionCookie, err := req.Cookie("session")
    if err != nil {
        return
    }

    w.Write([]byte(api.GetUserFromSession(sessionCookie.Value)))
}

func PostUser(w http.ResponseWriter, req *http.Request) {
    username := req.FormValue("username")
    password := req.FormValue("password")

    if len(username) == 0 || len(password) == 0 {
        w.WriteHeader(http.StatusNotFound)
    } else {
        api.PostUser(username, password)
    }
}

func PostSession(w http.ResponseWriter, req *http.Request) {
    username := req.FormValue("username")
    password := req.FormValue("password")

    if len(username) == 0 || len(password) == 0 {
        w.WriteHeader(http.StatusNotFound)
    } else {
        session := api.PostSession(username, password)
        if len(session) == 0 {
            w.WriteHeader(http.StatusNotFound)
        } else {
            w.Header().Set("Set-Cookie","session="+session+"; HttpOnly; Path=/;")
            w.WriteHeader(http.StatusOK)
        }
    }
}

func GetPost(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)

    postId := vars["post_id"]
    postIdInteger, err := strconv.Atoi(postId)
    if err != nil {
        return
    }

    w.Write([]byte(api.GetPostFromId(postIdInteger)))
}

func PostPost(w http.ResponseWriter, req *http.Request) {
    url := req.FormValue("url")
    title := req.FormValue("title")
    body := req.FormValue("body")

    sessionCookie, err := req.Cookie("session")
    if err != nil {
        if err != http.ErrNoCookie {
            log.Fatal(err)
        } else {
            log.Println("No cookie")
            return
        }
    }

    api.PostPost(sessionCookie.Value, url, title, body, 0)
}

func PostReply(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)

    postId := vars["post_id"]
    postIdInteger, err := strconv.Atoi(postId)
    if err != nil {
        return
    }

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

    api.PostPost(cookie.Value, url, title, body, postIdInteger)
}

func GetThreads(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte(api.GetThreads()))
}

func GetPostChildren(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)

    postId := vars["post_id"]
    postIdInteger, err := strconv.Atoi(postId)
    if err != nil {
        return
    }

    w.Write([]byte(api.GetChildrenFromParentId(postIdInteger)))
}

func main() {
    db.Connect();

    r := mux.NewRouter()

    // all of these handle funcs return the same HTML content. the router on the client is supposed to take care of rendering the proper page so on the backend we just expose the endpoints
    r.HandleFunc(".", serve)
    r.HandleFunc("/login", serve)
    r.HandleFunc("/register", serve)
    r.HandleFunc("/submit", serve)
    r.HandleFunc("/post", serve)

    r.HandleFunc("/api/user", GetUserFromSession).Methods("GET")
    r.HandleFunc("/api/user/{user_id}", GetUser).Methods("GET")
    r.HandleFunc("/api/user", PostUser).Methods("POST")
    r.HandleFunc("/api/session", PostSession).Methods("POST")
    r.HandleFunc("/api/post/{post_id}", GetPost).Methods("GET")
    r.HandleFunc("/api/post", PostPost).Methods("POST")
    r.HandleFunc("/api/post/{post_id}", PostReply).Methods("POST")
    r.HandleFunc("/api/threads", GetThreads).Methods("GET")
    r.HandleFunc("/api/post/{post_id}/children", GetPostChildren).Methods("GET")

    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./client-web/build")))

    err := http.ListenAndServe(":8080", r)
    if err != nil {
        log.Fatal(err)
    }
}