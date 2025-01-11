package adapters

import (
	"gorm.io/gorm"
	"HexagonalExample/core"
)

// Secondary adapter
// รับสิ่งที่เป็น config ของฝั่ง adapter ตัวจริง (Database) 
type GormOrderRepository struct {
	db *gorm.DB
}

// สร้าง instance ของ Secondary adapter
func NewGormOrderRepository(db *gorm.DB) core.OrderRepository {
	return &GormOrderRepository{db: db}
}

// implement ตาม Secondary Port
func (r *GormOrderRepository) Save(order core.Order) error {
	if result := r.db.Create(&order); result.Error != nil {
		return result.Error
	}

	return nil
}