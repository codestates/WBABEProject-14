package services

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wba/go-mvc-procjet/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TakerServiceImpl struct {
	menuCollection *mongo.Collection
	ctx            context.Context
}

func NewTakerService(mc *mongo.Collection, ctx context.Context) (TakerService, error) {
	return &TakerServiceImpl{
		menuCollection: mc,
		ctx:            ctx,
	}, nil
}

/* 메뉴 등록 */
func (o *TakerServiceImpl) CreateMenu(menu *model.Menu) error {
	/* 메뉴이름 중복 검사 */
	var isExistMenu model.Menu
	filter := bson.D{bson.E{Key: "menuname", Value: menu.MenuName}}
	o.menuCollection.FindOne(o.ctx, filter).Decode(&isExistMenu)
	if len(isExistMenu.MenuName) > 0 {
		fmt.Println(isExistMenu.MenuName)
		return errors.New("동일한 메뉴 이름이 존재합니다")
	}
	menu.CreatedAt = time.Now()
	_, err := o.menuCollection.InsertOne(o.ctx, menu)
	return err
}

/* 메뉴 수정 */
func (o *TakerServiceImpl) UpdateMenu(menu *model.Menu) error {
	filter := bson.M{"menuname": menu.MenuName}
	query := bson.M{
		"$set": bson.M{
			"price":       menu.Price,
			"origin":      menu.Origin,
			"orderstatus": menu.OrderStatus,
		},
	}
	_, err := o.menuCollection.UpdateOne(o.ctx, filter, query)

	return err
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
func (o *TakerServiceImpl) UpdateMenuRecommend(menu *model.Menu) error {
	var result model.Menu
	filter := bson.M{"menuname": menu.MenuName}
	o.menuCollection.FindOne(o.ctx, filter).Decode(&result)
	query := bson.M{
		"$set": bson.M{
			"recommend": !result.Recommend,
		},
	}
	_, err := o.menuCollection.UpdateOne(o.ctx, filter, query)
	return err
}
