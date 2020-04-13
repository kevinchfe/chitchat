package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/chitchat?parseTime=true")
	if err != nil {
		log.Fatalln(err.Error())
	}

	Db.SetMaxIdleConns(10)
	Db.SetMaxOpenConns(100)

	if err := Db.Ping(); err != nil {
		log.Fatalln(err.Error())
	}

}
