package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/model"
	"github.com/rizkyfazri23/dripay/usecase"
)

type DepositController struct {
	router  *gin.RouterGroup
	usecase usecase.DepositUsecase
}

func NewDepositController(r *gin.RouterGroup, u usecase.DepositUsecase) *DepositController {
	controller := DepositController{
		router:  r,
		usecase: u,
	}

	return &controller
}

func (c *DepositController) GetAll(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("member_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res, err := c.usecase.GetAll(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *DepositController) GetByCode(ctx *gin.Context) {
	code, err := strconv.Atoi(ctx.Param("deposit_code"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res, err := c.usecase.GetByCode(code)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *DepositController) GetByGateway(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("payment_gateway_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res, err := c.usecase.GetByGateway(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *DepositController) Add(ctx *gin.Context) {
	var deposit *model.Deposit

	if err := ctx.BindJSON(&deposit); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res, err := c.usecase.Add(deposit)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
