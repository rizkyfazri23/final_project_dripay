package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/middlewares"
	"github.com/rizkyfazri23/dripay/model/app_error"
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/usecase"
	"github.com/rizkyfazri23/dripay/utils"
)

type TransferController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.TransferUsecase
}

func NewTransferController(r *gin.RouterGroup, u usecase.TransferUsecase) *TransferController {
	controller := TransferController{
		router:  r,
		usecase: u,
	}

	trGroup := r.Group("/transfer")
	trGroup.Use(middlewares.JwtAuthMiddleware())

	trGroup.POST("/", controller.AddTransfer)

	return &controller
}

func (c *TransferController) AddTransfer(ctx *gin.Context) {
	var newTransfer *entity.TransferInfo
	sender_Id, _ := utils.ExtractTokenID(ctx)
	if err := ctx.BindJSON(&newTransfer); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		return
	}

	if newTransfer.ReceiptUsername == "" {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("one or more required fields are missing"))
		return
	}
	if newTransfer.TransferAmount < 1 {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid amount"))
		return
	}

	res, err := c.usecase.TransferBalance(newTransfer, sender_Id)

	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to transfer fund"))
		return
	}

	c.Success(ctx, http.StatusCreated, "01", "Successfully transfer fund", res)
}
