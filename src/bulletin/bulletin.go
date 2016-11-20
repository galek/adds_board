package main

import (
	"database/sql"
	"fmt"
	"net/http"

	//_ "github.com/mattn/go-sqlite3"

	_ "github.com/go-sql-driver/mysql"
)

func Header(w http.ResponseWriter) {
	fmt.Fprint(w, "<html>")
	fmt.Fprint(w, "<body>")
}

func Footer(w http.ResponseWriter) {
	fmt.Fprint(w, "</body>")
	fmt.Fprint(w, "</html>")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	Header(w)
	CategoriesShow(w)
	ListOfAddsShow(w)
	Footer(w)
}

func main() {

	//{
	//	if _, err = os.Stat("./bulletin.db"); os.IsNotExist(err) {
	//		println("database ./bulletin.db doesn't exist")
	//		return
	//	}
	//}

	connectionToDB()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", viewHandler)
	defer stmtCateg.Close() // Close the statement when we leave main() / the program terminates
	defer stntAdds.Close()
	defer db.Close()

	http.ListenAndServe(":8080", nil)
}

type Page struct {
	Title string
	Body  []byte
}

var db *sql.DB
var stmtCateg *sql.Stmt     //List of categories
var stntAdds *sql.Stmt      // list of all adds by categoryID
var stnAddsCatIds *sql.Stmt // list of all categoriesIDs
var err error

func printError() {
	if err != nil {
		println("Error: with DB ", err.Error())
	}
}

/**/
func ListOfAddsQuery() {
	// Prepare statement for inserting data
	stntAdds, err = db.Prepare("SELECT caption FROM postings WHERE categoryID=1")
	printError()
}

func ListOfAddsShow(w http.ResponseWriter) {
	var rows *sql.Rows
	rows, err = stntAdds.Query()

	printError()
	defer rows.Close()

	var value string
	for rows.Next() {
		rows.Scan(&value)
		fmt.Fprintf(w, "<a href='http://google.com'>%s</a>\n", value)
	}

	printError()
}

func CategoriesQuery() {
	// Prepare statement for inserting data
	stmtCateg, err = db.Prepare("SELECT name FROM categories")
	printError()
}

func CategoriesShow(w http.ResponseWriter) {
	var rows *sql.Rows
	rows, err = stmtCateg.Query()

	printError()
	defer rows.Close()

	var value string
	for rows.Next() {
		rows.Scan(&value)
		fmt.Fprintf(w, "<a href='http://google.com'>%s</a>\n", value)
	}

	printError()
}

/**/
func connectionToDB() {
	//	db, err = sql.Open("sqlite3", "./bulletin.db")
	db, err = sql.Open("mysql", "root:@/_abito")

	if err != nil {
		println(err.Error())
	}
	err = db.Ping()

	printError()

	CategoriesQuery()
	ListOfAddsQuery()
}
