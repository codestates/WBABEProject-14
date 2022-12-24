package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* 주문 */
type Order struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" example:"63a73a1c8d989838729bc114"`
	MenuName      string             `json:"menuname" bson:"menuname"  example:"떡볶이"`
	Phone         string             `json:"phone" bson:"phone" example:"01020221205"`
	Address       string             `json:"address" bson:"address" example:"서울시 강남구 위메이드동 1212-202 101호"`
	CreatedAt     time.Time          `json:"createdat" bson:"createdat" example:"2022-12-24T16:17:12.793+00:00"`
	IsDelete      bool               `json:"isdelete" bson:"isdelete" example:"false"`
	IsExistReview bool               `json:"isexistreview" bson:"isexistreview" example:"false"`
	Status        int                `json:"status" bson:"status" example:"0"`
	OrderNumber   int                `json:"ordernumber" bson:"ordernumber" example:"11"`
}

/* 상태 */
