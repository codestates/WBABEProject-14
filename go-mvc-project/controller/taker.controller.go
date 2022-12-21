package controller

import (
	"fmt"
	"net/http"
	"wba/go-mvc-procjet/model"
	"wba/go-mvc-procjet/services"

	"github.com/gin-gonic/gin"
)

type TakerController struct {
	TakerController services.TakerService
}

func TakerControllerNew(takerservice services.TakerService) TakerController {
	return TakerController{
		TakerController: takerservice,
	}
}

// 피주문자 api 작성
// 주문자 api 컨트롤러 작성
func (oc *TakerController) CreateMenu(ctx *gin.Context) {
	var menu model.Menu
	fmt.Println(ctx.Request)
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := oc.TakerController.CreateMenu(&menu)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
