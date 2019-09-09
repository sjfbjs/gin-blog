package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func Load() {
	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blade_blog")
	if nil != err {
		log.Fatalln(err)
	}
	// 设置最大打开的连接数，默认值为0表示不限制。
	db.SetMaxOpenConns(2000)
	// 设置闲置的连接数。
	db.SetMaxIdleConns(1000)
	DB = db
}