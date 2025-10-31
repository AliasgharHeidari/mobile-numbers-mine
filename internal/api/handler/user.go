package handler

import (
	"strconv"

	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/model"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/service"
	"github.com/gofiber/fiber/v2"
)

// GetUserList godoc
// @Summary List users
// @Tags Users
// @Description Get list of users
// @Produce json
// @Success 200 {array} object "users"
// @Failure 401 {object} map[string]interface{} "unauthorized"
// @Security ApiKeyAuth
// @Router /user [get]
func GetUserList(c *fiber.Ctx) error {
	usersList, err := service.GetUserList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user list",
		})
	}
	return c.Status(fiber.StatusOK).JSON(usersList)
}

// CreateUser 	godoc
// @Summary 	Create user
// @Tags 		Users
// @Description Create a new user
// @Accept 		json
// @Produce 	json
// @Param 		newUserDetails body model.CreateUserRequest true "new user payload"
// @Param 		Authorization header string true "Bearer {token}"
// @Success 	201 {object} model.CreateUserSuccessResponse "user created successfully"
// @Failure 	400 {object} map[string]interface{} "bad request"
// @Failure 	401 {object} map[string]interface{} "unauthorized"
// @Security 	ApiKeyAuth
// @Router 		/user [post]
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

// GetUserByID godoc
// @Summary Get user by ID
// @Tags Users
// @Description Get user by ID
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} object "user"
// @Failure 401 {object} map[string]interface{} "unauthorized"
// @Failure 404 {object} map[string]interface{} "not found"
// @Security ApiKeyAuth
// @Router /user/{id} [get]
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

// UpdateUserByID godoc
// @Summary Update user
// @Tags Users
// @Description Update user by ID
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body object true "user payload"
// @Success 200 {object} object "updated user"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 401 {object} map[string]interface{} "unauthorized"
// @Failure 404 {object} map[string]interface{} "not found"
// @Security ApiKeyAuth
// @Router /user/{id} [put]
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

// DeleteUserByID godoc
// @Summary Delete user
// @Tags Users
// @Description Delete user by ID
// @Param id path string true "User ID"
// @Success 204 {string} string "no content"
// @Failure 401 {object} map[string]interface{} "unauthorized"
// @Failure 404 {object} map[string]interface{} "not found"
// @Security ApiKeyAuth
// @Router /user/{id} [delete]
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
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user deleted successfully",
	})

}

// AddMobileNumber godoc
// @Summary Add mobile number
// @Tags MobileNumbers
// @Description Add a mobile number to user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param number body object true "mobile number payload"
// @Success 201 {object} object "added number"
// @Failure 400 {object} map[string]interface{} "bad request"
// @Failure 401 {object} map[string]interface{} "unauthorized"
// @Security ApiKeyAuth
// @Router /user/{id}/mobile-number [post]
func AddMobileNumber(c *fiber.Ctx) error {
	userID := c.Params("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user ID",
		})
	}
	var mobileNumber model.MobileNumber
	if err := c.BodyParser(&mobileNumber); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := service.AddMobileNumber(id, mobileNumber); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Mobile Number addded successfully",
	})
}

// DeleteMobileNumber godoc
// @Summary Delete mobile number
// @Tags MobileNumbers
// @Description Delete a mobile number of a user
// @Param id path string true "User ID"
// @Param number path string true "mobile number"
// @Success 204 {string} string "no content"
// @Failure 401 {object} map[string]interface{} "unauthorized"
// @Failure 404 {object} map[string]interface{} "not found"
// @Security ApiKeyAuth
// @Router /user/{id}/mobile-number/{number} [delete]
func DeleteMobileNumber(c *fiber.Ctx) error {
	userID := c.Params("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user ID",
		})
	}

	number := c.Params("number")

	if err := service.DeleteMobileNumber(id, number); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "MobileNumber deleted successfully",
	})
}
