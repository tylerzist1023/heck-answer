package db

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

var database *sql.DB

type User struct {
	Username string
	Password string
}

func Connect() {
	cfg := mysql.Config {
		User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "ANSWERHECK",
	}

	var err error
	database, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database!")
}

func AddUser(username string, password string) {
	var user User

	// check if the user exists
	err := database.QueryRow("SELECT * FROM USER WHERE username = ?", username).Scan(&user.Username, &user.Password)
	if err != nil {
		if err != sql.ErrNoRows {
			 log.Fatal(err)
		}
	}

	// we now know the user does not exist, so insert it into the table
	_, err = database.Exec("INSERT INTO USER (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUser(username string) User {
	var user User

	row := database.QueryRow("SELECT * FROM USER WHERE username = ?", username)
	if err := row.Scan(&user.Username, &user.Password); err != nil {
		if err != sql.ErrNoRows {
			 log.Fatal(err)
		}
	}

	return user
}