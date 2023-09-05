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

func PostPost(session string, url string, title string, body string) {
	user := db.GetUserFromSession(session)
	if user.Id != 0 {
		db.AddPost(user.Id, url, title, body)
	}
}

func GetPostFromId(id int) db.Post {
	return db.GetPostFromId(id)
}