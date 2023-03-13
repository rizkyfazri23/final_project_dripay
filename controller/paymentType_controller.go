package controller

import (
	"github.com/gin-gonic/gin"
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
	TrTypGroup.PUT("/", controller.Edit)
	TrTypGroup.DELETE("/", controller.Remove)
	return &controller
}
