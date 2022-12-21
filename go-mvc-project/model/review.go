package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 리뷰
type Review struct {
	MenuName  string             `json:"menuname" bson:"menuname"`
	Star      int                `json:"star" bson:"star"`
	Comment   string             `json:"comment" bson:"comment"`
	CreatedAt primitive.DateTime `json:"createdat" bson:"createdat"`
	IsDetele  bool               `json:"isdelete" bson:"isdelete"`
}
