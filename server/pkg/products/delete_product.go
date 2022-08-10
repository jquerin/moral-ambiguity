package products

import (
	"fmt"

	"github.com/jquerin/moral-ambiguity/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

// DeleteProduct removes a product by ID
// @Summary Remove product by ID
// @Description Remove product by ID
// @Security ApiKeyAuth
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.ResponseHTTP{}
// @Failure 404 {object} models.ResponseHTTP{}
// @Failure 503 {object} models.ResponseHTTP{}
// @Router /api/products/{id} [delete]
func (h handler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		switch result.Error.Error() {
		case "record not found":
			return c.Status(fiber.StatusNotFound).JSON(models.ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Book with ID %v not found.", id),
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

	h.DB.Delete(&product)

	return c.Status(fiber.StatusOK).JSON(models.ResponseHTTP{
		Success: true,
		Message: "Successfully deleted product",
		Data:    nil,
	})
}
