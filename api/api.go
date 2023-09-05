package api

import (
    db "answer-heck/db"
)

func PostUser(username string, password string) {
    db.AddUser(username, password)
}

func GetUserFromId(id int) db.User {
    return db.GetUserFromId(id)
}

func GetUserFromUsername(username string) db.User {
    return db.GetUserFromUsername(username)
}

func GetUserFromSession(session string) db.User {
    return db.GetUserFromSession(session)
}

func PostSession(username string, password string) string {
    return db.NewSession(username, password)
}

func GetThreads() []db.Post {
    return db.GetPostsFromParent(0)
}

func GetChildrenFromParentId(parentid int) []db.Post {
    var children = make([]db.Post, 0, 8)

    var first_decendants = db.GetPostsFromParent(parentid)
    for _,v := range(first_decendants) {
        children = append(children, v)
        children = append(children, GetChildrenFromParentId(v.Id)...)
    }
    return children
}

func PostPost(session string, url string, title string, body string, parentid int) {
    user := db.GetUserFromSession(session)
    if user.Id != 0 {
        db.AddPost(user.Id, url, title, body, parentid)
    }
}

func GetPostFromId(id int) db.Post {
    return db.GetPostFromId(id)
}