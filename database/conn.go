package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)


func ConnetionMysql() *sql.DB {
    config := mysql.Config{
        User: "root",
        Passwd: "123456",
        Net:  "tcp",
        Addr: "localhost:3306",
        DBName: "gymapp2",
    }
    db, err := sql.Open("mysql", config.FormatDSN())
    if err != nil {
        panic(err)
    }
    if err := db.Ping(); err != nil {
        panic(err)
    }
    return db
}
