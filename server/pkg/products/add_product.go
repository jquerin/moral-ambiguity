package products

import (
	"github.com/jquerin/moral-ambiguity/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type AddProductRequestBody struct {
	Name  string `json:"name"`
	Stock int32  `json:"stock"`
	Price int32  `json:"price"`
}

// AddProduct creates a new product
// @Summary Create new product
// @Description Create new product and store in the database
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Add product"
// @Success 200 {object} models.ResponseHTTP{data=models.Product}
// @Failure 400 {object} models.ResponseHTTP{}
// @Router /api/products [post]
func (h handler) AddProduct(c *fiber.Ctx) error {
	body := AddProductRequestBody{}

	// parse body, attach to AddProductRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var product models.Product

	product.Name = body.Name
	product.Stock = body.Stock
	product.Price = body.Price

	// insert new db entry
	if result := h.DB.Create(&product); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseHTTP{
			Success: false,
			Message: result.Error.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(models.ResponseHTTP{
		Success: true,
		Message: "Successfully created a new product in the system",
		Data:    product,
	})
}
