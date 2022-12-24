package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 리뷰
type Review struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" example:"63a73a1c8d989838729bc114"`
	MenuName  string             `json:"menuname" bson:"menuname" example:"떡볶이"`
	Grade     float64            `json:"grade" bson:"grade" binding:"required" example:"5.0"`
	Comment   string             `json:"comment" bson:"comment" binding:"required" example:"맛있었습니다."`
	CreatedAt time.Time          `json:"createdat" bson:"createdat" example:"2022-12-24T17:42:53.949+00:00"`
	IsDetele  bool               `json:"isdelete" bson:"isdelete" example:"false"`
}
