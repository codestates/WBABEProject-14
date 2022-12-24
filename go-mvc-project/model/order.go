package model

import (
	"time"
)

// 주문
type Order struct {
	MenuName      string    `json:"menuname" bson:"menuname"`
	Phone         string    `json:"phone" bson:"phone"`
	Address       string    `json:"address" bson:"address"`
	CreatedAt     time.Time `json:"createdat" bson:"createdat"`
	IsDelete      bool      `json:"isdelete" bson:"isdelete"`
	IsExistReview bool      `json:"isexistreview" bson:"isexistreview"`
	Status        int       `json:"status" bson:"status"`
}
