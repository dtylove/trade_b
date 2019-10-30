package models

import (
	"dtyTrade/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

type Store struct {
	db *gorm.DB
}

var once sync.Once
var store Store

func InitOnce() {
	dataDriver := config.Conf.DataSource
	url := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dataDriver.User,
		dataDriver.Password,
		dataDriver.Addr,
		dataDriver.Database,
	)
	db, err := gorm.Open(dataDriver.DriverName, url)

	if err != nil {
		panic(err)
	}

	store.db = db
	store.db.SingularTable(true)       //全局设置表名不可以为复数形式。
	store.db.DB().SetMaxIdleConns(10)  //SetMaxOpenConns用于设置最大打开的连接数
	store.db.DB().SetMaxOpenConns(100) //SetMaxIdleConns用于设置闲置的连接数

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "g_" + defaultTableName
	}

	if dataDriver.Migrate {
		store.db.CreateTable(&User{}, &Order{}, &Product{})
	}
}

func InitDB() {
	once.Do(InitOnce)
}

func GetDB() *gorm.DB {
	return store.db
}

func (s *Store) BeginTx() (*Store, error) {
	db := s.db.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	return &Store{db: db}, nil
}

func (s *Store) Rollback() error {
	return s.db.Rollback().Error
}

func (s *Store) CommitTx() error {
	return s.db.Commit().Error
}
