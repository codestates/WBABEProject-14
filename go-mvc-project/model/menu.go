package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 메뉴
type Menu struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" example:"63a73a1c8d989838729bc114"`
	MenuName    string             `json:"menuname" bson:"menuname" example:"떡볶이"`
	OrderStatus bool               `json:"orderstatus" bson:"orderstatus" example:"true"`
	LimitOrder  int                `json:"limitorder" bson:"limitorder" example:"100"`
	Origin      string             `json:"origin" bson:"origin" example:"국내산"`
	Price       int                `json:"price" bson:"price" example:"150000"`
	Spice       int                `json:"spice" bson:"spice" example:"3"`
	IsDelete    bool               `json:"isdelete" bson:"isdelete" example:"false"`
	Reorder     int                `json:"reorder" bson:"reorder" example:"30"`
	Recommend   bool               `json:"recommend" bson:"recommend" example:"true"`
	CreatedAt   time.Time          `json:"createdat" bson:"createdat" example:"2022-12-24T16:17:12.793+00:00"`
	Grade       float64            `json:"grade" bson:"grade" example:"3.5"`
}
