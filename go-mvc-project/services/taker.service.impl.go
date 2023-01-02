package services

import (
	"context"
	"errors"
	"time"
	"wba/go-mvc-procjet/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TakerServiceImplement struct {
	menuCollection  *mongo.Collection
	orderCollection *mongo.Collection
	ctx             context.Context
}

func NewTakerService(mc *mongo.Collection, oc *mongo.Collection, ctx context.Context) (TakerService, error) {
	return &TakerServiceImplement{
		menuCollection:  mc,
		orderCollection: oc,
		ctx:             ctx,
	}, nil
}

/* 메뉴 등록 */
func (o *TakerServiceImplement) CreateMenu(menu *model.Menu) error {
	/*
		각 가게별로 메뉴 이름은 중복될 수 있습니다. 그런 경우에는 어떻게 처리해야 할까요?
		e.g. A 가게의 김치찌개, B 가게의 김치찌개

		유니크함을 보장하기 위해서라면 다른 값을 활용하시는 것이 좋아보입니다. (MenuId와 같은)

		=> 각 가게를 구별할 수 있는것 , 가게 ID or 메뉴를 추가한 사용자의 ID 가 적합해보입니다.
		=> 메뉴 추가시 가게이름을 받는다고 가정하고, 메뉴 필드에 가게번호를 추가하면 가게별로 메뉴 이름 중복 문제를 해결할 수 있을 것이라 생각합니다.
		메뉴 추가시 요청 파라미터 1. 가게번호 2. 추가할 메뉴 이름
		해당 가게번호에 해당 메뉴가 존재하는지 검사
	*/

	/* 메뉴이름 중복 검사 */
	var isExistMenu model.Menu
	/* 삭제되지 않은 메뉴중 동일한 메뉴이름이 존재 */
	filter := bson.M{"storenumber": menu.StoreNumber, "menuname": menu.MenuName, "isdelete": false}
	o.menuCollection.FindOne(o.ctx, filter).Decode(&isExistMenu)
	if len(isExistMenu.MenuName) > 0 {
		return errors.New("동일한 메뉴 이름이 존재합니다")
	}
	menu.Grade = 0
	menu.Reorder = 0
	menu.IsDelete = false
	menu.CreatedAt = time.Now()
	menu.UpdatedAt = time.Now()
	_, err := o.menuCollection.InsertOne(o.ctx, menu)
	return err
}

/* 메뉴 수정 */
func (o *TakerServiceImplement) UpdateMenu(menuname string, menu *model.Menu) error {
	filter := bson.M{"menuname": menuname}
	query := bson.M{
		"$set": bson.M{
			"price":       menu.Price,
			"origin":      menu.Origin,
			"orderstatus": menu.OrderStatus,
			"updatedat":   time.Now(),
		},
	}
	result := o.menuCollection.FindOneAndUpdate(o.ctx, filter, query)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

/* 메뉴 삭제 */
func (o *TakerServiceImplement) DeleteMenu(menu *model.Menu) error {
	filter := bson.M{"menuname": menu.MenuName}
	query := bson.M{
		"$set": bson.M{
			"isdelete": true,
		},
	}
	_, err := o.menuCollection.UpdateOne(o.ctx, filter, query)
	return err
}

/* 금일 추천 메뉴 변경 */
func (o *TakerServiceImplement) UpdateMenuRecommend(menu *model.Menu) ([]*model.Menu, error) {
	var result model.Menu
	filter := bson.M{"menuname": menu.MenuName}
	o.menuCollection.FindOne(o.ctx, filter).Decode(&result)
	query := bson.M{
		"$set": bson.M{
			"recommend": !result.Recommend,
		},
	}
	if _, err := o.menuCollection.UpdateOne(o.ctx, filter, query); err != nil {
		return nil, err
	}
	recommend := bson.M{"recommend": true}
	opts := options.Find().SetSort(bson.D{{Key: "createdat", Value: -1}}) //최신순으로

	var recommendMenus []*model.Menu
	if cursor, err := o.menuCollection.Find(o.ctx, recommend, opts); err != nil {
		panic(err)
	} else if err := cursor.All(o.ctx, &recommendMenus); err != nil {
		panic(err)
	}
	return recommendMenus, nil
}

/* 현재 주문 내역 조회 */
func (o *TakerServiceImplement) GetOrderList() ([]*model.Order, error) {

	filter := bson.M{"status": 0}                                         //접수중인 주문만
	opts := options.Find().SetSort(bson.D{{Key: "createdat", Value: -1}}) //최신순으로

	var orderList []*model.Order

	if cursor, err := o.orderCollection.Find(o.ctx, filter, opts); err != nil {
		panic(err)
	} else if err := cursor.All(o.ctx, &orderList); err != nil {
		panic(err)
	}
	return orderList, nil
}

/* 각 메뉴별 주문 상태 변경 */
func (o *TakerServiceImplement) UpdateOrderStatus(menuname string, status model.Status) error {
	/* 해당 메뉴의 주문들 다음단계로 상태 저장 */
	if status > model.Delivering || status < model.Ordering {
		return errors.New("잘못된 요청입니다")
	}
	filter := bson.M{"menuname": menuname, "status": status}
	query := bson.M{
		"$set": bson.M{
			"status": status + 1,
		},
	}

	if _, err := o.orderCollection.UpdateMany(o.ctx, filter, query); err != nil {
		return err
	}
	return nil
}
