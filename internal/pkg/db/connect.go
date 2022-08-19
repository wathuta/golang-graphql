package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect()  {
	db,err:=sql.Open("mysql","root:root@(localhost:3306)/graphql")
	if err!=nil{
		log.Fatal(err)
	}
	if err:=db.Ping();err!=nil{
		log.Fatal(err)
	}
	Db = db
}
func CloseDb()error{
	return Db.Close()
}