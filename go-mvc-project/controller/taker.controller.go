package controller

import (
	"net/http"
	"time"
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

/* 메뉴 등록 */
func (oc *TakerController) CreateMenu(ctx *gin.Context) {
	var menu model.Menu
	/* MENU 바인딩, ERROR 체크 */
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	menu.CreatedAt = time.Now()
	err := oc.TakerService.CreateMenu(&menu)
	/* 메뉴 추가 로직에 에러 검증 */
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	/* 성공 여부를 리턴 */
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
