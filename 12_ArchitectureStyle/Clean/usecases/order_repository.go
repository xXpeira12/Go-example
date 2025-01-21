package usecases

import "CleanExample/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}