package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// type Welcome struct {
	// Sale string
	// Time string
// }

var db *sql.DB

func main() {
	welcome := Welcome{"Sale Begins Now", time.Now().Format(time.Stamp)}
	template := template.Must(template.ParseFiles("template/studentportal.html"))
	var err error
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/studentms")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/insert", inserthandler)
		if sale := r.FormValue("sale"); sale != "" {
			welcome.Sale = sale
		}
		if err := template.ExecuteTemplate(w, "studentportal.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":9000", nil)) //port 8000 for login and port 9000 for portals
}
