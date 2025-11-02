package handler

import (
	"strconv"

	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/model"
	"github.com/AliasgharHeidari/mobile-numbers-mine/internal/service"
	"github.com/gofiber/fiber/v2"
)

// GetUserList 	godoc
// @Summary 	List users
// @Tags 		Users
// @Description Get list of users
// @Produce 	json
// @Success 	200 {object} model.GetUserListSuccessResponse
// @Failure 	401 {object} model.GetUserListFailureResponse
// @Security 	BearerAuth
// @Router 		/user [get]
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
// @Success 	201 {object} model.CreateUserSuccessResponse "user created successfully"
// @Failure 	400 {object} model.CreateUserFailureResponse "invalid request body"
// @Failure 	401 {object} model.StatusUnauthorizedResponse "unauthorized"
// @Security 	BearerAuth
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

// GetUserByID 	godoc
// @Summary 	Get user by ID
// @Tags		Users
// @Description Get user by ID
// @Produce 	json
// @Param 		id path string true "User ID"
// @Success 	200 {object} model.GetUserByIDSuccessResponse "user retrieved successfully"
// @Failure 	401 {object} model.StatusUnauthorizedResponse "unauthorized"
// @Failure 	404 {object} model.StatusNotFoundResponse "not found"
// @Security 	BearerAuth
// @Router 		/user/{id} [get]
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
// @Param user body model.UpdateUserByIDRequest true "user payload"
// @Success 200 {object} model.UpdateUserByIDSuccessResponse "user updated successfully"
// @Failure 400 {object} model.UpdateUserByIDFailureResponse "bad request"
// @Failure 401 {object} model.StatusUnauthorizedResponse "unauthorized"
// @Failure 404 {object} model.StatusNotFoundResponse "not found"
// @Security BearerAuth
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
// @Success 204 {object} model.DeleteUserByIDSuccessResponse "user deleted successfully"
// @Failure 401 {object} model.StatusUnauthorizedResponse "unauthorized"
// @Failure 404 {object} model.StatusNotFoundResponse "not found"
// @Security BearerAuth
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
// @Param number body model.AddMobileNumberRequest true "mobile number payload"
// @Success 201 {object} model.AddMobileNumberSuccessResponse "mobile number added successfully"
// @Failure 400 {object} model.AddMobileNumberFailureResponse "bad request"
// @Failure 401 {object} model.StatusUnauthorizedResponse "unauthorized"
// @Security BearerAuth
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
// @Success 204 {object} model.DeleteMobileNumberSuccessResponse "mobile number deleted successfully"
// @Failure 401 {object} model.StatusUnauthorizedResponse "unauthorized"
// @Failure 404 {object} model.StatusNotFoundResponse "not found"
// @Security BearerAuth
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
