package main

import (
	"awesomeProject5/initializers"
	"awesomeProject5/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
