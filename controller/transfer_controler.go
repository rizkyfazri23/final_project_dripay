package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/model"
	"github.com/rizkyfazri23/dripay/usecase"
)

type TransferController struct {
	router  *gin.RouterGroup
	usecase usecase.TransferUsecase
}

func NewTransferController(r *gin.RouterGroup, u usecase.TransferUsecase) *TransferController {
	controller := TransferController{
		router:  r,
		usecase: u,
	}

	return &controller
}

func (c *TransferController) GetTransfer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Input",
		})
		return
	}
	res, err := c.usecase.TransferHistory(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *TransferController) AddTransfer(ctx *gin.Context) {
	var newTransfer *model.TransferInfo

	if err := ctx.BindJSON(&newTransfer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Input",
		})
	}
	res, err := c.usecase.TransferBalance(newTransfer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
