package databases

import (
	config2 "chitchat/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	config := config2.LoadConfig()
	driver := config.Db.Driver
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", config.Db.User, config.Db.Password, config.Db.Address, config.Db.Database)
	//Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/chitchat?parseTime=true")
	Db, err = sql.Open(driver, source)
	if err != nil {
		log.Fatalln(err.Error())
	}

	Db.SetMaxIdleConns(10)
	Db.SetMaxOpenConns(100)

	if err := Db.Ping(); err != nil {
		log.Fatalln(err.Error())
	}

}
