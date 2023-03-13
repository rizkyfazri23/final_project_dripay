package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/middlewares"
	"github.com/rizkyfazri23/dripay/model/app_error"
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/usecase"
	"github.com/rizkyfazri23/dripay/utils"
)

type SplitController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.SplitUsecase
}

func NewSplitController(r *gin.RouterGroup, u usecase.SplitUsecase) *SplitController {
	controller := SplitController{
		router:  r,
		usecase: u,
	}

	dGroup := r.Group("/payment/split")
	dGroup.Use(middlewares.JwtAuthMiddleware())

	dGroup.POST("/", controller.Add)

	return &controller
}

func (c *SplitController) Add(ctx *gin.Context) {
	var split *entity.SplitRequest

	if err := ctx.BindJSON(&split); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		return
	}

	member_id, err := utils.ExtractTokenID(ctx)

	res, err := c.usecase.Add(split, member_id)

	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(err.Error()))
		return
	}

	c.Success(ctx, http.StatusCreated, "01", "Successfully split bill", res)
}
