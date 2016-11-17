package main

import (
	"database/sql"
	"fmt"
	"net/http"
  "os"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var rows *sql.Rows
	rows, err = stmtCateg.Query()

	printError()
	defer rows.Close()

	var value string
	for rows.Next() {
		rows.Scan(&value)
		fmt.Fprintf(w, "%s", value)
	}

	printError()
}

func main() {
  if _, err = os.Stat("./bulletin.db");os.IsNotExist(err) {
      println("database ./bulletin.db doesn't exist" )
    return
  }
	connectionToDB()
  http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", viewHandler)
	defer stmtCateg.Close() // Close the statement when we leave main() / the program terminates
	defer db.Close()

	http.ListenAndServe(":8080", nil)
}

type Page struct {
	Title string
	Body  []byte
}

var db *sql.DB
var stmtCateg *sql.Stmt
var err error

func printError() {
	if err != nil {
		println("Error: with DB ", err.Error())
	}
}

/**/
func connectionToDB() {
	//db, err = sql.Open("mysql", "root:@/_abito")
  db, err = sql.Open("sqlite3","./bulletin.db")
  if err != nil {
    println(err.Error())
  }
	err = db.Ping()

	printError()

	// Prepare statement for inserting data
	stmtCateg, err = db.Prepare("SELECT name FROM categories")
	printError()
}
