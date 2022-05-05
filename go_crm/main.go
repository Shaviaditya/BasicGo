package main

import (
	"fmt"
	"github.com/Shaviaditya/BasicGo/go_crm/lead"
	"github.com/Shaviaditya/BasicGo/go_crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func setupRoutes(app *fiber.App){
	app.Get("/lead",lead.GetLeads)
	app.Get("/lead/:id",lead.GetLead)
	app.Post("/lead",lead.NewLead)
	app.Delete("/lead/:id",lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("mysql","rudraditya:MySQL#311@/leadsDb?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		panic("Failed to Connect\n")
	}
	fmt.Println("Connection opened to Database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated\n")
}

func main(){
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}