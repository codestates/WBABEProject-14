package model

import (
	"time"
)

// 주문
type Order struct {
	MenuName      string    `json:"menuname" bson:"menuname" validate:"required"`
	Phone         string    `json:"phone" bson:"phone" validate:"required"`
	Address       string    `json:"address" bson:"address" validate:"required"`
	CreatedAt     time.Time `json:"createdat" bson:"createdat" validate:"required"`
	IsDelete      bool      `json:"isdelete" bson:"isdelete" validate:"required"`
	IsExistReview bool      `json:"isexistreview" bson:"isexistreview" validate:"required"`
	Status        int       `json:"status" bson:"status" validate:"required"`
}
