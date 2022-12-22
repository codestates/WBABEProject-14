package model

import (
	"time"
)

// 메뉴
type Menu struct {
	MenuName    string    `json:"menuname" bson:"menuname" validate:"required"`
	OrderStatus bool      `json:"orderstatus" bson:"orderstatus" validate:"required"`
	LimitOrder  int       `json:"limitorder" bson:"limitorder" validate:"required"`
	Origin      string    `json:"origin" bson:"origin" validate:"required"`
	Price       int       `json:"price" bson:"price" validate:"required"`
	Spice       int       `json:"spice" bson:"spice" validate:"required"`
	IsDelete    bool      `json:"isdelete" bson:"isdelete" validate:"required"`
	Reorder     int       `json:"reorder" bson:"reorder" validate:"required"`
	Recommend   bool      `json:"recommend" bson:"recommend" validate:"required"`
	CreatedAt   time.Time `json:"createdat" bson:"createdat" validate:"required"`
	Grade       float64   `json:"grade" bson:"grade"`
}
