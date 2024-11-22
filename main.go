package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wpcodevo/go-postgres-jwt-auth-api/initializers"
)

// main is the entry point of the application.
func main() {
	// Load environment variables from the specified file.
	env, err := initializers.LoadEnv(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// Connect to the database using the loaded environment variables.
	initializers.ConnectDB(&env)

	// Create a new Fiber application instance.
	app := fiber.New()

	// Create a new Fiber instance for microservices.
	micro := fiber.New()

	// Mount the microservice instance to the main app under the /api route.
	app.Mount("/api", micro)

	// Use the logger middleware to log HTTP requests.
	app.Use(logger.New())

	// Use the CORS middleware to handle Cross-Origin Resource Sharing.
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	// Setup application routes.
	SetupRoutes(micro)

	// Define a health check endpoint.
	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "JSON Web Token Authentication and Authorization in Golang",
		})
	})

	// Handle all undefined routes with a 404 Not Found response.
	app.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

	// Start the Fiber application on port 8020.
	log.Fatal(app.Listen(":8020"))
}
