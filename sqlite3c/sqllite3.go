package sqlite3c

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func StartSqlite() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	if db.Ping() != nil {
		panic(db.Ping())
	}

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
	fmt.Println(stmt)
	log.Println(db)
}
