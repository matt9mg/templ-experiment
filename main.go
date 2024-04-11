package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/matt9mg/templ-experiment/controllers"
	"log"
)

func main() {
	app := fiber.New()

	homeController := controllers.NewHome()

	app.Static("/assets", "./assets/src/dist")
	app.Get("/", homeController.Index)
	app.Post("/test", homeController.Save)

	log.Fatalln(app.Listen(":8000"))
}
