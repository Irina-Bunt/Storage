package DatabaseStorage

import (
	"dataset/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var dbase *gorm.DB

func Init() (*gorm.DB, error) {
	dsn := "host=localhost user=admin password=admin dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	db.AutoMigrate(&model.Tokens{}, &model.Task{}, &model.User{})

	return db, nil
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase, _ = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Printf("Database is unavailable, wait for %d sec\n", sleep)
			time.Sleep(sleep * time.Second)
			dbase, _ = Init()
		}
	}
	return dbase
}
