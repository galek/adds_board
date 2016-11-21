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
		ShowErrorPage(w)
		return
	}
	id := int(i)

	ListOfAddsQuery(id)

	Header(w)
	ListOfAddsShow(w)
	Footer(w)
}

func MessageShowHandler(w http.ResponseWriter, r *http.Request) {
	println("MessageShowHandler Body: with DB ", r.FormValue("id"))

	i, err := strconv.ParseInt(r.FormValue("id")[0:], 10, 32)
	if err != nil {
		println("Invalid error id ", i)
		ShowErrorPage(w)
		return
	}
	id := int(i)

	Header(w)
	GetMessageBody(w, id)
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
	http.HandleFunc("/mymessages", MyMessagesHandler)
	http.HandleFunc("/deletemessage", DeleteMessageHandler)

	defer stmtCateg.Close() // Close the statement when we leave main() / the program terminates
	defer stntAdds.Close()
	defer db.Close()

	http.ListenAndServe(":8080", nil)
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

/*Получает тело */
func GetMessageBody(w http.ResponseWriter, id int) {
	var req string = "SELECT caption, content, phonenumber, created FROM postings WHERE id=" + strconv.Itoa(id)
	var stntMessageBody *sql.Stmt // list of all adds by categoryID
	stntMessageBody, err = db.Prepare(req)
	printError()

	//Читаем все значения
	var rows *sql.Rows
	rows, err = stntMessageBody.Query()

	printError()

	var caption string
	var content string
	var phonenumber string
	var created int
	for rows.Next() {
		rows.Scan(&caption, &content, &phonenumber, &created)
		fmt.Fprintf(w, "[DEBUG ONLY-GetMessageBody] <p>%s</p> <p>%s</p> <p>%s</p> <p>%d</p>\n", caption, content, phonenumber, created)
	}

	printError()

	if caption == "" && content == "" && phonenumber == "" && created == 0 {
		ShowErrorPage(w)
	}

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

func ShowErrorPage(w http.ResponseWriter) {
	Header(w)
	fmt.Fprintf(w, "ERROR - 404")
	Footer(w)
}

//========================================
// УДАЛЕНИЕ СООБЩЕНИЙ - НЕ ЗАКОНЧЕНО
// тут надо определиться, будем удалять на 1 странице со списком, или на отдельной странице
//========================================
func DeleteMessageReq(w http.ResponseWriter, cookie string, id string) {
	// TODO: 10 for tests
	var req string = "DELETE FROM postings WHERE cookie='" + cookie + "'" + "AND id ='" + id + "'"
	var stntMessageBody *sql.Stmt // list of all adds by categoryID
	stntMessageBody, err = db.Prepare(req)
	printError()

	//Читаем все значения
	var rows *sql.Rows
	rows, err = stntMessageBody.Query()

	printError()

	defer rows.Close()
	defer stntMessageBody.Close()
}

// TODO: Передавать куку запросом
func DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	println("DeleteMessageHandler Body: with DB ", r.FormValue("cookie"))
	println("DeleteMessageHandler Body: with DB ", r.FormValue("id"))

	cookieStr := r.FormValue("cookie")
	if cookieStr == "" {
		println("[DEBUG DeleteMessageHandler]INVALID COOKIE")
		ShowErrorPage(w)
		return
	}
	idStr := r.FormValue("id")
	if cookieStr == "" {
		println("[DEBUG DeleteMessageHandler]INVALID id")
		ShowErrorPage(w)
		return
	}

	// TODO:TESTING OF CREATION NEW MESSAGE
	if cookieStr == "800" && idStr == "800" {
		CreateNewMessage(w, "sdwe", "800")
	}

	Header(w)
	DeleteMessageReq(w, cookieStr, idStr)
	Footer(w)
}

//========================================
// СОЗДАНИЕ НОВОГО СООБЩЕНИЯ  - НЕ ЗАКОНЧЕНО
// 99999-DATE
func CreateNewMessage(w http.ResponseWriter, cookie string, id string) {
	_CreateEmptyMessage(w, cookie, id)
	UpdateMessage(w, id, "98", cookie, "caption", "content", "phonenumber", "99999")
}

func _CreateEmptyMessage(w http.ResponseWriter, cookie string, id string) {
	println("[CreateEmptyMessage]")

	var req string = "INSERT INTO `postings` VALUES ('" + id + "'," + "''," + "'" + cookie + "'," + "'', '', '', '9999')"

	var stntMessageBody *sql.Stmt // list of all adds by categoryID
	stntMessageBody, err = db.Prepare(req)
	printError()

	//Читаем все значения
	var rows *sql.Rows
	rows, err = stntMessageBody.Query()

	printError()
	defer rows.Close()
	defer stntMessageBody.Close()
}
func UpdateMessage(w http.ResponseWriter, id string, categoryID string, cookie string, caption string, content string, phonenumber string, created string) {
	var req string = "UPDATE postings SET categoryID='" + categoryID + "'" + ", caption='" + caption + "'" + ", content='" + content + "'" + ", phonenumber='" + phonenumber + "'" + ", created='" + created + "'" + "WHERE id='" + id + "'" + "AND cookie='" + cookie + "'"

	var stntMessageBody *sql.Stmt // list of all adds by categoryID
	stntMessageBody, err = db.Prepare(req)
	printError()

	//Читаем все значения
	var rows *sql.Rows
	rows, err = stntMessageBody.Query()

	printError()
	defer rows.Close()
	defer stntMessageBody.Close()
}

//========================================
// СПИСОК МОИХ СООБЩЕНИЙ - НЕ ЗАКОНЧЕНО
func MyMessagesShow(w http.ResponseWriter, cookie string) {
	var req string = "SELECT caption, content, phonenumber, created FROM postings WHERE cookie=" + cookie
	var stntMessageBody *sql.Stmt // list of all adds by categoryID
	stntMessageBody, err = db.Prepare(req)
	printError()

	//Читаем все значения
	var rows *sql.Rows
	rows, err = stntMessageBody.Query()

	printError()

	var caption string
	var content string
	var phonenumber string
	var created int
	for rows.Next() {
		rows.Scan(&caption, &content, &phonenumber, &created)
		fmt.Fprintf(w, "[DEBUG ONLY-MyMessagesShow] <p>%s</p> %s %s <p>%d</p>\n", caption, content, phonenumber, created)
	}

	printError()
	defer rows.Close()
	defer stntMessageBody.Close()
}

// TODO: Передавать куку запросом

func MyMessagesHandler(w http.ResponseWriter, r *http.Request) {
	println("MyMessagesHandler Body: with DB ", r.FormValue("cookie"))

	cookieStr := r.FormValue("cookie")
	if cookieStr == "" {
		println("[DEBUG MyMessagesHandler]INVALID COOKIE")
		ShowErrorPage(w)
		return
	}

	Header(w)
	MyMessagesShow(w, cookieStr)
	Footer(w)
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
