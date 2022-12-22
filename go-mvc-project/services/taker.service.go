package services

import "wba/go-mvc-procjet/model"

type TakerService interface {
	CreateMenu(*model.Menu) error
	UpdateMenu(*model.Menu) error
	DeleteMenu(*model.Menu) error
	UpdateMenuRecommend(*model.Menu) error
}
