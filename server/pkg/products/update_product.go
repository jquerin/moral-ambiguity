package products

import (
	"github.com/jquerin/moral-ambiguity/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int32  `json:"stock"`
	Price int32  `json:"price"`
}

// UpdateProduct updates a product with new values
// @Summary Updates existing product
// @Description Updates existing product and stores in the database
// @Security ApiKeyAuth
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.ResponseHTTP{data=models.Product}
// @Failure 400 {object} models.ResponseHTTP{}
// @Router /api/products [put]
func (h handler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateProductRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseHTTP{
			Success: false,
			Message: result.Error.Error(),
			Data:    nil,
		})
	}

	product.Name = body.Name
	product.Stock = body.Stock
	product.Price = body.Price

	h.DB.Save(&product)

	return c.Status(fiber.StatusOK).JSON(models.ResponseHTTP{
		Success: true,
		Message: "Successfully updated product in the system",
		Data:    product,
	})
}
