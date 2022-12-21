package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// 메뉴
type Menu struct {
	MenuName    string             `json:"menuname" bson:"menuname"`
	OrderStatus bool               `json:"orderstatus" bson:"orderstatus"`
	LimitOrder  int                `json:"limitorder" bson:"limitorder"`
	Origin      string             `json:"origin" bson:"origin"`
	Price       int                `json:"price" bson:"price"`
	Spice       int                `json:"spice" bson:"spice"`
	IsDelete    bool               `json:"isdelete" bson:"isdelete"`
	Reorder     int                `json:"reorder" bson:"reorder"`
	CreatedAt   primitive.DateTime `json:"createdat" bson:"createdat"`
}
