package funcs

import (
	"html/template"
	"log"
	"net/http"
	"php16_html_template/practiceA/types"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ViewHandler(writer http.ResponseWriter, request *http.Request) {
	// 使用直接輸出
	//msg := []byte("signature list goes here:")
	//_, err := writer.Write(msg)
	//check(err)

	// 讀取資料庫資料顯示
	db, err := GetDb()
	//查詢資料
	rows, err := db.Query("SELECT * FROM guestbook_sign")
	check(err)
	var guestbooks types.Guestbook
	var guestNames []string
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		check(err)
		guestNames = append(guestNames, name)
	}
	guestbooks = types.Guestbook{
		SCount:     len(guestNames),
		Signatures: guestNames,
	}

	// 使用模板
	// 路徑要為go path
	html, err := template.ParseFiles("practice/src/php16_html_template/practiceA/views/view.html")
	check(err)
	err = html.Execute(writer, guestbooks)
	check(err)
}

func NewHandler(writer http.ResponseWriter, request *http.Request) {
	// 使用模板
	// 路徑要為go path
	html, err := template.ParseFiles("practice/src/php16_html_template/practiceA/views/new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func CreateHandler(writer http.ResponseWriter, request *http.Request) {
	// 取得request資料
	signature := request.FormValue("signature")

	// 寫入資料庫
	db, err := GetDb()
	check(err)

	stmt, err := db.Prepare("INSERT guestbook_sign SET name=?")
	check(err)

	_, err = stmt.Exec(signature)
	check(err)

	//id, err := res.LastInsertId()
	//check(err)

	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}
