package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"mooblog/utils"
	"os"
	"time"
)

var db *gorm.DB //指定数据原型
var err error   //接受错误

func InitDb() { //数据库模型
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))

	if err != nil {
		fmt.Println("连接出错", err)
		os.Exit(1)
	}

	db.SingularTable(true) //禁用数据表复数形式user->users

	db.AutoMigrate(&User{}, &Article{}, &Category{})
	//{
	//	//fmt.Println(ok,er)
	//	//if err != nil {
	//	//	fmt.Println("连接出错", er)
	//	//
	//	//}
	//}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)
	//db.Close()
}
