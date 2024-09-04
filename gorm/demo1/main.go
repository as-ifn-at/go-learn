package main

import (
	"fmt"

	postgres "github.com/yugabyte/gorm-yugabytedb"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "yugabyte"
	password = "yugabyte"
	dbname   = "gorm_test"
)

type Employee struct {
	Id       int64  `gorm:"primary_key"`
	Name     string `gorm:"size:255"`
	Age      int64
	Language string `gorm:"size:255"`
}

func main() {

	baseUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		user, password, host, port, dbname)
	url := fmt.Sprintf("%s?load_balance=true&yb_servers_refresh_interval=240", baseUrl)

	// db, err := gorm.Open(postgres.Open(url), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	fmt.Printf("url: %v\n", url)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(&Employee{})

	db.Create(&Employee{Id: 12, Name: "John", Age: 35, Language: "Golang-GORM"})
	db.Create(&Employee{Id: 22, Name: "Smith", Age: 24, Language: "Golang-GORM"})
	db.Create(&Employee{Id: 32, Name: "Asif", Age: 24, Language: "GolaGORM"})

}
