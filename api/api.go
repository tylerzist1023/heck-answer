package api

import db "answer-heck/db"

func PostUser(username string, password string) {
	db.AddUser(username, password)
}

func GetUser(username string) db.User {
	return db.GetUser(username)
}