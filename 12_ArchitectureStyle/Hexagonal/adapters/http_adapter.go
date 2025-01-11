package adapters

import (
	"HexagonalExample/core"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

// Primary adapter ฝั่งขาเข้า
// ต้องเชื่อมไปยัง Primary Port
// เวลาเรียกใช้ จะเรียกคำสั่งใน Primary Port
type HttpOrderHandler struct {
	service core.OrderService
}

// สร้าง instance ของ Primary adapter
func NewHttpOrderHandler(service core.OrderService) *HttpOrderHandler {
  return &HttpOrderHandler{service: service}
}

// ฟังก์ชันที่สร้างมาเพื่อเชื่อมกับ end point
// ไม่ได้ implement ตามใคร ใช้ชื่อเดียวกัน แค่กันความสับสน
func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order core.Order
	if err := c.BodyParser(&order); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// ใช้งานฟังก์ชันที่อยู่ใน Primary Port
	if err := h.service.CreateOrder(order); err != nil {
		fmt.Println("Hello")
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	fmt.Println("Hello")

	return c.Status(fiber.StatusCreated).JSON(order)
}