package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/model/app_error"
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/usecase"
)

type TransactionTypeController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.PaymentTypeUsecase
}

func NewTransactionTypeController(r *gin.RouterGroup, u usecase.PaymentTypeUsecase) *TransactionTypeController {
	controller := TransactionTypeController{
		router:  r,
		usecase: u,
	}

	TrTypGroup := r.Group("/trType")
	TrTypGroup.POST("/", controller.AddType)
	TrTypGroup.GET("/", controller.GetAll)
	TrTypGroup.GET("/:id", controller.GetOne)
	TrTypGroup.PUT("/edit/:id", controller.Edit)
	TrTypGroup.DELETE("/delete/:id", controller.Remove)
	return &controller
}

func (c *TransactionTypeController) AddType(ctx *gin.Context) {
	var newType entity.TransactionTypeInput

	if err := ctx.BindJSON(&newType); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError("wrong input format"))
		return
	}
	if newType.TypeName == "" {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("one or more required fields are missing"))
		return
	}
	res, err := c.usecase.AddType(&newType)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to create transaction type"))
		return
	}

	c.Success(ctx, http.StatusCreated, "01", "Successfully created new transaction type", res)
}

func (c *TransactionTypeController) GetAll(ctx *gin.Context) {
	res, err := c.usecase.GetAll()
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(""))
		return
	}

	c.Success(ctx, http.StatusOK, "", "Successfully retrieved all transaction type", res)
}

func (c *TransactionTypeController) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid id"))
		return
	}
	res, err := c.usecase.GetOne(id)
	if err != nil {
		c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("type with id %d not found", id)))
		return
	}

	c.Success(ctx, http.StatusOK, "", fmt.Sprintf("type with %d found", id), res)
}

func (c *TransactionTypeController) Edit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid id"))
		return
	}
	var typeEdit entity.TransactionTypeInput

	if err := ctx.BindJSON(&typeEdit); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.InvalidError("invalid request body"))
		return
	}

	res, err := c.usecase.Edit(id, &typeEdit)
	if err != nil {
		c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("failed to make change")))
		return
	}

	c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Successfully updated transaction type with Id %d", id), res)
}

func (c *TransactionTypeController) Remove(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "X01", app_error.InvalidError("invalid id"))
		return
	}
	err = c.usecase.Remove(id)
	if err != nil {
		c.Failed(ctx, http.StatusNotFound, "X04", app_error.DataNotFoundError(fmt.Sprintf("member with id %d not found", id)))
		return
	}

	c.Success(ctx, http.StatusOK, "", fmt.Sprintf("Transaction type with id %d has been removed", id), nil)
}
