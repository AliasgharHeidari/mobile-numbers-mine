package handler

import (
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
	return c.SendString("Create User")
}

func GetUserByID(c *fiber.Ctx) error {
	return c.SendString("Get User By ID")
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
