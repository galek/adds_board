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
	MakeCookiesGreatAgain(w, r)
	Header(w)
	CategoriesShow(w)
	println("CategoriesHandler: with DB ", r.FormValue("id"))
	Footer(w)
}

func ListOfAddsHandler(w http.ResponseWriter, r *http.Request) {

	println("ListOfAddsHandler Body: with DB ", r.FormValue("id"))
	MakeCookiesGreatAgain(w, r)

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
	MakeCookiesGreatAgain(w, r)

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
	http.HandleFunc("/newmessage", NewMessageHandler)
	http.HandleFunc("/create_message", NewMessageHandlerUtil)
	http.HandleFunc("/deletemessage", DeleteMessageHandler)

	CookiesInit()

	defer stmtCateg.Close() // Close the statement when we leave main() / the program terminates
	defer stntAdds.Close()
	defer DB.Close()

	http.ListenAndServe(":8080", nil)
}

var DB *sql.DB
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
	stntAdds, err = DB.Prepare(req)
	printError()
}

/*Получает тело */
func GetMessageBody(w http.ResponseWriter, id int) {
	var req string = "SELECT caption, content, phonenumber, created FROM postings WHERE id=" + strconv.Itoa(id)
	var stntMessageBody *sql.Stmt // list of all adds by categoryID
	stntMessageBody, err = DB.Prepare(req)
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
	stmtCateg, err = DB.Prepare("SELECT id, name FROM categories")
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

	fmt.Fprintf(w, "<p align=\"right\"><a href='/mymessages'>My Messages</a>\n</p>")

	printError()
}

func ShowErrorPage(w http.ResponseWriter) {
	Header(w)
	fmt.Fprintf(w, "ERROR - 404")
	Footer(w)
}

//========================================
// СОЗДАНИЕ СООБЩЕНИЙ
//========================================
func NewMessageHandler(w http.ResponseWriter, r *http.Request) {
	Header(w)

	i, err := strconv.ParseInt(r.FormValue("categoryID")[0:], 10, 32)
	if err != nil {
		println("Invalid error categoryID ", i)
		ShowErrorPage(w)
		return
	}
	categoryID := int(i)

	fmt.Fprintf(w, "<p>")
	fmt.Fprintf(w, "<form id=\"create_message_form\" action=\"create_message\" method=\"get\">Заголовок")
	fmt.Fprintf(w, "</p>")

	fmt.Fprintf(w, "<p>")
	fmt.Fprintf(w, "<input type=\"text\" name=\"caption\">Текст объявления")
	fmt.Fprintf(w, "</p>")

	fmt.Fprintf(w, "<p>")
	fmt.Fprintf(w, "<input type=\"text\" name=\"telephone\">Номер телефона")
	fmt.Fprintf(w, "</p>")

	fmt.Fprintf(w, "<p>")
	fmt.Fprintf(w, "<input type=\"text\" name=\"categoryID\" value=\"%d\" style=\"display:none;\">", categoryID)
	fmt.Fprintf(w, "</p>")

	fmt.Fprintf(w, "<p>")
	fmt.Fprintf(w, "<textarea name=\"body\"></textarea><input type=\"submit\" value=\"Создать объявление\"></form>")
	fmt.Fprintf(w, "</p>")

	Footer(w)
}

func NewMessageHandlerUtil(w http.ResponseWriter, r *http.Request) {

	i, err := strconv.ParseInt(r.FormValue("categoryID")[0:], 10, 32)
	if err != nil {
		println("Invalid error categoryID ", i)
		ShowErrorPage(w)
		return
	}
	categoryID := int(i)

	caption := r.FormValue("caption")
	body := r.FormValue("body")
	telephone := r.FormValue("telephone")

	_CreateEmptyMessage(categoryID, caption, body, telephone)
}

//========================================
// УДАЛЕНИЕ СООБЩЕНИЙ
//========================================
func DeleteMessageReq(w http.ResponseWriter, id string) {
	// TODO: 10 for tests
	var req string = "DELETE FROM postings WHERE cookieid=? AND id = ?"
	var stntMessageBody *sql.Stmt // list of all adds by categoryID
	stntMessageBody, err = DB.Prepare(req)
	printError()

	println("CookieId: %d %d", CookieId, id)

	//Читаем все значения
	var rows *sql.Rows
	rows, err = stntMessageBody.Query(CookieId, id)

	printError()

	defer rows.Close()
	defer stntMessageBody.Close()
}

func DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	println("DeleteMessageHandler Body: with DB ", r.FormValue("cookie"))
	println("DeleteMessageHandler Body: with DB ", r.FormValue("id"))

	MakeCookiesGreatAgain(w, r)

	idStr := r.FormValue("id")

	Header(w)
	DeleteMessageReq(w, idStr)
	MyMessagesShow(w)
	Footer(w)
}

//========================================
// СОЗДАНИЕ НОВОГО СООБЩЕНИЯ  - НЕ ЗАКОНЧЕНО
// 99999-DATE

func _CreateEmptyMessage(categoryID int, caption string, body string, telephone string) {
	println("[CreateEmptyMessage]")

	//var req string = "INSERT INTO `postings` VALUES ('" + id + "'," + "''," + "'" + cookie + "'," + "'', '', '', '9999')"

	var req string = "INSERT INTO postings (categoryId, cookieid, caption, content, phonenumber, created) VALUES(?,?,?,?,?,9999)"

	var stntMessageBody *sql.Stmt // list of all adds by categoryID
	stntMessageBody, err = DB.Prepare(req)
	printError()

	//Читаем все значения
	var rows *sql.Rows
	rows, err = stntMessageBody.Query(categoryID, CookieId, caption, body, telephone)

	printError()
	defer rows.Close()
	defer stntMessageBody.Close()
}

//========================================
// СПИСОК МОИХ СООБЩЕНИЙ - НЕ ЗАКОНЧЕНО
func MyMessagesShow(w http.ResponseWriter) {
	var req string = "SELECT id, caption, content, phonenumber, created FROM postings WHERE cookieid=?" // + string(CookieId)
	var stntMessageBody *sql.Stmt                                                                       // list of all adds by categoryID
	stntMessageBody, err = DB.Prepare(req)
	printError()

	//Читаем все значения
	var rows *sql.Rows
	rows, err = stntMessageBody.Query(CookieId)

	printError()

	var caption string
	var id int
	var content string
	var phonenumber string
	var created int
	for rows.Next() {
		rows.Scan(&id, &caption, &content, &phonenumber, &created)
		fmt.Fprintf(w, "[DEBUG ONLY-MyMessagesShow] <p>%s</p> %s %s <p>%d</p>\n", caption, content, phonenumber, created)

		fmt.Fprintf(w, "<button onclick=\"location.href='/deletemessage?id=%d'\">Delete</button>", id)
	}

	printError()
	defer rows.Close()
	defer stntMessageBody.Close()
}

// TODO: Передавать куку запросом
func MyMessagesHandler(w http.ResponseWriter, r *http.Request) {
	println("MyMessagesHandler Body: with DB ", r.FormValue("cookie"))
	MakeCookiesGreatAgain(w, r)

	Header(w)
	MyMessagesShow(w)
	Footer(w)
}

//========================================
/**/
func connectionToDB() {
	//	DB, err = sql.Open("sqlite3", "./bulletin.db")
	DB, err = sql.Open("mysql", "root:@/_abito")

	if err != nil {
		println(err.Error())
	}
	err = DB.Ping()

	printError()

	CategoriesQuery()
}
