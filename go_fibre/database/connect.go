package database

import (
	"fmt"
	"log"
	"strconv"
	"github.com/Shaviaditya/BasicGo/go_fibre/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB;

func ConnectDB(){
	var err error
	p:= config.Config("DB_PORT")
	port, err := strconv.ParseUint(p,10,32)

	if err!=nil {
		log.Println("Error while parsing")
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB,err := gorm.Open(postgres.Open(dsn))

	if err != nil {
        panic("failed to connect database")
    }
	fmt.Print(DB)
    fmt.Println("Connection Opened to Database")
}
