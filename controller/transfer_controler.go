package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/model/app_error"
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/usecase"
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

	r.GET("/transfer", controller.AddTransfer)

	return &controller
}

func (c *TransferController) AddTransfer(ctx *gin.Context) {
	var newTransfer *entity.TransferInfo

	if err := ctx.BindJSON(&newTransfer); err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(""))
	}
	res, err := c.usecase.TransferBalance(newTransfer)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to add deposit"))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
