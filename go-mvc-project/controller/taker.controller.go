package controller

import (
	"net/http"
	"wba/go-mvc-procjet/model"
	"wba/go-mvc-procjet/services"

	"github.com/gin-gonic/gin"
)

type TakerController struct {
	TakerService services.TakerService
}

func NewTakerController(takerservice services.TakerService) (TakerController, error) {
	return TakerController{
		TakerService: takerservice,
	}, nil
}

// @Summary 💡API 신규 메뉴 등록
// @Description 신규 메뉴를 생성합니다.
// @Param  request body model.ApiCreateMenuRequest true "메뉴 등록 요청 Body"
// @name CreateMenu
// @Accept   json
// @Produce  json
// @Router /api/v01/taker/menu [POST]
// @Success 201 {object} model.Success "메뉴 생성 성공"
// @Failure 400 {object} model.Failure "메뉴 생성 실패(동일 메뉴이름이 존재)"
func (oc *TakerController) CreateMenu(ctx *gin.Context) {
	var menu model.Menu
	/* BINDING */
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	/* CREATE */
	if err := oc.TakerService.CreateMenu(&menu); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	/* RESPONSE */
	ctx.JSON(http.StatusCreated, gin.H{"message": "success"})
}

// @Summary 💡API 기존 메뉴 수정
// @Description 기존 메뉴를 수정합니다.
// @Param menuname path string true "Menu Name"
// @Parma menuname
// @Param request body model.ApiUpdateMenuRequest true "메뉴 수정 요청 Body"
// @Accept  json
// @Produce json
// @Router /api/v01/taker/menu/{menuname} [PUT]
// @Success 200 {object} model.Success "메뉴 수정 성공"
// @Failure 400 {object} model.Failure "메뉴 수정 실패"
func (oc *TakerController) UpdateMenu(ctx *gin.Context) {
	var menu model.Menu
	menuname := ctx.Param("menuname")
	/* BINDING */
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	/* UPDATE */
	if err := oc.TakerService.UpdateMenu(menuname, &menu); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	/* RESPONSE */
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// @Summary 💡API 기존 메뉴 삭제
// @Description 기존 메뉴 삭제 플래그를 변경(false -> true)합니다.
// @name DeleteMenu
// @Param  request body model.ApiUpdateOrderRequest true "메뉴 삭제 요청 Body"
// @Accept  json
// @Produce  json
// @Router /api/v01/taker/menu [DELETE]
// @Success 200 {object} model.Success "메뉴 삭제 성공"
// @Failure 400 {object} model.Failure "메뉴 삭제 실패"
func (oc *TakerController) DeleteMenu(ctx *gin.Context) {
	var menu model.Menu
	/* BINDING */
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	/* DELETE */
	if err := oc.TakerService.DeleteMenu(&menu); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	/* RESPONSE */
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// @Summary 💡API 금일 추천 메뉴 변경
// @Description 추천 메뉴를 변경합니다. 요청한 메뉴가  금일 추천 메뉴라면 추천하지 않음으로, 추천 메뉴가 아니라면 추천 메뉴로 설정합니다. 응답으로 추천 메뉴들을 반환합니다.
// @name UpdateMenuRecommend
// @Param  request body model.ApiUpdateOrderRequest true "금일 추천 메뉴 변경 요청 Body"
// @Accept   json
// @Produce  json
// @Router /api/v01/taker/menu [PATCH]
// @Success 200 {object} model.ApiUpdateMenuRecommend "추천 메뉴 변경 성공"
// @Failure 400 {object} model.Failure "추천 메뉴 변경 실패"
func (oc *TakerController) UpdateMenuRecommend(ctx *gin.Context) {
	var menu *model.Menu
	/* BINDING */
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	/* UPDATE */
	menus, err := oc.TakerService.UpdateMenuRecommend(menu)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	/* RESPONSE */
	ctx.JSON(http.StatusOK, menus)
}

// @Summary 💡API 주문 리스트 조회
// @Description 주문 리스트를 조회합니다. 주문 상태가 접수되기전 상태인 주문들만 조회합니다 (주문이 막 들어온 상태).
// @name GetOrderList
// @Accept  json
// @Produce  json
// @Router /api/v01/taker/orders [GET]
// @Success 200 {object} model.ApiGetOrderListResponse "주문 조회 성공"
// @Failure 400 {object} model.Failure "주문 조회 실패"
func (oc *TakerController) GetOrderList(ctx *gin.Context) {
	/* RESPONSE */
	if orderList, err := oc.TakerService.GetOrderList(); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, orderList)
	}
}

// @Summary 💡API  각 메뉴별 주문 상태 변경
// @Description 각 메뉴별 주문들의 상태를 다음 단계로 변경합니다. 메뉴이름 과 현재 상태를 요청합니다. 해당 메뉴의 0(접수중)인 주문들을 1(접수완료) 상태로 업데이트 합니다.
// @name UpdateOrderStatus
// @Param  request body model.ApiUpdateOrderStatusRequest true "메뉴별 상태 업데이트 요청 Body"
// @Param  menuname path string true "Menu Name"
// @Accept  json
// @Produce  json
// @Router /api/v01/taker/orders/{menuname} [PATCH]
// @Success 200 {object} model.Success "메뉴별 주문 상태 저장 성공"
// @Failure 400 {object} model.Failure "저장 실패"
func (oc *TakerController) UpdateOrderStatus(ctx *gin.Context) {
	var order *model.Order
	orderId := ctx.Param("menuname")
	/* BINDING */
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	/* UPDATE */
	if err := oc.TakerService.UpdateOrderStatus(orderId, order.Status); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	/* RESPONSE */
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
