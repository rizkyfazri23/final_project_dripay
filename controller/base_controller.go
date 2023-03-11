/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Sat Mar 04 2023 10:08:36 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package controller

import (
	"github.com/rizkyfazri23/dripay/model/dto/res"
	"github.com/gin-gonic/gin"
)

type BaseController struct {}

func (b *BaseController) Success(c *gin.Context, httpCode int, code string, msg string, data any) {
	res.NewSuccessJsonResponse(c, httpCode, code, msg, data).Send()
}

func (b *BaseController) Failed(c *gin.Context, httpCode int, code string, err error) {
	res.NewErrorJsonResponse(c, httpCode, code, err).Send()
}