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
	Id int 			`json:"id"`
	Username string `json:"username"`
	password string
	session string
}

type Post struct {
	Id int			`json:"id"`
	UserId string	`json:"userid"`
	Title string	`json:"title"`
	Url string		`json:"url"`
	Body string		`json:"body"`
	Score int		`json:"score"`
	ParentId int	`json:"parentid"`
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
	// UPDATE: we're using user ids now but I still do not want duplicate usernames to exist
	err := database.QueryRow("SELECT * FROM USER WHERE username = ?", username).Scan(&user.Id, &user.Username, &user.password, &user.session)
	if err != nil {
		if err != sql.ErrNoRows {
			 log.Fatal(err)
		} else {
			// we now know the user does not exist, so insert it into the table
			_, err = database.Exec("INSERT INTO USER (username, password, session) VALUES (?, ?, ?)", username, password, generateSessionCookie(64))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func GetUserFromId(id int) User {
	var user User

	row := database.QueryRow("SELECT * FROM USER WHERE id = ?", id)
	if err := row.Scan(&user.Id, &user.Username, &user.password, &user.session); err != nil {
		if err != sql.ErrNoRows {
			 log.Fatal(err)
		}
	}

	return user
}

func GetUserFromUsername(username string) User {
	var user User

	row := database.QueryRow("SELECT * FROM USER WHERE username = ?", username)
	if err := row.Scan(&user.Id, &user.Username, &user.password, &user.session); err != nil {
		if err != sql.ErrNoRows {
			 log.Fatal(err)
		}
	}

	return user
}

func GetUserFromSession(session string) User {
	var user User

	row := database.QueryRow("SELECT * FROM USER WHERE session = ?", session)
	if err := row.Scan(&user.Id, &user.Username, &user.password, &user.session); err != nil {
		if err != sql.ErrNoRows {
			 log.Fatal(err)
		}
	}

	return user
}

func NewSession(username string, password string) string {
	user := GetUserFromUsername(username)

	// check if user exists
	if len(user.Username) == 0 {
		return ""
	}

	// check if password matches
	if user.password != password {
		return ""
	}

	var err error

	session := generateSessionCookie(64)
	// delete the previous session
	_, err = database.Exec("UPDATE USER SET session = ? WHERE username = ?", session, username)
	if err != nil {
		if err != sql.ErrNoRows {
			 log.Fatal(err)
		}
	}

	return session
}

func GetPostsFromParent(parent int) []Post {
	posts := make([]Post, 0, 8)

	rows,err := database.Query("SELECT * FROM POST WHERE parentid = ?", parent)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		if err = rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Url, &post.Body, &post.Score, &post.ParentId); err != nil {
			log.Fatal(err)
		}

		posts = append(posts, post)
	}

	return posts
}

func AddPost(url string, title string, body string) {
	_, err := database.Exec("INSERT INTO POST (userid, title, url, body, score, parentid) VALUES (?, ?, ?, ?, ?, ?)", 0, title, url, body, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
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
