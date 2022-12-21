package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 주문
type Order struct {
	MenuName  string             `json:"menuname" bson:"menuname"`
	Phone     string             `json:"phone" bson:"phone"`
	Address   string             `json:"address" bson:"address"`
	CreatedAt primitive.DateTime `json:"createdat" bson:"createdat"`
	Status    int                `json:"status" bson:"status"`
}
