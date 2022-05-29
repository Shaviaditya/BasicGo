package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/Shaviaditya/BasicGo/go_jwt/database"
	"github.com/Shaviaditya/BasicGo/go_jwt/routes"
	"github.com/Shaviaditya/BasicGo/go_jwt/models"
);


func controllers(app *fiber.App){
	app.Get("/",func(c *fiber.Ctx) error {
		return c.SendString("hello world!")
	})
	app.Post("/signup",routes.Signup);
	app.Post("/login",routes.Login);
	app.Get("/users",routes.AuthRequired(),routes.GetUsers);
	app.Get("/user/:id",routes.AuthRequired(),routes.GetUserId);
	app.Delete("/user/:id",routes.AuthRequired(),routes.DeleteUserId);
}

func main(){
	app := fiber.New();
	controllers(app);
	database.ConnectDB()
	database.DBConn.AutoMigrate(&models.Details{})
	log.Fatal(app.Listen(":5700")) 
	defer database.DBConn.Close()
}