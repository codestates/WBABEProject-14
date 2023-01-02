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

/*
네이밍이 직관적이지 못해 보입니다. Impl은 무엇을 의미하나요?
*/
type OrdererServiceImplement struct {
	orderCollection  *mongo.Collection
	reviewCollection *mongo.Collection
	menuCollection   *mongo.Collection
	ctx              context.Context
}

func NewOrdererService(mc *mongo.Collection, oc *mongo.Collection, rc *mongo.Collection, ctx context.Context) (OrdererService, error) {
	return &OrdererServiceImplement{
		menuCollection:   mc,
		orderCollection:  oc,
		reviewCollection: rc,
		ctx:              ctx,
	}, nil
}

/* 주문 생성 */
func (o *OrdererServiceImplement) CreateOrder(order *model.Order) (int, error) {
	/* 메뉴 추가로 인한 신규주문 체크 (ObjectId 가 생성된 채로 넘어오는지)*/
	if order.ID != primitive.NilObjectID {
		order.ID = primitive.NewObjectID()
	}
	order.CreatedAt = time.Now()
	order.IsDelete = false
	order.Status = 0
	order.IsExistReview = false

	/* 일련번호 - 오늘 날짜 기준  ( 🔥 UTC 한국날짜 기준 -9 시간 생각하기 ) */
	/*
		하루를 빼는 이유는 무엇인가요? UTC와 한국시간의 차이라면 9시간을 더하거나 뺴주어야 할 것 같습니다.
		= 일련번호는 하루 단위로 초기화 되기 때문에 하루 전 00시 ~ 현재 시간을 기준으로 하기 때문입니다.
	*/
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
func (o *OrdererServiceImplement) GetAllMenu(sort string) ([]model.Menu, error) {
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
func (o *OrdererServiceImplement) GetAllReiview(menuName string) (float64, []model.Review, error) {

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
func (o *OrdererServiceImplement) CreateReview(review *model.Review, orderId string) error {
	var order model.Order
	objId, _ := primitive.ObjectIDFromHex(orderId)
	filter := bson.M{"_id": objId}
	o.orderCollection.FindOne(o.ctx, filter).Decode(&order)
	/* 예외처리 조건 : 주문 상태가 5(배달완료)가 아니거나 이미 리뷰가 존재하는 주문이라면 */

	/*
		모델에서도 언급하였지만 Status 처럼 여러 상태값을 가지는 경우에는 일반적으로 Enum을 활용하는 편이
		가독성 측면에서 좋습니다. 현재와 같은 경우 5번이 무엇인지를 의미하는데 알기가 힘듭니다.
	*/
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

	return nil
}

/* 메뉴 변경 */
func (o *OrdererServiceImplement) UpdateOrder(id string, flag int, menuname string) (int, error) {
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": objid}
	/* 해당 주문 상태 가져오기 */
	var or model.Order
	o.orderCollection.FindOne(o.ctx, filter).Decode(&or)

	/*
		코드가 많이 길어지는 경우, 관련 로직만을 모아 따로 함수로 분리하는 것을 추천드립니다. 가독성이 매우 높아지고, 테스트를 작성하기에도 쉬워집니다.
		메뉴 추가에 대한 함수, 메뉴 변경에 대한 함수로 분리할 수 있겠습니다.
	*/
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
		/*
			변경할 수 없는 경우를 하나의 메시지로 처리하는 것은 어떤가요?
			여러 케이스로 세분화 하는 것 보다는, 주문을 변경할 수 없는 상황이라면 하나의 메시지로 전달해도 무방해보이고, 코드도 깔끔해질 것 같습니다.
		*/
		if or.Status >= model.OrderCancel && or.Status <= model.Complete || or.MenuName == menuname {
			return -1, errors.New("주문을 변경할 수 없습니다")
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
func (o *OrdererServiceImplement) GetOrders() ([]model.Order, []model.Order, error) {

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
