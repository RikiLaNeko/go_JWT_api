package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wpcodevo/go-postgres-jwt-auth-api/handlers"
)

// SetupRoutes sets up the application routes.
// It defines the routes for authentication and user-related operations.
// The routes are grouped and protected by middleware functions.
func SetupRoutes(app fiber.Router) {
	// Group routes under /auth for authentication-related operations.
	authRoutes := app.Group("/auth")

	// Route for user registration.
	authRoutes.Post("/register", handlers.SignUpUser)

	// Route for user login.
	authRoutes.Post("/login", handlers.SignInUser)

	// Route for user logout, protected by the DeserializeUser middleware.
	authRoutes.Get("/logout", DeserializeUser, handlers.LogoutUser)

	// Route to get the authenticated user's information, protected by the DeserializeUser middleware.
	app.Get("/users/me", DeserializeUser, handlers.GetMeHandler)

	// Route to get all users, protected by the DeserializeUser and allowedRoles middleware.
	app.Get("/users/", DeserializeUser, allowedRoles([]string{"admin", "moderator"}), handlers.GetUsersHandler)
}
