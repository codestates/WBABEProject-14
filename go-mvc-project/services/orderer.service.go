package services

import "wba/go-mvc-procjet/model"

type OrdererService interface {
	CreateOrder(*model.Order) error
}
