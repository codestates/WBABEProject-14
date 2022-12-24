package controller

import (
	"net/http"
	"wba/go-mvc-procjet/model"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}

func (p *Controller) RespOK(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}
