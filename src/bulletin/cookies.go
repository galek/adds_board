package main

import (
	"database/sql"
	"math/rand"
	"net/http"
	"time"
)

var CookieId int64
var stmtCheckCookie *sql.Stmt
var stmtInsertCookie *sql.Stmt

func MakeCookiesGreatAgain(w http.ResponseWriter, r *http.Request) {
	var cookie *http.Cookie
	cookie, err = r.Cookie("id")
	if err != nil {
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie = &http.Cookie{
			Name:    "id",
			Value:   randStringRunes(16),
			Expires: expiration,
		}
	}
	row := stmtCheckCookie.QueryRow(cookie.Value)

	var CookieName string
	err = nil
	err = row.Scan(&CookieId, &CookieName)

	var result sql.Result

	if err != nil {
		result, err = stmtInsertCookie.Exec(cookie.Value)
		printError()
		if err == nil {
			CookieId, err = result.LastInsertId()
		} else {
			printError()
		}
	}
	http.SetCookie(w, cookie)
}

func CookiesInit() {
	rand.Seed(time.Now().UnixNano())
	stmtCheckCookie, err = DB.Prepare("SELECT id, cookie FROM cookies WHERE cookies.cookie = ?")

	printError()
	stmtInsertCookie, err = DB.Prepare("INSERT INTO cookies (cookie) VALUES (?)")

	printError()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
