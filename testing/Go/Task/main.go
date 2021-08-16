package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// "log"
	// "net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Port = ":8080"
)

type Teacher struct {
	Name     string
	Email    string
	Password string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234"
	dbName := "studentms"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

//var tmpl = template.HTMLEscaper("C:/Program Files/Go/src/SMS/Go/Task/form/*"))

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		insForm, err := db.Prepare("INSERT INTO Employee(name, email, password) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, email)
		log.Println("INSERT: Name: " + name + " | email: " + email + " | password: " + password)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 3306)
}

func main() {
	//template.ParseFiles("form/index.html")
	log.Println("Server started on: http://localhost/127.0.0.1:3306")
	http.HandleFunc("/insert", Insert)
	fmt.Println("Welcome TO SMS")

	// db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/studentms")
	// if err != nil {
	// 	panic(err.Error())
}
