package products

import (
	"github.com/jquerin/moral-ambiguity/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

// GetAllProducts is a function to retrieve all products data from database
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseHTTP{data=[]models.Product}
// @Failure 503 {object} models.ResponseHTTP{}
// @Router /api/products [get]
func (h handler) GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product

	if result := h.DB.Find(&products); result.Error != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(models.ResponseHTTP{
			Success: false,
			Message: result.Error.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseHTTP{
		Success: true,
		Message: "Successfully retrieved all products",
		Data:    products,
	})
}
