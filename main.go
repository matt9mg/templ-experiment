package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/matt9mg/templ-experiment/controllers"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", controllers.NewHome().Index)

	log.Fatalln(app.Listen(":8000"))
}
