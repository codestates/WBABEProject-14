package services

import (
	"context"
	"errors"
	"strings"
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
func (o *OrdererServiceImpl) CreateOrder(order *model.Order) (int, error) {
	/* 메뉴 추가로 인한 신규주문 체크 (ObjectId 가 생성된 채로 넘어오는지)*/
	if order.ID != primitive.NilObjectID {
		order.ID = primitive.NewObjectID()
	}
	order.CreatedAt = time.Now()
	order.IsDelete = false
	order.Status = 0
	order.IsExistReview = false

	/* 일련번호 - 오늘 날짜 기준  ( 🔥 UTC 한국날짜 기준 -9 시간 생각하기 ) */
	standard := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, 0, 00, 00, 0, time.UTC)
	findQuery := bson.M{"createdat": bson.M{"$gte": standard, "$lt": order.CreatedAt}}

	orderNumber, countErr := o.orderCollection.CountDocuments(o.ctx, findQuery)

	if countErr != nil {
		panic(countErr)
	}
	order.OrderNumber = int(orderNumber)
	_, err := o.orderCollection.InsertOne(o.ctx, order)
	return int(orderNumber), err
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
	var menu *model.Menu
	o.menuCollection.FindOne(o.ctx, bson.M{"menuname": menuName}).Decode(&menu)

	return menu.Grade, reivewlist, nil
}

/* 리뷰 작성 */
func (o *OrdererServiceImpl) CreateReview(review *model.Review, orderId string) error {
	var order model.Order
	objId, _ := primitive.ObjectIDFromHex(orderId)
	filter := bson.M{"_id": objId}
	o.orderCollection.FindOne(o.ctx, filter).Decode(&order)
	/* 예외처리 조건 : 주문 상태가 5(배달완료)가 아니거나 이미 리뷰가 존재하는 주문이라면 */
	if order.IsExistReview || order.Status != 5 {
		return errors.New("리뷰를 작성할 수 없습니다")
	}
	review.MenuName = order.MenuName
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
func (o *OrdererServiceImpl) UpdateOrder(id string, flag int, menuname string) (int, error) {
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": objid}
	/* 해당 주문 상태 가져오기 */
	var or model.Order
	o.orderCollection.FindOne(o.ctx, filter).Decode(&or)

	/* 메뉴 추가 */
	if flag == 0 {
		/* 배달중일경우 */
		if or.Status == 4 {
			/* 신규 주문으로 처리 */
			or.MenuName = menuname
			or.CreatedAt = time.Now()
			if orderNumber, err := o.CreateOrder(&or); err != nil {
				return -1, errors.New("메뉴 추가에 실패하였습니다. 다시 시도해주세요")
			} else {
				return orderNumber, nil
			}
		} else {
			/* 메뉴 추가 성공 */
			str_slices := []string{or.MenuName, menuname}
			menunames := strings.Join(str_slices, ",")
			query := bson.M{
				"$set": bson.M{
					"menuname": menunames,
				},
			}
			if _, err := o.orderCollection.UpdateByID(o.ctx, objid, query); err != nil {
				return -1, err
			} else {
				return -1, nil
			}
		}
		/* 메뉴 변경 */
	} else if flag == 1 {
		/* 조리중 배달중 배달완료 에러처리 */
		if or.Status == 3 {
			return -1, errors.New("해당 주문은 조리중입니다")
		} else if or.Status == 4 {
			return -1, errors.New("해당 주문은 배달중입니다")
		} else if or.Status == 5 {
			return -1, errors.New("배달이 완료된 주문입니다 ")
		} else if or.MenuName == menuname {
			return -1, errors.New("동일한 메뉴로 변경할 수 없습니다")
		} else {
			/* 메뉴 변경 성공 */
			query := bson.M{
				"$set": bson.M{
					"menuname": menuname,
				},
			}
			if _, err := o.orderCollection.UpdateByID(o.ctx, objid, query); err != nil {
				return -1, err
			} else {
				return -1, nil
			}
		}
	} else {
		return -1, errors.New("잘못된 요청")
	}
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
