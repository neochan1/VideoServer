package database

import "database/sql"

var (
	DbConnection *sql.DB
	err error
)

func init()  {
	DbConnection,err = sql.Open("mysql","root:123!@#@tcp(localhost:3306)/video_server?charset=utf8")
}