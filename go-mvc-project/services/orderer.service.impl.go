package services

import (
	"context"
	"wba/go-mvc-procjet/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrdererServiceImpl struct {
	ordercollection *mongo.Collection
	ctx             context.Context
}

func NewOrdererService(orderercollection *mongo.Collection, ctx context.Context) OrdererService {
	return &OrdererServiceImpl{
		ordercollection: orderercollection,
		ctx:             ctx,
	}
}

func (o *OrdererServiceImpl) CreateOrder(order *model.Order) error {
	_, err := o.ordercollection.InsertOne(o.ctx, order)
	return err
}
