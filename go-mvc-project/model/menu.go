package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
1. example 넣어주신 점 좋습니다.

2. 필수 데이터라면 binding required를 이용해 필수값으로 변경할 수 있겠습니다.
*/

// 메뉴
type Menu struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" example:"63a73a1c8d989838729bc114"`
	StoreNumber int                `json:"storenumber" bson:"storenumber" example:"001"`
	MenuName    string             `json:"menuname" bson:"menuname" example:"떡볶이"`
	OrderStatus bool               `json:"orderstatus" bson:"orderstatus" example:"true"`
	LimitOrder  int                `json:"limitorder" bson:"limitorder" example:"100"`
	Origin      string             `json:"origin" bson:"origin" example:"국내산"`
	Price       int                `json:"price" bson:"price" example:"150000"`
	Spice       int                `json:"spice" bson:"spice" example:"3"`
	IsDelete    bool               `json:"isdelete" bson:"isdelete" example:"false"`
	Reorder     int                `json:"reorder" bson:"reorder" example:"30"`
	Recommend   bool               `json:"recommend" bson:"recommend" example:"true"`
	/*
		일반적으로 created_at, updated_at은 한 세트로 이루어집니다.
		그 이유는 추후 값 변경의 히스토리 추적을 위해서 updated_at이 용이합니다. 또한, 가장 최근에 변경된 순으로 정렬을 하는 경우에도 용이하구요.
	*/
	CreatedAt time.Time `json:"createdat" bson:"createdat" example:"2022-12-24T16:17:12.793+00:00"`
	UpdatedAt time.Time `json:"updatedat" bson:"updatedat" example:"2022-12-25T16:17:12.793+00:00"`
	Grade     float64   `json:"grade" bson:"grade" example:"3.5"`
}
