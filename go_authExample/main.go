package main

import (
	"log"
	"github.com/Shaviaditya/BasicGo/go_authExample/controller"
	"github.com/gofiber/fiber/v2"
);
func controllers(app *fiber.App){
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./statics/index.html",true)
	})
	app.Get("/login", controller.Login)
	app.Get("/google/callback",controller.Callback)
}
func main(){
	app := fiber.New()
	controllers(app)
	app.Static("/","./statics")
	log.Fatal(app.Listen(":5500"))
}