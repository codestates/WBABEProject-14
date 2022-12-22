package controller

import (
	"fmt"
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

// 주문자 api 컨트롤러 작성
func (oc *OrdererController) CreateOrder(ctx *gin.Context) {
	var order model.Order
	fmt.Println(ctx.Request)
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := oc.OrdererService.CreateOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
