package adapters

import (
	"CleanExample/entities"
	"CleanExample/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpOrderHandler struct {
	orderUseCase usecases.OrderUseCase
}

func NewHttpOrderHandler(orderUseCase usecases.OrderUseCase) *HttpOrderHandler {
	return &HttpOrderHandler{orderUseCase: orderUseCase}
}

func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	order := entities.Order{}
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := h.orderUseCase.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}