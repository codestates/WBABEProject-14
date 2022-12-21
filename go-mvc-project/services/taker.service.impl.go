package services

import (
	"context"
	"wba/go-mvc-procjet/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type TakerServiceImpl struct {
	takercollection *mongo.Collection
	ctx             context.Context
}

func NewTakerService(takercollection *mongo.Collection, ctx context.Context) TakerService {
	return &TakerServiceImpl{
		takercollection: takercollection,
		ctx:             ctx,
	}
}

func (o *TakerServiceImpl) CreateMenu(order *model.Menu) error {
	_, err := o.takercollection.InsertOne(o.ctx, order)
	return err
}
