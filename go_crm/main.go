package main

import (
	"fmt"
	"log"
	"github.com/Shaviaditya/BasicGo/go_crm/lead"
	"github.com/Shaviaditya/BasicGo/go_crm/database"
	"github.com/gofiber/fiber/v2"
);

func setupRoutes(app *fiber.App){
	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Hello CRM");
	})
	app.Get("/api/lead/:id",lead.GetLead)
	app.Get("/api/lead",lead.GetLeads)
	app.Post("/api/lead/",lead.NewLead)
	app.Delete("/api/lead/:id",lead.DeleteLead)
}


func main(){
	app := fiber.New()
	setupRoutes(app)
	database.InitDatabase()
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
	log.Fatal(app.Listen(":3000"))
	defer database.DBConn.Close()
}