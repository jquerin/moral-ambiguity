package users

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jquerin/moral-ambiguity/pkg/auth"
	"github.com/jquerin/moral-ambiguity/pkg/common/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func validUser(id string, p string, db *gorm.DB) bool {
	var user models.User
	db.First(&user, id)
	if user.Username == "" {
		return false
	}
	if !auth.CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}

// GetUserById is a function to get a user by ID
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.ResponseHTTP{data=[]models.User}
// @Failure 404 {object} models.ResponseHTTP{}
// @Failure 503 {object} models.ResponseHTTP{}
// @Router /api/users/{id} [get]
func (h DbHandler) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		switch result.Error.Error() {
		case "record not found":
			return c.Status(fiber.StatusNotFound).JSON(models.ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Product with ID %v not found.", id),
				Data:    nil,
			})
		default:
			return c.Status(fiber.StatusServiceUnavailable).JSON(models.ResponseHTTP{
				Success: false,
				Message: result.Error.Error(),
				Data:    nil,
			})
		}
	}

	// clear password
	user.Password = ""

	return c.Status(fiber.StatusOK).JSON(models.ResponseHTTP{
		Success: true,
		Message: "Successfully retrieved product",
		Data:    user,
	})
}

// AddUser creates a new user
// @Summary Create new user
// @Description Create new user and store in the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "Add user"
// @Success 201 {object} models.ResponseHTTP{data=models.User}
// @Failure 500 {object} models.ResponseHTTP{}
// @Router /api/users [post]
func (h DbHandler) AddUser(c *fiber.Ctx) error {
	body := models.User{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// empty Password is not valid
	if body.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseHTTP{
			Success: false,
			Message: "Password cannot be empty",
			Data:    nil,
		})
	}

	hash, err := hashPassword(body.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseHTTP{
			Success: false,
			Message: fmt.Sprintf("Failed to hash password: %s", err.Error()),
			Data:    nil,
		})
	}

	var user models.User

	user.Email = body.Email
	user.Username = body.Username
	user.Password = hash

	if err := h.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseHTTP{
			Success: false,
			Message: fmt.Sprintf("Failed to create user: %s", err.Error()),
			Data:    nil,
		})
	}

	// wipe password before returning data
	user.Password = ""

	return c.Status(fiber.StatusCreated).JSON(models.ResponseHTTP{
		Success: true,
		Message: "Successfully created new user",
		Data:    user,
	})
}

// UpdateUser updates a user
// @Summary Update a user
// @Description Updates the user's names field
// @Security ApiKeyAuth
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.ResponseHTTP{data=models.User}
// @Failure 500 {object} models.ResponseHTTP{}
// @Router /api/users/{id} [patch]
func (h DbHandler) UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		Names string `json:"names"`
	}

	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseHTTP{
			Success: false,
			Message: fmt.Sprintf("Review your input. Exception: %s", err.Error()),
			Data:    nil,
		})
	}

	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseHTTP{
			Success: false,
			Message: "Invalid token ID",
			Data:    nil,
		})
	}

	var user models.User

	if err := h.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseHTTP{
			Success: false,
			Message: fmt.Sprintf("Failed to retrieve user: %s", err.Error()),
			Data:    nil,
		})
	}

	user.Names = uui.Names
	if err := h.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseHTTP{
			Success: false,
			Message: fmt.Sprintf("Failed to update user: %s", err.Error()),
			Data:    nil,
		})
	}

	user.Password = ""
	return c.Status(fiber.StatusOK).JSON(models.ResponseHTTP{
		Success: true,
		Message: "Successfully updated user's names",
		Data:    user,
	})
}

// DeleteUser Deletes a user
// @Summary Deletes a user from the database
// @Description Deletes a user from the database
// @Security ApiKeyAuth
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.ResponseHTTP{}
// @Failure 500 {object} models.ResponseHTTP{}
// @Router /api/users/{id} [delete]
func (h DbHandler) DeleteUser(c *fiber.Ctx) error {
	type PasswordInput struct {
		Password string `json:"password"`
	}
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseHTTP{
			Success: false,
			Message: fmt.Sprintf("Review your input. Exception: %s", err.Error()),
			Data:    nil,
		})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseHTTP{
			Success: false,
			Message: "Invalid token ID",
			Data:    nil,
		})
	}

	if !validUser(id, pi.Password, h.DB) {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseHTTP{
			Success: false,
			Message: "Not a valid user",
			Data:    nil,
		})
	}

	var user models.User
	// find user
	if err := h.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseHTTP{
			Success: false,
			Message: fmt.Sprintf("Failed to retrieve user: %s", err.Error()),
			Data:    nil,
		})
	}

	// delete user from DB
	if err := h.DB.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseHTTP{
			Success: false,
			Message: fmt.Sprintf("Failed to delete user: %s", err.Error()),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseHTTP{
		Success: true,
		Message: "Successfully deleted user",
		Data:    nil,
	})
}
