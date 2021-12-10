package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql....")
	if db == nil {
		db, _ = sql.Open("mysql", "andy:Andy990515@tcp(127.0.0.1:3306)/blog_gin?charset=utf8")
		CreateTableWithUser()
	}
}

func CreateTableWithUser(){
	sqlString := `CREATE TABLE IF NOT EXISTS users(
        id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64),
        status INT,
        createTime BIGINT
        );`
	ModifyDB(sqlString)
}

func ModifyDB(sql string, args ...interface{}) (int64, error){
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil{
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}