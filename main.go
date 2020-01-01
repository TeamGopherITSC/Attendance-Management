package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"

)

type Student struct {
	Id       int
	Name     string
	Dept     string
	Batch    int
	Semester int
	Email    string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "attmgmt"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseFiles("index.html", "reset.html", "signup.html"))
var tmpl1 = template.Must(template.ParseGlob("admin/*"))
var tmpl2 = template.Must(template.ParseGlob("student/*"))
var tmpl3 = template.Must(template.ParseGlob("teacher/*"))
var tmpl4 = template.Must(template.ParseGlob("template/*"))

func Reset(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "reset.html", nil)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "signup.html", nil)
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	tmpl1.ExecuteTemplate(w, "index.html", nil)
}

func StudentIndex(w http.ResponseWriter, r *http.Request) {
	tmpl2.ExecuteTemplate(w, "index.html", nil)
}

func TeacherIndex(w http.ResponseWriter, r *http.Request) {
	tmpl3.ExecuteTemplate(w, "index.html", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
	// db := dbConn()
	// selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	// if err != nil {
	//     panic(err.Error())
	// }
	// emp := Employee{}
	// res := []Employee{}
	// for selDB.Next() {
	//     var id int
	//     var name, city string
	//     err = selDB.Scan(&id, &name, &city)
	//     if err != nil {
	//         panic(err.Error())
	//     }
	//     emp.Id = id
	//     emp.Name = name
	//     emp.City = city
	//     res = append(res, emp)
	// }
	// tmpl.ExecuteTemplate(w, "Index", res)
	// defer db.Close()
}

// func Show(w http.ResponseWriter, r *http.Request) {
//     db := dbConn()
//     nId := r.URL.Query().Get("id")
//     selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
//     if err != nil {
//         panic(err.Error())
//     }
//     emp := Employee{}
//     for selDB.Next() {
//         var id int
//         var name, city string
//         err = selDB.Scan(&id, &name, &city)
//         if err != nil {
//             panic(err.Error())
//         }
//         emp.Id = id
//         emp.Name = name
//         emp.City = city
//     }
//     tmpl.ExecuteTemplate(w, "Show", emp)
//     defer db.Close()
// }

// func New(w http.ResponseWriter, r *http.Request) {
//     tmpl.ExecuteTemplate(w, "New", nil)
// }

// func Edit(w http.ResponseWriter, r *http.Request) {
//     db := dbConn()
//     nId := r.URL.Query().Get("id")
//     selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
//     if err != nil {
//         panic(err.Error())
//     }
//     emp := Employee{}
//     for selDB.Next() {
//         var id int
//         var name, city string
//         err = selDB.Scan(&id, &name, &city)
//         if err != nil {
//             panic(err.Error())
//         }
//         emp.Id = id
//         emp.Name = name
//         emp.City = city
//     }
//     tmpl.ExecuteTemplate(w, "Edit", emp)
//     defer db.Close()
// }

//INSERT USED TO ADD INFORMATION TO NEW USERS

func Login(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	

	if r.Method == "POST" {
		usname := r.FormValue("username")
		pass := r.FormValue("password")
		etype := r.FormValue("type")
		selDb, err := db.Query("SELECT * FROM admininfo WHERE username=? AND password=?", usname,pass)
		for selDb.Next() {
						var ettype string
						var usnamee string
						var passs string
						var phone string
						var email string
						var fname string
						err = selDb.Scan(&usnamee, &passs, &email, &fname, &phone, &ettype) 
		if err == nil && etype==ettype{
			log.Println(usname + " " + pass + " " )
			log.Println(usnamee+" "+passs+" "+ettype)
			if etype == "admin" {
				http.Redirect(w, r, "/adminindex", 301)
			} else if etype == "student" {
				http.Redirect(w, r, "/studentindex", 301)
			} else if etype == "teacher" {
				http.Redirect(w, r, "/teacherindex", 301)
			}
		} else {
			http.Redirect(w,r,"/",301)
		}
	}
	}

	// if r.Method == "POST" {
	// 	uname := r.FormValue("username")
	// 	pass := r.FormValue("password")
	// 	// empl := r.FormValue("type")
	// 	selDb, err := db.Query("SELECT * FROM admininfo WHERE username=?", uname)
	// 	if err != nil {
	// 		log.Println("ERROR")
	// 		log.Println(pass)
	// 	} else {
	// 		log.Println("NO ERRORS")
	// 		for selDb.Next() {
	// 			var etype string
	// 			var usname string
	// 			var passs string
	// 			var phone string
	// 			var email string
	// 			var fname string
	// 			err = selDb.Scan(&usname, &passs, &email, &fname, &phone, &etype)
	// 			if err != nil {
	// 				panic(err.Error())
	// 			} else{
	// 				log.Println(usname+" "+passs)
	// 			}
	// 			if etype == "admin" {
	// 				http.Redirect(w, r, "/adminindex", 301)
	// 			} else if etype == "student" {
	// 				http.Redirect(w, r, "/studentindex", 301)
	// 			} else if etype == "teacher" {
	// 				http.Redirect(w, r, "/teacherindex", 301)
	// 			}
	// 	}
	// }

	// }
	defer db.Close()

}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		fname := r.FormValue("fname")
		pass := r.FormValue("pass")
		email := r.FormValue("email")
		uname := r.FormValue("uname")
		phone := r.FormValue("phone")
		empl := r.FormValue("type")
		insForm, err := db.Prepare("INSERT into admininfo(username,password,email,fname,phone,type) VALUES(?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(uname, pass, email, fname, phone, empl)
		// tmpl2.ExecuteTemplate(w, "success",nil)
		// http.Redirect(w,r,"/",301)
		http.Redirect(w, r, "/", 301)

	}
	defer db.Close()

}



// func Update(w http.ResponseWriter, r *http.Request) {
//     db := dbConn()
//     if r.Method == "POST" {
//         name := r.FormValue("name")
//         city := r.FormValue("city")
//         id := r.FormValue("uid")
//         insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
//         if err != nil {
//             panic(err.Error())
//         }
//         insForm.Exec(name, city, id)
//         log.Println("UPDATE: Name: " + name + " | City: " + city)
//     }
//     defer db.Close()
//     http.Redirect(w, r, "/", 301)
// }

// func Delete(w http.ResponseWriter, r *http.Request) {
//     db := dbConn()
//     emp := r.URL.Query().Get("id")
//     delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
//     if err != nil {
//         panic(err.Error())
//     }
//     delForm.Exec(emp)
//     log.Println("DELETE")
//     defer db.Close()
//     http.Redirect(w, r, "/", 301)
// }

func main() {
	log.Println("Server started on: http://localhost:8080")
	mux := http.NewServeMux()
	fz := http.FileServer(http.Dir("assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fz))

	mux.HandleFunc("/", Index)
	mux.HandleFunc("/signup", Signup)
	mux.HandleFunc("/reset", Reset)
	mux.HandleFunc("/insert", Insert)
	mux.HandleFunc("/adminindex", AdminIndex)
	mux.HandleFunc("/studentindex", StudentIndex)
	mux.HandleFunc("/teacherindex", TeacherIndex)
	mux.HandleFunc("/login", Login)

	//http.HandleFunc("/admin/signup",Insert)
	// http.HandleFunc("/show", Show)
	// http.HandleFunc("/new", New)
	// http.HandleFunc("/edit", Edit)
	// http.HandleFunc("/insert", Insert)
	// http.HandleFunc("/update", Update)
	// http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", mux)
}