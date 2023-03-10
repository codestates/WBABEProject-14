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

// @Summary ๐กAPI ์ ๊ท ๋ฉ๋ด ๋ฑ๋ก
// @Description ์ ๊ท ๋ฉ๋ด๋ฅผ ์์ฑํฉ๋๋ค.
// @Param  request body model.ApiCreateMenuRequest true "๋ฉ๋ด ๋ฑ๋ก ์์ฒญ Body"
// @name CreateMenu
// @Accept   json
// @Produce  json
// @Router /api/v01/taker/menu [POST]
// @Success 201 {object} model.Success "๋ฉ๋ด ์์ฑ ์ฑ๊ณต"
// @Failure 400 {object} model.Failure "๋ฉ๋ด ์์ฑ ์คํจ(๋์ผ ๋ฉ๋ด์ด๋ฆ์ด ์กด์ฌ)"
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

// @Summary ๐กAPI ๊ธฐ์กด ๋ฉ๋ด ์์ 
// @Description ๊ธฐ์กด ๋ฉ๋ด๋ฅผ ์์ ํฉ๋๋ค.
// @Param menuname path string true "Menu Name"
// @Parma menuname
// @Param request body model.ApiUpdateMenuRequest true "๋ฉ๋ด ์์  ์์ฒญ Body"
// @Accept  json
// @Produce json
// @Router /api/v01/taker/menu/{menuname} [PUT]
// @Success 200 {object} model.Success "๋ฉ๋ด ์์  ์ฑ๊ณต"
// @Failure 400 {object} model.Failure "๋ฉ๋ด ์์  ์คํจ"
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

// @Summary ๐กAPI ๊ธฐ์กด ๋ฉ๋ด ์ญ์ 
// @Description ๊ธฐ์กด ๋ฉ๋ด ์ญ์  ํ๋๊ทธ๋ฅผ ๋ณ๊ฒฝ(false -> true)ํฉ๋๋ค.
// @name DeleteMenu
// @Param  request body model.ApiUpdateOrderRequest true "๋ฉ๋ด ์ญ์  ์์ฒญ Body"
// @Accept  json
// @Produce  json
// @Router /api/v01/taker/menu [DELETE]
// @Success 200 {object} model.Success "๋ฉ๋ด ์ญ์  ์ฑ๊ณต"
// @Failure 400 {object} model.Failure "๋ฉ๋ด ์ญ์  ์คํจ"
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

// @Summary ๐กAPI ๊ธ์ผ ์ถ์ฒ ๋ฉ๋ด ๋ณ๊ฒฝ
// @Description ์ถ์ฒ ๋ฉ๋ด๋ฅผ ๋ณ๊ฒฝํฉ๋๋ค. ์์ฒญํ ๋ฉ๋ด๊ฐ  ๊ธ์ผ ์ถ์ฒ ๋ฉ๋ด๋ผ๋ฉด ์ถ์ฒํ์ง ์์์ผ๋ก, ์ถ์ฒ ๋ฉ๋ด๊ฐ ์๋๋ผ๋ฉด ์ถ์ฒ ๋ฉ๋ด๋ก ์ค์ ํฉ๋๋ค. ์๋ต์ผ๋ก ์ถ์ฒ ๋ฉ๋ด๋ค์ ๋ฐํํฉ๋๋ค.
// @name UpdateMenuRecommend
// @Param  request body model.ApiUpdateOrderRequest true "๊ธ์ผ ์ถ์ฒ ๋ฉ๋ด ๋ณ๊ฒฝ ์์ฒญ Body"
// @Accept   json
// @Produce  json
// @Router /api/v01/taker/menu [PATCH]
// @Success 200 {object} model.ApiUpdateMenuRecommend "์ถ์ฒ ๋ฉ๋ด ๋ณ๊ฒฝ ์ฑ๊ณต"
// @Failure 400 {object} model.Failure "์ถ์ฒ ๋ฉ๋ด ๋ณ๊ฒฝ ์คํจ"
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

// @Summary ๐กAPI ์ฃผ๋ฌธ ๋ฆฌ์คํธ ์กฐํ
// @Description ์ฃผ๋ฌธ ๋ฆฌ์คํธ๋ฅผ ์กฐํํฉ๋๋ค. ์ฃผ๋ฌธ ์ํ๊ฐ ์ ์๋๊ธฐ์  ์ํ์ธ ์ฃผ๋ฌธ๋ค๋ง ์กฐํํฉ๋๋ค (์ฃผ๋ฌธ์ด ๋ง ๋ค์ด์จ ์ํ).
// @name GetOrderList
// @Accept  json
// @Produce  json
// @Router /api/v01/taker/orders [GET]
// @Success 200 {object} model.ApiGetOrderListResponse "์ฃผ๋ฌธ ์กฐํ ์ฑ๊ณต"
// @Failure 400 {object} model.Failure "์ฃผ๋ฌธ ์กฐํ ์คํจ"
func (oc *TakerController) GetOrderList(ctx *gin.Context) {
	/* RESPONSE */
	if orderList, err := oc.TakerService.GetOrderList(); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, orderList)
	}
}

// @Summary ๐กAPI  ๊ฐ ๋ฉ๋ด๋ณ ์ฃผ๋ฌธ ์ํ ๋ณ๊ฒฝ
// @Description ๊ฐ ๋ฉ๋ด๋ณ ์ฃผ๋ฌธ๋ค์ ์ํ๋ฅผ ๋ค์ ๋จ๊ณ๋ก ๋ณ๊ฒฝํฉ๋๋ค. ๋ฉ๋ด์ด๋ฆ ๊ณผ ํ์ฌ ์ํ๋ฅผ ์์ฒญํฉ๋๋ค. ํด๋น ๋ฉ๋ด์ 0(์ ์์ค)์ธ ์ฃผ๋ฌธ๋ค์ 1(์ ์์๋ฃ) ์ํ๋ก ์๋ฐ์ดํธ ํฉ๋๋ค.
// @name UpdateOrderStatus
// @Param  request body model.ApiUpdateOrderStatusRequest true "๋ฉ๋ด๋ณ ์ํ ์๋ฐ์ดํธ ์์ฒญ Body"
// @Param  menuname path string true "Menu Name"
// @Accept  json
// @Produce  json
// @Router /api/v01/taker/orders/{menuname} [PATCH]
// @Success 200 {object} model.Success "๋ฉ๋ด๋ณ ์ฃผ๋ฌธ ์ํ ์ ์ฅ ์ฑ๊ณต"
// @Failure 400 {object} model.Failure "์ ์ฅ ์คํจ"
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
