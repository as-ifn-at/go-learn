package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	
	
	
	conn := fmt.Sprintf("host= %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err := gorm.Open("postgres", conn)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(&Employee{})

	db.Create(&Employee{Id: 1, Name: "John", Age: 35, Language: "Golang-GORM"})
	db.Create(&Employee{Id: 2, Name: "Smith", Age: 24, Language: "Golang-GORM"})
	db.Create(&Employee{Id: 3, Name: "Asif", Age: 24, Language: "GolaGORM"})

	// Display input data
	var employees []Employee
	db.Find(&employees)
	for _, employee := range employees {
		fmt.Printf("Employee ID:%d\nName:%s\nAge:%d\nLanguage:%s\n", employee.Id, employee.Name, employee.Age, employee.Language)
		fmt.Printf("--------------------------------------------------------------\n")
	}
}
