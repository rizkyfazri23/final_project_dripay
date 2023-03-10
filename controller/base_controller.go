package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/model/dto/res"
)

type BaseController struct{}

func (b *BaseController) Success(c *gin.Context, httpCode int, code string, msg string, data any) {
	res.NewSuccessJsonResponse(c, httpCode, code, msg, data).Send()
}

func (b *BaseController) Failed(c *gin.Context, httpCode int, code string, err error) {
	res.NewErrorJsonResponse(c, httpCode, code, err).Send()
}
