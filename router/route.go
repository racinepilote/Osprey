package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/racinepilote/go_rest_api/handler"
	"github.com/racinepilote/go_rest_api/middleware"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/user", middleware.JWTMiddleware) // Apply middleware here
	// routes
	v1.Get("/", handler.GetAllUsers)
	v1.Get("/:id", handler.GetSingleUser)
	v1.Post("/", handler.CreateUser)
	v1.Put("/:id", handler.UpdateUser)
	v1.Delete("/:id", handler.DeleteUserByID)
}
