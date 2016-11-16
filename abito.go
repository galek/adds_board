// db
package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var rows *sql.Rows
	rows, err = stmtInsCateg.Query()

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
	connectionToDB()
	http.HandleFunc("/", viewHandler)
	defer stmtInsCateg.Close() // Close the statement when we leave main() / the program terminates
	defer db.Close()

	http.ListenAndServe(":8080", nil)
}

type Page struct {
	Title string
	Body  []byte
}

var db *sql.DB
var stmtInsCateg *sql.Stmt
var err error

func printError() {
	if err != nil {
		println("Error: with DB ", err.Error())
	}
}

/**/
func connectionToDB() {
	db, err = sql.Open("mysql", "root:@/_abito")

	//var result sql.Result
	err = db.Ping()

	printError()

	// Prepare statement for inserting data
	stmtInsCateg, err = db.Prepare("SELECT name FROM categories")
	printError()

	// Categories list from DB
	//?место куда я вписываю свои данные

}
