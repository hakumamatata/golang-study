package funcs

import "database/sql"

const DbSourceName = "root:1234@/local?charset=utf8mb4"

func GetDb() (*sql.DB, error) {
	return sql.Open("mysql", DbSourceName)
}
