package api

import (
    db "answer-heck/db"
    "encoding/json"
    "log"
)

func getJson(obj any) string {
    objJson, err := json.Marshal(obj)
    if err != nil {
        log.Fatal(err)
    }
    return string(objJson)
}

func PostUser(username string, passwordHash string) {
    db.AddUser(username, passwordHash)
}

func GetUserFromId(id int) string {
    return getJson(db.GetUserFromId(id))
}

func GetUserFromUsername(username string) string {
    return getJson(db.GetUserFromUsername(username))
}

func GetUserFromSession(session string) string {
    return getJson(db.GetUserFromSession(session))
}

func PostSession(username string, password string) string {
    return db.NewSession(username, password)
}

func GetThreads() string {
    return getJson(db.GetPostsFromParent(0))
}

func GetChildrenFromParentId(parentid int) string {
    var children = make([]db.Post, 0, 8)
    children = append(children, getChildrenFromParentId(parentid, 0)...)
    return getJson(children)
}

func getChildrenFromParentId(parentid int, depth int) []db.Post {
    var children = make([]db.Post, 0, 8)

    var first_decendants = db.GetPostsFromParent(parentid)
    for _,v := range(first_decendants) {
        children = append(children, v)
        children_of_parentid := getChildrenFromParentId(v.Id, depth+1)
        children = append(children, children_of_parentid...)
        if depth+1 >= 10 {
            break
        }
    }
    return children
}

func PostPost(session string, url string, title string, body string, parentid int) {
    user := db.GetUserFromSession(session)
    if user.Id != 0 {
        db.AddPost(user.Id, url, title, body, parentid)
    }
}

func GetPostFromId(id int) string {
    return getJson(db.GetPostFromId(id))
}