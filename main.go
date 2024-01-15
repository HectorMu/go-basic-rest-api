package main

import (
	"fmt"
	"os"

	"github.com/HectorMu/go-rest-api/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	println("Hello world")

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", controllers.GetUsers)

	app.Get("/every", controllers.HandleEvery)

	app.Get("/mapped", controllers.HandleMap)

	app.Get("/filtered", controllers.HandleFilter)

	app.Post("/", controllers.SaveUser)

	app.Delete("/:id", controllers.RemoveUser)

	app.Listen(":3000")
	fmt.Println("Server running on port 3000")
}
