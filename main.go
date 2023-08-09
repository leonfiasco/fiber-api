package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leonfiasco/fiber-api/database"
	"github.com/leonfiasco/fiber-api/routes"
)

func welcome(c * fiber.Ctx) error{
return c.SendString(("welcome to my awesome API"))
}

func setupRoutes(app *fiber.App) {
	// welcome endpoint
	app.Get("/api", welcome)
	// user endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)


	// products endpoints
	app.Get("/api/products")
	app.Post("/api/products", routes.CreateProduct)
}


func main()  {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":2402"))
}