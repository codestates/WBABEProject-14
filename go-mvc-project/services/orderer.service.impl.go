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
func (o *OrdererServiceImpl) GetAllOrder(sort string) ([]model.Menu, error) {
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
	fmt.Println(results[0]["avg_grade"])

	avgGrade := results[0]["avg_grade"].(float64)
	return avgGrade, reivewlist, nil
}

/* 리뷰 작성 */
func (o *OrdererServiceImpl) CreateReview(review *model.Review, id string) error {
	var order *model.Order

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	o.orderCollection.FindOne(o.ctx, filter).Decode(&order)

	/* 예외처리 조건 : 주문 상태가 5 (배달완료) 가 아니거나 이미 리뷰가 존재하는 주문이라면 + */
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
	_, err := o.orderCollection.UpdateByID(o.ctx, objId, query)

	return err
}
