package core

// Secondary Port
type OrderRepository interface {
	Save(order Order) error
}