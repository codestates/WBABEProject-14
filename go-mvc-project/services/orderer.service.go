package services

import "wba/go-mvc-procjet/model"

type OrdererService interface {
	CreateOrder(*model.Order) error
	GetAllOrder(sort string) ([]model.Menu, error)
	GetAllReiview(menuName string) (float64, []model.Review, error)
	CreateReview(*model.Review, string) error
}
