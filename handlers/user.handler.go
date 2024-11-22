package handlers

import (
	"github.com/RikiLaNeko/go-postgres-jwt-auth-api/initializers"
	"github.com/RikiLaNeko/go-postgres-jwt-auth-api/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// GetMeHandler handles the request to get the authenticated user's information.
// It retrieves the user information from the context and returns it in the response.
func GetMeHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

// GetUsersHandler handles the request to get a list of users.
// It supports pagination through the "page" and "limit" query parameters.
// It retrieves the users from the database, converts them to UserResponse format, and returns them in the response.
func GetUsersHandler(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var users []models.User
	results := initializers.DB.Limit(intLimit).Offset(offset).Find(&users)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, models.UserResponse{
			ID:        *user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      *user.Role,
			Photo:     *user.Photo,
			Provider:  *user.Provider,
			CreatedAt: *user.CreatedAt,
			UpdatedAt: *user.UpdatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(userResponses), "users": userResponses})
}
