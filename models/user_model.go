package models

import (
	"blogweb_gin/database"
	"fmt"
)

type User struct {
	Id int
	Username string
	Password string
	Status int // 0: 正常状态 1: 已删除
	CreateTime int64
}

func InsertUser(user User) (int64, error){
	return database.ModifyDB(
		"insert into users(username, password, status, createTime) values (?,?,?,?);",
		user.Username,
		user.Password,
		user.Status,
		user.CreateTime,
		)
}

func QueryUserWithCondition(con string) (int, error) {
	sqlStatement := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sqlStatement)
	row := database.QueryRowDB(sqlStatement)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func QueryWithUsername(username string) (int, error) {
	statement := fmt.Sprintf("where username='%s';", username)
	return QueryUserWithCondition(statement)
}