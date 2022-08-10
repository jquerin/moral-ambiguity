package auth

import (
	"github.com/jquerin/moral-ambiguity/pkg/common/config"
	"github.com/jquerin/moral-ambiguity/pkg/common/models"

	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (h DbHandler) getUserByEmail(e string) (*models.User, error) {
	var user models.User
	if err := h.DB.Where(&models.User{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (h DbHandler) getUserByUsername(u string) (*models.User, error) {
	var user models.User
	if err := h.DB.Where(&models.User{Username: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Login login with provided credentials
// @Summary Login user
// @Description Login user with provided credentials. Returns JWT Token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "Login User"
// @Success 200 {object} models.ResponseHTTP{data=string}
// @Failure 400 {object} models.ResponseHTTP{}
// @Failure 401 {object} models.ResponseHTTP{}
// @Failure 500 {object} models.ResponseHTTP{}
// @Router /api/auth/login [post]
func (h DbHandler) Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var input LoginInput
	var ud UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseHTTP{
			Success: false,
			Message: "Error on login request",
			Data:    err.Error(),
		})
	}
	identity := input.Identity
	pass := input.Password

	email, err := h.getUserByEmail(identity)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseHTTP{
			Success: false,
			Message: "Issue with email",
			Data:    err.Error(),
		})
	}

	user, err := h.getUserByUsername(identity)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseHTTP{
			Success: false,
			Message: "Issue with username",
			Data:    err.Error(),
		})
	}

	if email == nil && user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseHTTP{
			Success: false,
			Message: "User not found",
			Data:    nil,
		})
	}

	if email == nil {
		ud = UserData{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		}
	} else {
		ud = UserData{
			ID:       email.ID,
			Username: email.Username,
			Email:    email.Email,
			Password: email.Password,
		}
	}

	if !CheckPasswordHash(pass, ud.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ResponseHTTP{
			Success: false,
			Message: "Incorrect password",
			Data:    nil,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.GetValue("SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseHTTP{
			Success: false,
			Message: "Issue generating token",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseHTTP{
		Success: true,
		Message: "Successful login",
		Data:    t,
	})
}
