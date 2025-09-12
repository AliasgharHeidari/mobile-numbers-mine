package handler

import (
	"strconv"

	"github.com/Golang-Training-entry-3/mobile-numbers/internal/model"
	"github.com/Golang-Training-entry-3/mobile-numbers/internal/service"
	"github.com/gofiber/fiber/v2"
)

func GetUserList(c *fiber.Ctx) error {
	usersList, err := service.GetUserList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user list",
		})
	}
	return c.Status(fiber.StatusOK).JSON(usersList)
}

func CreateUser(c *fiber.Ctx) error {
	var newUser model.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	id, err := service.CreateUser(newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user_id": id,
	})
}

func GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := service.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func UpdateUserByID(c *fiber.Ctx) error {
	return c.SendString("Update User By ID")
}

func DeleteUserByID(c *fiber.Ctx) error {
	return c.SendString("Delete User By ID")
}

func AddMobileNumber(c *fiber.Ctx) error {
	return c.SendString("Add Mobile Number")
}

func DeleteMobileNumber(c *fiber.Ctx) error {
	return c.SendString("Delete Mobile Number")
}
