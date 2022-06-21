package main

import (
	"github.com/Shaviaditya/BasicGo/go_fibre/database"
	"github.com/gofiber/fiber/v2"
)

func main(){
	app:= fiber.New()
	database.ConnectDB()
	app.Get("/",func(c *fiber.Ctx) error {
		err:= c.SendString("API is up & running")
		return err
	})

	app.Listen(":3500")
}