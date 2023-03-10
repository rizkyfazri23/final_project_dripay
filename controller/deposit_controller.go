package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/middlewares"
	"github.com/rizkyfazri23/dripay/model/app_error"
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/usecase"
)

type DepositController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.DepositUsecase
}

func NewDepositController(r *gin.RouterGroup, u usecase.DepositUsecase) *DepositController {
	controller := DepositController{
		router:  r,
		usecase: u,
	}

	dGroup := r.Group("/deposit")
	dGroup.Use(middlewares.JwtAuthMiddleware())

	dGroup.POST("/", controller.Add)

	return &controller
}

func (c *DepositController) Add(ctx *gin.Context) {
	var deposit *entity.DepositRequest

	if err := ctx.BindJSON(&deposit); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		return
	}

	if deposit.Member_Username == "" {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("one or more required fields are missing"))
		return
	}

	res, err := c.usecase.Add(deposit)

	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to add deposit"))
		return
	}

	c.Success(ctx, http.StatusCreated, "01", "Successfully created new deposit", res)
}
