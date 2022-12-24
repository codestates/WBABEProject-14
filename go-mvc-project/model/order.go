package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 주문
type Order struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id" example:"63a73a1c8d989838729bc114"`
	MenuName      string             `json:"menuname" bson:"menuname"`
	Phone         string             `json:"phone" bson:"phone"`
	Address       string             `json:"address" bson:"address"`
	CreatedAt     time.Time          `json:"createdat" bson:"createdat"`
	IsDelete      bool               `json:"isdelete" bson:"isdelete"`
	IsExistReview bool               `json:"isexistreview" bson:"isexistreview"`
	Status        int                `json:"status" bson:"status"`
	OrderNumber   int                `json:"ordernumber" bson:"ordernumber"`
}
