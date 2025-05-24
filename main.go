package main

import (
	database "backend/config"
	"backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	
	app := fiber.New()
	
	database.InitDB()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	
	routes.SetupRoutes(app)


	log.Println("Server running on port 8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Failed to start server", err)
	}
}