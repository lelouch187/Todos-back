package router

import (
	"awesomeProject5/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/columns")

	api.Get("/", controllers.FindColumns)

	api.Post("/", controllers.CreateColumnHandler)

}
