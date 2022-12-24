package controller

import (
	"net/http"
	"strconv"
	"wba/go-mvc-procjet/model"
	"wba/go-mvc-procjet/services"

	"github.com/gin-gonic/gin"
)

type OrdererController struct {
	OrdererService services.OrdererService
}

func NewOrdererController(ordererservice services.OrdererService) (OrdererController, error) {
	return OrdererController{
		OrdererService: ordererservice,
	}, nil
}

// @Summary 💡API 주문 생성 요청
// @Description 새로운 주문을 생성합니다. json 포맷으로 요청합니다. 반환값은 주문번호 입니다.
// @name CreateOrder
// @Param  request body model.ApiMenuCreateRequest true "주문 생성 요청"
// @Accept  json
// @Produce  json
// @Router /api/v01/orderer/order [POST]
// @Success 201 {object} model.ApiMenuCreateResponse "주문 생성 성공"
// @Failure 400 {object} model.Failure "주문 생성 실패"
func (oc *OrdererController) CreateOrder(ctx *gin.Context) {
	var order model.Order
	/* BINDING */
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Failure{Message: err.Error()})
		return
	}
	/* CREATE */
	orderNumber, err := oc.OrdererService.CreateOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, model.Failure{Message: err.Error()})
		return
	}
	/* RETURN */
	ctx.JSON(http.StatusCreated, gin.H{"주문번호": orderNumber})
}

// @Summary 💡API 메뉴 리스트 조회
// @Description 메뉴 리스트를 조회합니다. Path 로 정렬할 데이터를 받습니다. [recommend, grade, reorder, createdat]
// @name GetAllMenu
// @Param sort path string true "Sorting Parameter"
// @Accept  json
// @Produce  json
// @Router /api/v01/orderer/menu/{sort} [GET]
// @Success 200 {object} model.ApiMenuListResponse "메뉴 리스트 조회 성공"
// @Failure 400 {object} model.Failure "조회 실패"
func (oc *OrdererController) GetAllMenu(ctx *gin.Context) {
	menulist, err := oc.OrdererService.GetAllMenu(ctx.Param("sort"))
	/* GET */
	if err != nil {
		ctx.JSON(http.StatusBadGateway, model.Failure{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menulist)
}

// @Summary 💡API 리뷰 작성 요청
// @Description 완료된 주문에 리뷰를 작성합니다. 해당 주문의 ID를 받습니다. 성공/실패 여부를 반환합니다.
// @name CreateReview
// @Param  orderID path string true "Review ID"
// @Param  request body model.ApiCreateReviewRequest true "리뷰 생성 요청 Body"
// @Accept  json
// @Produce  json
// @Router /api/v01/orderer/review/{orderID} [POST]
// @Success 201 {object} model.Success "리뷰 작성 성공"
// @Failure 400 {object} model.Failure "리뷰 작성 실패"
func (oc *OrdererController) CreateReview(ctx *gin.Context) {
	var review model.Review
	id := ctx.Param("orderID")
	/* BINDING */
	if err := ctx.ShouldBindJSON(&review); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Failure{Message: err.Error()})
		return
	}
	/* CREATE */
	err := oc.OrdererService.CreateReview(&review, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Failure{Message: err.Error()})
		return
	}
	/* RETURN */
	ctx.JSON(http.StatusOK, model.Success{Message: "success"})
}

// @Summary 💡API  메뉴 상세 조회
// @Description 각 메뉴별 상세 정보를 조회합니다. 메뉴 이름을 받습니다. 해당 메뉴의 평점과 리뷰 리스트를 반환합니다.
// @name GetMenuDetail
// @Param  menuname path string true "Menu Name"
// @Accept  json
// @Produce  json
// @Router /api/v01/orderer/detailMenu/{menuname} [GET]
// @Success 200 {object} model.ApiGetMenuDetailResponse "메뉴 상세 조회 성공"
// @Failure 400 {object} model.Failure "메뉴 상세 조회 실패"
func (oc *OrdererController) GetMenuDetail(ctx *gin.Context) {
	avgGrade, reviewlist, err := oc.OrdererService.GetAllReiview(ctx.Param("menuname"))
	/* GET */
	if err != nil {
		ctx.JSON(http.StatusBadGateway, model.Failure{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"avgGrade": avgGrade, "reviewlist": reviewlist})
}

// @Summary 💡API 주문 변경 ( 메뉴 변경 및 추가 )
// @Description 주문한 메뉴를 변경 또는 추가 합니다. 주문ID, 추가 또는 변경의 Flag,  메뉴를 Body 로 받습니다. flag 가 0 일땐 메뉴 추가이며 배달중일 경우 신규주문으로 처리, 배달중이 아닐경우 성공을 반환합니다. flag 가 1일땐 메뉴 변경이며 조리중,배달중 일경우 실패, 아닐경우 성공을 반환합니다.
// @name UpdateOrder
// @Param  orderId path string true "Order Id"
// @Param  flag path string true "AddMenu : 0, UpdateMenu : 1"
// @Param  request body model.ApiUpdateOrderRequest true "변경 또는 추가할 메뉴"
// @Accept  json
// @Produce  json
// @Router /api/v01/orderer/order/{orderId}/{flag} [PATCH]
// @Success 201 {object} model.ApiUpdateOrderResponse "메뉴 추가 실패시 신규주문처리"
// @Success 200 {object} model.Success "메뉴 추가 성공시"
// @Failure 400 {object} model.Failure "메뉴 변경 실패시"
func (oc *OrdererController) UpdateOrder(ctx *gin.Context) {
	var order *model.Order
	id := ctx.Param("orderId")
	status, err := strconv.Atoi(ctx.Param("flag"))
	/* BINDING */
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Failure{Message: err.Error()})
		return
	}
	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Failure{Message: err.Error()})
		return
	}
	/* UPDATE */
	orderNumber, updateError := oc.OrdererService.UpdateOrder(id, status, order.MenuName)
	if updateError != nil {
		ctx.JSON(http.StatusBadRequest, model.Failure{Message: updateError.Error()})
		return
	}

	/* RESPONSE */
	if orderNumber == -1 {
		/* 메뉴 추가 성공 */
		ctx.JSON(http.StatusCreated, model.Success{Message: "success"})
		return
	}
	/* 메뉴 추가 실패 , 신규 주문으로 처리, 신규 주문번호 반환 */
	ctx.JSON(http.StatusCreated, gin.H{"message": "해당 주문은 배달중입니다. 신규주문으로 처리되었습니다.", "주문 번호": orderNumber})
}

// @Summary 💡API 주문 내역 조회
// @Description 주문 내역을 조회 합니다. 현재 주문과 주문이 완료된 이전 주문들을 반환합니다.
// @name GetOrders
// @Accept  json
// @Produce  json
// @Router /api/v01/orderer/orders [GET]
// @Success 200 {object} model.ApiGetOrdersResponse "주문 내역 조회 성공"
// @Failure 400 {object} model.Failure "주문 내역 조회 실패"
func (oc *OrdererController) GetOrders(ctx *gin.Context) {
	currents, pastorders, err := oc.OrdererService.GetOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Failure{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"currentOrders": currents, "pastOrders": pastorders})

}
