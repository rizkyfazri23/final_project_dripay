package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rizkyfazri23/dripay/controller"
	"github.com/rizkyfazri23/dripay/repository"
	"github.com/rizkyfazri23/dripay/usecase"
)

func GatewayRoutes(c *gin.Engine, db *sqlx.DB) {
	gatewayRepo := repository.NewGatewayRepo(db)
	gatewayUsecase := usecase.NewGatewayUsecase(gatewayRepo)
	gatewayController := controller.NewGatewayController(gatewayUsecase)

	r := c.Group("/payment/gateway")
	r.GET("/", gatewayController.ReadGateway)
	r.POST("/", gatewayController.CreateGateway)
	r.PUT("/:id", gatewayController.UpdateGateway)
	r.DELETE("/", gatewayController.DeleteGateway)

}
