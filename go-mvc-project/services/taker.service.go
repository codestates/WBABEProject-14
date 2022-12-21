package services

import "wba/go-mvc-procjet/model"

type TakerService interface {
	CreateMenu(*model.Menu) error
}
