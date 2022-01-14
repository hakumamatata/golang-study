package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"php16_html_template/practiceA/funcs"
)

func main() {
	http.HandleFunc("/guestbook", funcs.ViewHandler)
	http.HandleFunc("/guestbook/new", funcs.NewHandler)
	http.HandleFunc("/guestbook/create", funcs.CreateHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
