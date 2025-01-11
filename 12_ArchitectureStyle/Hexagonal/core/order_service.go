package core

import (
	"errors"
)

// Primary Port
type OrderService interface {
	CreateOrder(order Order) error
}

// Application Core ที่คุยระหว่าง Primary Port และ Secondary Port
// Implement Implementation ของ Primary Port
// ประกาศรับ struct ของ Secondary Port
type OrderServiceImpl struct {
	repo OrderRepository
}

// สร้าง instance ของ Application Core
func NewOrderService(repo OrderRepository) OrderService {
	return &OrderServiceImpl{repo: repo}
}

// คุยผ่าน Primary Port เป็นช่องทางรับข้อมูลจาก CreateOrder
// คุยผ่าน Secondary Port ผ่าน service OrderServiceImpl
func (s *OrderServiceImpl) CreateOrder(order Order) error {
	//Business logic function
	if order.Total <= 0 {
		return errors.New("total must be positive")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}

	return nil
}