package db

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

var database *sql.DB

type User struct {
	Username string
	Password string
}

type Session struct {
	Cookie string
	Username string
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
		} else {
			// we now know the user does not exist, so insert it into the table
			_, err = database.Exec("INSERT INTO USER (username, password) VALUES (?, ?)", username, password)
			if err != nil {
				log.Fatal(err)
			}
		}
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

func AddSession(username string, password string) string {
	user := GetUser(username)

	// check if user exists
	if len(user.Username) == 0 {
		return ""
	}

	// check if password matches
	if user.Password != password {
		return ""
	}

	var err error

	// delete the previous session
	_, err = database.Exec("DELETE FROM SESSION WHERE username = ?", username)
	if err != nil {
		if err != sql.ErrNoRows {
			 log.Fatal(err)
		}
	}

	cookie := generateSessionCookie(64)
	_, err = database.Exec("INSERT INTO SESSION (cookie, username) VALUES (?, ?)", cookie, username)
	if err != nil {
		log.Fatal(err)
	}

	return cookie
}

func GetSession(cookie string) Session {
	var session Session

	row := database.QueryRow("SELECT * FROM SESSION WHERE cookie = ?", cookie)
	if err := row.Scan(&session.Cookie, &session.Username); err != nil {
		if err != sql.ErrNoRows {
			 log.Fatal(err)
		}
	}

	return session
}

var charSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func generateSessionCookie(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
    b := make([]rune, n)
    for i := range b {
        b[i] = charSet[rand.Intn(len(charSet))]
    }
    return string(b)
}