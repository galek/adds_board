package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	//_ "github.com/mattn/go-sqlite3"

	_ "github.com/go-sql-driver/mysql"
)

//========================================
func Header(w http.ResponseWriter) {
	fmt.Fprint(w, "<html>")
	fmt.Fprint(w, "<body>")
}

func Footer(w http.ResponseWriter) {
	fmt.Fprint(w, "</body>")
	fmt.Fprint(w, "</html>")
}

//========================================
func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	Header(w)
	CategoriesShow(w)
	println("CategoriesHandler: with DB ", r.FormValue("id"))
	Footer(w)
}

func ListOfAddsHandler(w http.ResponseWriter, r *http.Request) {

	println("ListOfAddsHandler Body: with DB ", r.FormValue("id"))

	i, err := strconv.ParseInt(r.FormValue("id")[0:], 10, 32)
	if err != nil {
		println("Invalid error id ", i)
		return
	}
	id := int(i)

	ListOfAddsQuery(id)

	Header(w)
	ListOfAddsShow(w)
	Footer(w)
}

func MessageShowHandler(w http.ResponseWriter, r *http.Request) {
	println("MessageShowHandler Body: with DB ", r.FormValue("id"),
		r.FormValue("caption"), r.FormValue("content"), r.FormValue("phonenumber"), r.FormValue("created"))

	Header(w)
	GetMessageBody(w)
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

	http.HandleFunc("/", CategoriesHandler)
	http.HandleFunc("/adds", ListOfAddsHandler)
	http.HandleFunc("/showmessage", MessageShowHandler)

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
var stmtCateg *sql.Stmt //List of categories
var stntAdds *sql.Stmt  // list of all adds by categoryID
//var stntMessageBody *sql.Stmt // list of all adds by categoryID
var err error

func printError() {
	if err != nil {
		println("Error: with DB ", err.Error())
	}
}

/*
Получает список всех объявлений, по выбранной категории
*/
func ListOfAddsQuery(selectedCategoryID int) {
	var req string = "SELECT id,caption FROM postings WHERE categoryID=" + strconv.Itoa(selectedCategoryID)
	stntAdds, err = db.Prepare(req)
	printError()
}

type Message struct {
	id          int
	cookie      string
	caption     string
	content     string
	phonenumber string
	created     int
}

/*Получает тело */
func GetMessageBody(w http.ResponseWriter) {
	// TODO: Replace
	var req string = "SELECT * FROM postings WHERE id=" + strconv.Itoa(1)
	var stntMessageBody *sql.Stmt // list of all adds by categoryID
	stntMessageBody, err = db.Prepare(req)
	printError()

	//Читаем все значения
	var rows *sql.Rows
	rows, err = stntMessageBody.Query()

	printError()

	var value string
	for rows.Next() {
		rows.Scan(&value)
		fmt.Fprintf(w, "<a href=''>GetMessageBody %s</a>\n", value, value)
	}

	printError()
	defer rows.Close()
	defer stntMessageBody.Close()
}

//========================================
func ListOfAddsShow(w http.ResponseWriter) {
	var rows *sql.Rows
	rows, err = stntAdds.Query()

	printError()
	defer rows.Close()

	var value string
	var id int
	for rows.Next() {
		rows.Scan(&id, &value)
		fmt.Fprintf(w, "<p><a href='/showmessage?id=%d'>[DEBUG ONLY ListOfAddsShow]%s</a>\n</p>", id, value)
	}

	printError()
}

//========================================
func CategoriesQuery() {
	// Prepare statement for inserting data
	stmtCateg, err = db.Prepare("SELECT id, name FROM categories")
	printError()
}

func CategoriesShow(w http.ResponseWriter) {
	var rows *sql.Rows
	rows, err = stmtCateg.Query()

	printError()
	defer rows.Close()

	var value string
	var id int
	for rows.Next() {
		rows.Scan(&id, &value)
		fmt.Fprintf(w, "<p><a href='/adds?id=%d'>[DEBUG ONLY CategoriesShow]%s</a>\n</p>", id, value)
	}

	printError()
}

//========================================
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
}
