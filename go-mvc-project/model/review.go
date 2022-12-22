package model

import (
	"time"
)

// 리뷰
type Review struct {
	MenuName  string    `json:"menuname" bson:"menuname" binding:"required"`
	Grade     float64   `json:"grade" bson:"grade" binding:"required"`
	Comment   string    `json:"comment" bson:"comment" binding:"required"`
	CreatedAt time.Time `json:"createdat" bson:"createdat"`
	IsDetele  bool      `json:"isdelete" bson:"isdelete"`
}
