package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id       int64
	Username string
	Kind     string
	Password string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/seckill?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("main.init err:", err)
	}
	defer db.Close()
	//*********************************在这里进行练习*********************************//
	var users = []User{}
	res := db.Exec("update users set password = ? where id = ?", "123456", 1)
	//*********************************在这里进行练习*********************************//
	if res.Error != nil {
		log.Fatal("create fail err:", res.Error)
	}
	fmt.Println(res.RowsAffected, res.Value)
	fmt.Println(users)
}
