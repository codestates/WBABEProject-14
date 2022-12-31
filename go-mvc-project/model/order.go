package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status int

/* 주문 */
type Order struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" example:"63a73a1c8d989838729bc114"`
	MenuName      string             `json:"menuname" bson:"menuname"  example:"떡볶이"`
	Phone         string             `json:"phone" bson:"phone" example:"01020221205"`
	Address       string             `json:"address" bson:"address" example:"서울시 강남구 위메이드동 1212-202 101호"`
	CreatedAt     time.Time          `json:"createdat" bson:"createdat" example:"2022-12-24T16:17:12.793+00:00"`
	IsDelete      bool               `json:"isdelete" bson:"isdelete" example:"false"`
	IsExistReview bool               `json:"isexistreview" bson:"isexistreview" example:"false"`
	/*
		Status의 경우 상수를 통해 입력받는다면 사용하는 곳에서 코드를 읽기에 쉬워집니다.
		관련 키워드 : Enum, Const
	*/
	Status      Status `json:"status" bson:"status" example:"0"`
	OrderNumber int    `json:"ordernumber" bson:"ordernumber" example:"11"`
}

/* 상태 */
const (
	Ordering Status = iota
	OrderCancel
	AddOrder
	Cooking
	Delivering
	Complete
)
