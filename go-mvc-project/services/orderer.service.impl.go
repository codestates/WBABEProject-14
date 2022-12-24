package services

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wba/go-mvc-procjet/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrdererServiceImpl struct {
	orderCollection  *mongo.Collection
	reviewCollection *mongo.Collection
	menuCollection   *mongo.Collection
	ctx              context.Context
}

func NewOrdererService(mc *mongo.Collection, oc *mongo.Collection, rc *mongo.Collection, ctx context.Context) (OrdererService, error) {
	return &OrdererServiceImpl{
		menuCollection:   mc,
		orderCollection:  oc,
		reviewCollection: rc,
		ctx:              ctx,
	}, nil
}

/* 주문 생성 */
func (o *OrdererServiceImpl) CreateOrder(order *model.Order) error {
	order.CreatedAt = time.Now()
	_, err := o.orderCollection.InsertOne(o.ctx, order)
	return err
}

/* 모든 메뉴 리스트 조회 */
func (o *OrdererServiceImpl) GetAllMenu(sort string) ([]model.Menu, error) {
	//sort = [recommend, grade, reorder, createdat]

	filter := bson.M{"isdelete": false}
	opts := options.Find().SetSort(bson.D{{Key: sort, Value: -1}})

	var menulist []model.Menu

	if corsur, err := o.menuCollection.Find(o.ctx, filter, opts); err != nil {
		panic(err)
	} else if err := corsur.All(o.ctx, &menulist); err != nil {
		panic(err)
	}
	return menulist, nil
}

/* 특정 메뉴에 대한 리뷰들 조회 */
func (o *OrdererServiceImpl) GetAllReiview(menuName string) (float64, []model.Review, error) {

	filter := bson.M{"menuname": menuName, "isdelete": false}
	var reivewlist []model.Review
	if corsur, err := o.reviewCollection.Find(o.ctx, filter); err != nil {
		panic(err)
	} else if err := corsur.All(o.ctx, &reivewlist); err != nil {
		panic(err)
	}

	/* 그룹 스테이지 생성 - 메뉴이름 기준 */
	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$menuname"},
			{Key: "avg_grade", Value: bson.D{{Key: "$avg", Value: "$grade"}}},
			{Key: "type_total", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}}

	/* 집계함수에 그룹스테이지 조건 대입 */
	cursor, err := o.reviewCollection.Aggregate(o.ctx, mongo.Pipeline{groupStage})
	if err != nil {
		panic(err)
	}

	/* 결과 표시 */
	var results []bson.M
	if err = cursor.All(o.ctx, &results); err != nil {
		panic(err)
	}
	avgGrade := results[0]["avg_grade"].(float64)

	return avgGrade, reivewlist, nil
}

/* 리뷰 작성 */
func (o *OrdererServiceImpl) CreateReview(review *model.Review, id string) error {
	var order *model.Order

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	o.orderCollection.FindOne(o.ctx, filter).Decode(&order)

	/* 예외처리 조건 : 주문 상태가 5 (배달완료) 가 아니거나 이미 리뷰가 존재하는 주문이라면 */
	if order.IsExistReview || order.Status != 5 {
		return errors.New("리뷰를 작성할 수 없습니다")
	}

	review.CreatedAt = time.Now()
	review.IsDetele = false
	if _, err := o.reviewCollection.InsertOne(o.ctx, review); err != nil {
		return err
	}
	/* 해당 주문에 대해서 리뷰 작성완료 업데이트 */
	query := bson.M{
		"$set": bson.M{
			"isexistreview": true,
		},
	}
	if _, err := o.orderCollection.UpdateByID(o.ctx, objId, query); err != nil {
		panic(err)
	}

	/* 메뉴 평점 업데이트 */
	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$menuname"},
			{Key: "avg_grade", Value: bson.D{{Key: "$avg", Value: "$grade"}}},
			{Key: "type_total", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}}

	/* 집계함수에 그룹스테이지 조건 대입 */
	cursor, err := o.reviewCollection.Aggregate(o.ctx, mongo.Pipeline{groupStage})
	if err != nil {
		panic(err)
	}

	/* 결과 표시 */
	var results []bson.M
	if err = cursor.All(o.ctx, &results); err != nil {
		panic(err)
	}
	avgGrade := results[0]["avg_grade"].(float64)
	query = bson.M{
		"$set": bson.M{
			"grade": avgGrade,
		},
	}
	o.menuCollection.FindOneAndUpdate(o.ctx, bson.M{"menuname": review.MenuName}, query)

	return err
}

/* 메뉴 변경 */
func (o *OrdererServiceImpl) UpdateOrder(id string, menuname string) error {
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": objid}
	/* 해당 주문 상태 가져오기 */
	var or model.Order
	o.orderCollection.FindOne(o.ctx, filter).Decode(&or)
	if or.Status == 3 {
		return errors.New("해당 주문은 조리중입니다")
	} else if or.Status == 4 {
		return errors.New("해당 주문은 배달중입니다")
	}
	query := bson.M{
		"$set": bson.M{
			"menuname": menuname,
		},
	}
	result, error := o.orderCollection.UpdateByID(o.ctx, objid, query)
	fmt.Println(&result)
	return error
}

/* 주문 내역 조회 */
func (o *OrdererServiceImpl) GetOrders() ([]model.Order, []model.Order, error) {

	filter := bson.M{}
	opts := options.Find().SetSort(bson.D{{Key: "createdat", Value: -1}})

	var currentOrders []model.Order
	var pastOrders []model.Order
	cursor, err := o.orderCollection.Find(o.ctx, filter, opts)
	if err != nil {
		panic(err)
	}
	for cursor.Next(o.ctx) {
		var result model.Order
		if err := cursor.Decode(&result); err != nil {
			panic(err)
		}
		if result.Status == 5 { // 배달완료된 이전 주문들
			pastOrders = append(pastOrders, result)
		} else { //현재 주문들
			currentOrders = append(currentOrders, result)
		}
	}
	return currentOrders, pastOrders, nil
}
