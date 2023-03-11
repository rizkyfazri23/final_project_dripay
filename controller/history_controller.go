package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/middlewares"
	"github.com/rizkyfazri23/dripay/model/app_error"
	"github.com/rizkyfazri23/dripay/usecase"
)

type HistoryController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.HistoryUsecase
}

func NewHistoryController(r *gin.RouterGroup, u usecase.HistoryUsecase) *HistoryController {
	controller := HistoryController{
		router:  r,
		usecase: u,
	}

	hGroup := r.Group("/history")
	hGroup.Use(middlewares.JwtAuthMiddleware())

	hGroup.GET("/", controller.GetAll)
	hGroup.GET("/payment", controller.GetAllPayment)
	hGroup.GET("/deposit", controller.GetAllDeposit)
	hGroup.GET("/transfer", controller.GetAllTransfer)

	return &controller
}

func (c *HistoryController) GetAll(ctx *gin.Context) {
	res, err := c.usecase.GetAll()
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(""))
		return
	}

	c.Success(ctx, http.StatusOK, "", "Successfully retrieved all history data", res)
}

func (c *HistoryController) GetAllPayment(ctx *gin.Context) {
	res, err := c.usecase.GetAllPayment()
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(""))
		return
	}

	c.Success(ctx, http.StatusOK, "", "Successfully retrieved all payment data", res)
}

func (c *HistoryController) GetAllTransfer(ctx *gin.Context) {
	res, err := c.usecase.GetAllTransfer()
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(""))
		return
	}

	c.Success(ctx, http.StatusOK, "", "Successfully retrieved all transfer data", res)
}

func (c *HistoryController) GetAllDeposit(ctx *gin.Context) {
	res, err := c.usecase.GetAllDeposit()
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(""))
		return
	}

	c.Success(ctx, http.StatusOK, "", "Successfully retrieved all deposit data", res)
}
