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

/* 메뉴 등록 */
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

/* 메뉴 수정 */
func (oc *TakerController) UpdateMenu(ctx *gin.Context) {
	var menu model.Menu
	/* BINDING */
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	/* UPDATE */
	if err := oc.TakerService.UpdateMenu(&menu); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	/* RESPONSE */
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

/* 메뉴 삭제 */
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

/* 금일 추천 메뉴 변경 */
func (oc *TakerController) UpdateMenuRecommend(ctx *gin.Context) {
	var menu model.Menu
	/* BINDING */
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	/* UPDATE */
	if err := oc.TakerService.UpdateMenuRecommend(&menu); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	/* RESPONSE */
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
