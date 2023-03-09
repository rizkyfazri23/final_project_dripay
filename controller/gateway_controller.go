package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/usecase"
)

type GatewayController struct {
	router         *gin.RouterGroup
	gatewayUsecase usecase.GatewayUsecase
}

func NewGatewayController(r *gin.RouterGroup, u usecase.GatewayUsecase) *GatewayController {
	controller := GatewayController{
		router:         r,
		gatewayUsecase: u,
	}

	return &controller
}

func (c *GatewayController) CreateGateway(w *gin.Context) {
	var gateway entity.Gateway
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

func (c *GatewayController) ReadGateway(w *gin.Context) {
	gateways, err := c.gatewayUsecase.ReadGateway()
	if err != nil {
		w.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	w.JSON(http.StatusOK, gateways)
}

func (c *GatewayController) UpdateGateway(w *gin.Context) {
	var gateway entity.Gateway
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

func (c *GatewayController) DeleteGateway(w *gin.Context) {
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
