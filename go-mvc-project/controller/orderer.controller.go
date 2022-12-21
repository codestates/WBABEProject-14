package controller

import (
	"fmt"
	"net/http"
	"wba/go-mvc-procjet/model"
	"wba/go-mvc-procjet/services"

	"github.com/gin-gonic/gin"
)

type OrdererController struct {
	OrdererController services.OrdererService
}

func OrdererControllerNew(ordererservice services.OrdererService) OrdererController {
	return OrdererController{
		OrdererController: ordererservice,
	}
}

// 주문자 api 컨트롤러 작성
func (oc *OrdererController) CreateOrder(ctx *gin.Context) {
	var order model.Order
	fmt.Println(ctx.Request)
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := oc.OrdererController.CreateOrder(&order)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
