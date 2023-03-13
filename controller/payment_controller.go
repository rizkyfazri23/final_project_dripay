package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/middlewares"
	"github.com/rizkyfazri23/dripay/model/app_error"
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/usecase"
	"github.com/rizkyfazri23/dripay/utils"
)

type PaymentController struct {
	BaseController
	router  *gin.RouterGroup
	usecase usecase.PaymentUsecase
}

func NewPaymentController(r *gin.RouterGroup, u usecase.PaymentUsecase) *PaymentController {
	controller := PaymentController{
		router:  r,
		usecase: u,
	}
	pGroup := r.Group("/payment")
	pGroup.Use(middlewares.JwtAuthMiddleware())
	pGroup.GET("/", controller.GetAll)
	pGroup.GET("/:id", controller.GetOne)
	pGroup.POST("/", controller.Create)
	pGroup.PUT("/:id", controller.Update)

	return nil
}

func (c *PaymentController) Create(ctx *gin.Context) {
	var payment entity.PaymentRequest

	if err := ctx.BindJSON(&payment); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		return
	}

	member_id, err := utils.ExtractTokenID(ctx)
	if err != nil {
		c.Failed(ctx, http.StatusUnauthorized, "", fmt.Errorf("failed to extract token"))
		return
	}

	res, err := c.usecase.CreatePayment(&payment, member_id)
	fmt.Println(err)
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to create payment"))
		return
	}
	c.Success(ctx, http.StatusCreated, "01", "Successfully created new payment", res)
}

func (c *PaymentController) GetAll(ctx *gin.Context) {
	res, err := c.usecase.GetAllPayment()
	if err != nil {
		c.Failed(ctx, http.StatusInternalServerError, "", app_error.UnknownError(""))
		return
	}

	c.Success(ctx, http.StatusOK, "", "Successfully Get Payment data", res)
}

func (c *PaymentController) GetOne(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.UnknownError(""))
		return
	}
	c.usecase.GetPayment(paymentId)

}

func (c *PaymentController) Update(ctx *gin.Context) {
	var paymentStatus entity.PaymentRequestStatus
	paymentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", fmt.Errorf("payment id required"))
		return
	}

	if err := ctx.BindJSON(&paymentStatus); err != nil {
		c.Failed(ctx, http.StatusBadRequest, "", app_error.DataNotFoundError(""))
		return
	}
	member_id, err := utils.ExtractTokenID(ctx)
	if err != nil {
		c.Failed(ctx, http.StatusUnauthorized, "", fmt.Errorf("failed to extract token"))
		return
	}
	res, err := c.usecase.UpdatePayment(paymentId, member_id)
	if err != nil {
		fmt.Print(err)
		c.Failed(ctx, http.StatusInternalServerError, "", fmt.Errorf("failed to update payment"))
		return
	}
	c.Success(ctx, http.StatusCreated, "01", "Successfully update payment status", res)

}
