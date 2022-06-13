package main

import (
	"fmt"
	"gorm-mysql/setting"
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

var DBSetting *setting.DBsetting

func init() {
	setting, err := setting.NewSetting()
	if err != nil {
		log.Fatal("--init setting err:--", err)
	}
	setting.ReadSection("DataBase", &DBSetting)
}

func main() {
	//db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/seckill?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", DBSetting.DataSource)

	if err != nil {
		log.Fatal("main.init err:", err)
	}
	defer db.Close()
	//*********************************在这里进行练习*********************************//
	var user []User
	res := db.Select("username").Find(&user)
	//*********************************在这里进行练习*********************************//
	if res.Error != nil {
		log.Fatal("create fail err:", res.Error)
	}
	fmt.Println(res.RowsAffected, res.Value)
	fmt.Println(user)
}
