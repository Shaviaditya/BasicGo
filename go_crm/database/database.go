package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Shaviaditya/BasicGo/go_crm/config"
);
var DBConn *gorm.DB

func InitDatabase(){
	var err error;
	conn := fmt.Sprintf("%s:%s#%s@/%s?charset=utf8&parseTime=True&loc=Local",config.Config("DB_USER"),config.Config("DB_PASSWORD"),config.Config("DB_PASS"),config.Config("DB_NAME"))
	fmt.Println(conn)
	DBConn,err = gorm.Open(config.Config("DB_DIALECT"),(conn))
	if err != nil {
		panic("Failed to Connect")
	}
	fmt.Println("Connection Opened")
}