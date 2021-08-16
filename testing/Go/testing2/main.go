package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234"
	dbName := "studentms"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Student struct {
	Token      string `json:"token"`
	Stdid      string `json:"stdid"`
	Stdname    string `json:"stdname"`
	Stdemail   string `json:"stdemail"`
	Stdclass   string `json:"stdclass"`
	Stdage     string `json:"stdage"`
	Stdcity    string `json:"stdcity"`
	Stdsubject string `json:"stdsubject"`
}

type Teacher struct {
	Tecid       string `json:"tecid"`
	Tecname     string `json:"tecname"`
	Tecemail    string `json:"tecemail"`
	Tecpassword string `json:"tecpassword"`
}

func Createteacher(w http.ResponseWriter, r *http.Request) {
	var teacher Teacher
	db := dbConn()
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		panic(err)
	}
	query := "insert into teacher (tecname,tecemail,tecpassword) values ( '" + teacher.Tecname + "' ,'" + teacher.Tecemail + "','" + teacher.Tecpassword + "')"
	querychk := "select * from teacher where tecemail='" + teacher.Tecemail + "'"
	resultchk, err := db.Query(querychk)
	if err != nil {
		panic(err)
	}
	defer resultchk.Close()
	var teccheck Teacher
	for resultchk.Next() {
		err := resultchk.Scan(&teccheck.Tecid, &teccheck.Tecname, &teccheck.Tecemail, &teccheck.Tecpassword)
		if err != nil {
			panic(err.Error())
		}
	}
	if teacher.Tecemail != teccheck.Tecemail {
		result, err := db.Exec(query)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
		defer db.Close()
		w.Header().Set("Content-Type", "application/json") //show meaning full massage
		w.WriteHeader(http.StatusCreated)                  //status code
		json.NewEncoder(w).Encode(teacher)                 // print values in post man to cheack what vvalue store
		json.NewEncoder(w).Encode(valmess)
	} else {
		var insmess = "Email already Registerd"
		json.NewEncoder(w).Encode(insmess)
	}
}
func Teacherlogin(w http.ResponseWriter, r *http.Request) {
	var teacher Teacher
	gtoken := "Token 123456"
	logresult := "Successfull Login"
	db := dbConn()
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		panic(err)
	}
	query := "SELECT * FROM teacher WHERE tecemail= '" + teacher.Tecemail + "' AND tecpassword= '" + teacher.Tecpassword + "'"
	result, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	var teccheck Teacher
	for result.Next() {
		err := result.Scan(&teccheck.Tecid, &teccheck.Tecname, &teccheck.Tecemail, &teccheck.Tecpassword)
		if err != nil {
			panic(err.Error())
		}
	}
	if teacher.Tecemail == teccheck.Tecemail && teacher.Tecpassword == teccheck.Tecpassword {
		json.NewEncoder(w).Encode(gtoken)
		json.NewEncoder(w).Encode(logresult)
		json.NewEncoder(w).Encode(teccheck) // print values in post man to cheack what vvalue store
	} else {
		var insmess = "Invalid Email or Password"
		json.NewEncoder(w).Encode(insmess)
	}
	defer db.Close()
	w.Header().Set("Content-Type", "application/json") //show meaning full massage
	w.WriteHeader(http.StatusFound)                    //status code
}
func Createstudent(w http.ResponseWriter, r *http.Request) {

	var student Student
	db := dbConn()
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		panic(err)
	}
	query, ttoken := "insert into student (stdname,stdemail,stdclass,stdage,stdcity,stdsubject) values ( '"+student.Stdname+"' ,'"+student.Stdemail+"','"+student.Stdclass+"','"+student.Stdage+"','"+student.Stdcity+"','"+student.Stdsubject+"')", student.Token
	querychk := "select * from student where stdemail ='" + student.Stdemail + "'"
	if ttoken == vtoken {
		resultchk, err := db.Query(querychk)
		if err != nil {
			panic(err)
		}
		defer resultchk.Close()
		var stdcheck Student
		for resultchk.Next() {
			err := resultchk.Scan(&stdcheck.Stdid, &stdcheck.Stdname, &stdcheck.Stdemail, &stdcheck.Stdclass, &stdcheck.Stdage, &stdcheck.Stdcity, &stdcheck.Stdsubject)
			if err != nil {
				panic(err.Error())
			}
		}
		if student.Stdemail != stdcheck.Stdemail {
			result, err := db.Exec(query)
			if err != nil {
				panic(err)
			}
			fmt.Println(result)
			defer db.Close()
			w.Header().Set("Content-Type", "application/json") //show meaning full massage
			w.WriteHeader(http.StatusCreated)                  //status code
			json.NewEncoder(w).Encode(student)                 // print values in post man to cheack what vvalue store
			json.NewEncoder(w).Encode(valmess)
		} else {
			var insmess = "Email already Registerd"
			json.NewEncoder(w).Encode(insmess)

		}
	} else {
		json.NewEncoder(w).Encode(invmess)

	}
}
func Updatestudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	db := dbConn()
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		panic(err)
	}
	query, ttoken := "UPDATE student SET stdname='"+student.Stdname+"', stdclass='"+student.Stdclass+"', stdcity='"+student.Stdcity+"', stdage='"+student.Stdage+"', stdsubject='"+student.Stdsubject+"' WHERE stdemail='"+student.Stdemail+"'", student.Token
	querychk := "select * from student where stdemail ='" + student.Stdemail + "'"
	if ttoken == vtoken {
		resultchk, err := db.Query(querychk)
		if err != nil {
			panic(err)
		}
		defer resultchk.Close()
		var stdcheck Student
		for resultchk.Next() {
			err := resultchk.Scan(&stdcheck.Stdid, &stdcheck.Stdname, &stdcheck.Stdemail, &stdcheck.Stdclass, &stdcheck.Stdage, &stdcheck.Stdcity, &stdcheck.Stdsubject)
			if err != nil {
				panic(err.Error())
			}
		}
		if student.Stdemail == stdcheck.Stdemail {
			result, err := db.Exec(query)
			if err != nil {
				panic(err)
			}
			fmt.Println(result)
			defer db.Close()
			w.Header().Set("Content-Type", "application/json") //show meaning full massage
			w.WriteHeader(http.StatusFound)                    //status code
			json.NewEncoder(w).Encode(student)                 // print values in post man to cheack what vvalue store
			json.NewEncoder(w).Encode(valmess)
		} else {
			var insmess = "Email Not Valid"
			json.NewEncoder(w).Encode(insmess)
		}
	} else {
		json.NewEncoder(w).Encode(invmess)
	}
}
func Deletestudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	db := dbConn()
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		panic(err)
	}
	query, ttoken := "DELETE FROM student WHERE stdemail='"+student.Stdemail+"'", student.Token
	querychk := "select * from student where stdemail ='" + student.Stdemail + "'"
	if ttoken == vtoken {
		resultchk, err := db.Query(querychk)
		if err != nil {
			panic(err)
		}
		defer resultchk.Close()
		var stdcheck Student
		for resultchk.Next() {
			err := resultchk.Scan(&stdcheck.Stdid, &stdcheck.Stdname, &stdcheck.Stdemail, &stdcheck.Stdclass, &stdcheck.Stdage, &stdcheck.Stdcity, &stdcheck.Stdsubject)
			if err != nil {
				panic(err.Error())
			}
		}
		if student.Stdemail == stdcheck.Stdemail {
			result, err := db.Exec(query)
			if err != nil {
				panic(err)
			}
			fmt.Println(result)
			defer db.Close()
			w.Header().Set("Content-Type", "application/json") //show meaning full massage
			w.WriteHeader(http.StatusAccepted)                 //status code
			json.NewEncoder(w).Encode(student)                 // print values in post man to cheack what vvalue store
			json.NewEncoder(w).Encode(valmess)
		} else {
			var insmess = "No Record found on this Email"
			json.NewEncoder(w).Encode(insmess)
		}
	} else {
		json.NewEncoder(w).Encode(invmess)
	}
}
func Showstudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	db := dbConn()
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		panic(err)
	}
	query, ttoken := "SELECT * FROM student WHERE stdemail='"+student.Stdemail+"'", student.Token
	querychk := "select * from student where stdemail ='" + student.Stdemail + "'"
	if ttoken == vtoken {
		resultchk, err := db.Query(querychk)
		if err != nil {
			panic(err)
		}
		defer resultchk.Close()
		var stdcheck Student
		for resultchk.Next() {
			err := resultchk.Scan(&stdcheck.Stdid, &stdcheck.Stdname, &stdcheck.Stdemail, &stdcheck.Stdclass, &stdcheck.Stdage, &stdcheck.Stdcity, &stdcheck.Stdsubject)
			if err != nil {
				panic(err.Error())
			}
		}
		if student.Stdemail == stdcheck.Stdemail {
			result, err := db.Query(query)
			if err != nil {
				panic(err)
			}
			defer result.Close()
			var stud Student
			for result.Next() {
				err := result.Scan(&stud.Stdid, &stud.Stdname, &stud.Stdemail, &stud.Stdclass, &stud.Stdage, &stud.Stdcity, &stud.Stdsubject)
				if err != nil {
					panic(err.Error())
				}
			}
			fmt.Println(stud)
			defer db.Close()
			w.Header().Set("Content-Type", "application/json") //show meaning full massage
			w.WriteHeader(http.StatusFound)                    //status code
			json.NewEncoder(w).Encode(stud)                    // print values in post man to cheack what vvalue store
			json.NewEncoder(w).Encode(valmess)
		} else {
			var insmess = "No Record found on this Email"
			json.NewEncoder(w).Encode(insmess)
		}
	} else {
		json.NewEncoder(w).Encode(invmess)
	}
}

var vtoken string = "123456"
var invmess = "Invalid Token/Enter Token"
var valmess = "Done with your task"

func main() {

	http.HandleFunc("/createteacher", Createteacher) //method post working
	http.HandleFunc("/teacherlogin", Teacherlogin)   //method check working
	http.HandleFunc("/createstudent", Createstudent) //method post working
	http.HandleFunc("/updatestudent", Updatestudent) //method update working issue
	http.HandleFunc("/deletestudent", Deletestudent) //method delete working
	http.HandleFunc("/showstudent", Showstudent)     //method get working
	fmt.Println(http.ListenAndServe(":9083", nil))
}
