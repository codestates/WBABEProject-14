package controller

import (
	"net/http"
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

/* 주문 생성 */
func (oc *OrdererController) CreateOrder(ctx *gin.Context) {
	var order model.Order
	/* BINDING */
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	/* CREATE */
	err := oc.OrdererService.CreateOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	/* RETURN */
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

/* 메뉴 리스트 조회 */
func (oc *OrdererController) GetAllOrder(ctx *gin.Context) {
	menulist, err := oc.OrdererService.GetAllOrder(ctx.Param("sort"))
	/* GET */
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menulist)
}

/* 리뷰 작성 */
func (oc *OrdererController) CreateReview(ctx *gin.Context) {
	var review model.Review
	id := ctx.Param("objectId")
	/* BINDING */
	if err := ctx.ShouldBindJSON(&review); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	/* CREATE */
	err := oc.OrdererService.CreateReview(&review, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	/* RETURN */
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

/* 메뉴별 평점 및 리뷰 조회 */
func (oc *OrdererController) GetAllReiview(ctx *gin.Context) {
	avgGrade, reviewlist, err := oc.OrdererService.GetAllReiview(ctx.Param("menuname"))
	/* GET */
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"avgGrade": avgGrade, "reviewlist": reviewlist})
}
