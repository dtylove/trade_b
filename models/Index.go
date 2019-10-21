package models

import (
	"dtyTrade/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var gdb *gorm.DB
var once sync.Once

func InitOnce() {
	dataDriver := config.Conf.DataSource
	url := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dataDriver.User,
		dataDriver.Password,
		dataDriver.Addr,
		dataDriver.Database,
	)
	gdb, err := gorm.Open(dataDriver.DriverName, url)

	if err != nil {
		panic(err)
	}

	gdb.SingularTable(true)       //全局设置表名不可以为复数形式。
	gdb.DB().SetMaxIdleConns(10)  //SetMaxOpenConns用于设置最大打开的连接数
	gdb.DB().SetMaxOpenConns(100) //SetMaxIdleConns用于设置闲置的连接数

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "g_" + defaultTableName
	}

	if dataDriver.Migrate {
		gdb.CreateTable(&User{})
	}
}

func InitDB() {
	once.Do(InitOnce)
}

func GetDB() *gorm.DB {
	return gdb
}
