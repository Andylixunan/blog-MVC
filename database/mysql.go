package database

import (
	"blogweb_gin/utils"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql() {
	utils.Logger.Println("InitMysql....")
	if db == nil {
		var err error
		db, err = sql.Open("mysql", "root:Andy990515!@tcp(127.0.0.1:3306)/micro_blog_gin?charset=utf8")
		if err != nil {
			utils.Logger.Fatal(err)
		}
		err = db.Ping()
		if err != nil {
			utils.Logger.Fatal(err)
		}
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(5 * time.Minute)
		CreateTableWithUser()
		CreateTableWithArticle()
	}
}

func CreateTableWithUser() {
	sqlString := `CREATE TABLE IF NOT EXISTS users(
        id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64),
        status INT,
        createTime BIGINT
        );`
	ModifyDB(sqlString)
}

func CreateTableWithArticle() {
	sql := `create table if not exists article(
        id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
        title varchar(30),
        author varchar(20),
        tags varchar(30),
        short varchar(255),
        content longtext,
        createTime BIGINT
        );`
	ModifyDB(sql)
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		utils.Logger.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		utils.Logger.Println(err)
		return 0, err
	}
	return count, nil
}

func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
