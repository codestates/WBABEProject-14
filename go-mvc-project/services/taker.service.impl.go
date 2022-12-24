package services

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wba/go-mvc-procjet/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TakerServiceImpl struct {
	menuCollection  *mongo.Collection
	orderCollection *mongo.Collection
	ctx             context.Context
}

func NewTakerService(mc *mongo.Collection, oc *mongo.Collection, ctx context.Context) (TakerService, error) {
	return &TakerServiceImpl{
		menuCollection:  mc,
		orderCollection: oc,
		ctx:             ctx,
	}, nil
}

/* 메뉴 등록 */
func (o *TakerServiceImpl) CreateMenu(menu *model.Menu) error {
	/* 메뉴이름 중복 검사 */
	var isExistMenu model.Menu
	/* 삭제되지 않은 메뉴중 동일한 메뉴이름이 존재 */
	filter := bson.M{"menuname": menu.MenuName, "isdelete": false}
	o.menuCollection.FindOne(o.ctx, filter).Decode(&isExistMenu)
	if len(isExistMenu.MenuName) > 0 {
		fmt.Println(isExistMenu.MenuName)
		return errors.New("동일한 메뉴 이름이 존재합니다")
	}
	menu.Grade = 0
	menu.Reorder = 0
	menu.IsDelete = false
	menu.CreatedAt = time.Now()
	_, err := o.menuCollection.InsertOne(o.ctx, menu)
	return err
}

/* 메뉴 수정 */
func (o *TakerServiceImpl) UpdateMenu(menuname string, menu *model.Menu) error {
	filter := bson.M{"menuname": menuname}
	query := bson.M{
		"$set": bson.M{
			"price":       menu.Price,
			"origin":      menu.Origin,
			"orderstatus": menu.OrderStatus,
		},
	}
	result := o.menuCollection.FindOneAndUpdate(o.ctx, filter, query)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

/* 메뉴 삭제 */
func (o *TakerServiceImpl) DeleteMenu(menu *model.Menu) error {
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
func (o *TakerServiceImpl) UpdateMenuRecommend(menu *model.Menu) ([]*model.Menu, error) {
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
func (o *TakerServiceImpl) GetOrderList() ([]*model.Order, error) {

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
func (o *TakerServiceImpl) UpdateOrderStatus(menuname string, status int) error {
	/* 해당 메뉴의 주문들 다음단계로 상태 저장 */
	if status > 4 || status < 0 {
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
