package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/model"
	"github.com/rizkyfazri23/dripay/usecase"
)

type GatewayController interface {
	CreateGateway(w *gin.Context)
	ReadGateway(w *gin.Context)
	UpdateGateway(w *gin.Context)
	DeleteGateway(w *gin.Context)
}

type gatewayController struct {
	gatewayUsecase usecase.GatewayUsecase
}

func NewGatewayController(gatewayUsecase usecase.GatewayUsecase) GatewayController {
	return &gatewayController{
		gatewayUsecase: gatewayUsecase,
	}
}

func (c *gatewayController) CreateGateway(w *gin.Context) {
	var gateway model.Gateway
	err := w.BindJSON(&gateway)
	if err != nil {
		w.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err = c.gatewayUsecase.CreateGateway(&gateway)
	if err != nil {
		w.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	w.JSON(http.StatusCreated, gateway)
}

func (c *gatewayController) ReadGateway(w *gin.Context) {
	gateways, err := c.gatewayUsecase.ReadGateway()
	if err != nil {
		w.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	w.JSON(http.StatusOK, gateways)
}

func (c *gatewayController) UpdateGateway(w *gin.Context) {
	var gateway model.Gateway
	err := w.BindJSON(&gateway)
	if err != nil {
		w.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err = c.gatewayUsecase.UpdateGateway(&gateway)
	if err != nil {
		w.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	w.JSON(http.StatusOK, gin.H{"message": "Update Success"})
}

func (c *gatewayController) DeleteGateway(w *gin.Context) {
	id, err := strconv.Atoi(w.Param("id"))
	if err != nil {
		w.JSON(http.StatusBadRequest, gin.H{"message": "ID must be Number"})
		return
	}
	err = c.gatewayUsecase.DeleteGateway(id)
	if err != nil {
		w.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	w.JSON(http.StatusOK, gin.H{"message": "Gateway Delete"})
}
