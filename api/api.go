package api

import db "answer-heck/db"

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

func PostPost(url string, title string, body string) {
	db.AddPost(url, title, body)
}