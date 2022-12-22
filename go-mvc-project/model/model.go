package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	Client *mongo.Client
}

func NewModel(mongoUrl string) (*mongo.Collection, *mongo.Collection, *mongo.Collection, error) {
	r := &Model{}

	var err error
	mgUrl := mongoUrl
	if r.Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, nil, nil, err
	} else if err := r.Client.Ping(context.Background(), nil); err != nil {
		return nil, nil, nil, err
	}
	colMenu := r.Client.Database("go-mvc-project").Collection("Menu")
	colOrder := r.Client.Database("go-mvc-project").Collection("Order")
	colReview := r.Client.Database("go-mvc-project").Collection("Review")
	fmt.Println("Mongo DB Successful Connected")

	return colMenu, colOrder, colReview, nil
}
