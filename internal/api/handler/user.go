package handler

import (
	"strconv"

	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/model"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/service"
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
	userID := c.Params("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid User ID",
		})
	}

	var updatedUser model.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := service.UpdateUserByID(id, updatedUser); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

func DeleteUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid ID",
		})
	}

	if err := service.DeleteUserByID(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "user deleted successfully",
	})

}

func AddMobileNumber(c *fiber.Ctx) error {
	userID := c.Params("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid user ID",
		})
	}
	var mobileNumber model.MobileNumber
	if err := c.BodyParser(&mobileNumber); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid request body",
 		})
	}

	if err := service.AddMobileNumber(id, mobileNumber); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}

   return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Mobile Number addded successfully",
   })
}

func DeleteMobileNumber(c *fiber.Ctx) error {
	userID := c.Params("id")
	id , err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid user ID",
		})
	}

	number := c.Params("number")

	if err := service.DeleteMobileNumber(id, number); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "MobileNumber deleted successfully",
	})
}
