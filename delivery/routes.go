package delivery

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rizkyfazri23/dripay/config"
	"github.com/rizkyfazri23/dripay/controller"
	"github.com/rizkyfazri23/dripay/manager"
)

type AppServer struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
	host           string
}

func (p *AppServer) v1() {
	v1Routes := p.engine.Group("/v1/dripay")

	gatewayRouterGroup := v1Routes.Group("/payment/gateway")
	p.gatewayController(gatewayRouterGroup)

	depositRouterGroup := v1Routes.Group("/deposit")
	p.depositController(depositRouterGroup)
}

func (p *AppServer) gatewayController(rg *gin.RouterGroup) {
	customerController := controller.NewGatewayController(rg, p.usecaseManager.GatewayUsecase())
	rg.GET("/", customerController.ReadGateway)
	rg.POST("/", customerController.CreateGateway)
	rg.PUT("/:id", customerController.UpdateGateway)
	rg.DELETE("/", customerController.DeleteGateway)
}

func (p *AppServer) depositController(rg *gin.RouterGroup) {
	depositController := controller.NewDepositController(rg, p.usecaseManager.DepositUsecase())
	rg.GET("/:member_id", depositController.GetAll)
	rg.GET("/:deposit_code", depositController.GetByCode)
	rg.GET("/:payment_gateway_id", depositController.GetByGateway)
	rg.POST("", depositController.Add)
}

func (p *AppServer) Run() {
	p.v1()
	err := p.engine.Run(p.host)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application failed to run", err)
		}
	}()
	if err != nil {
		log.Println(err)
	}
}

func Server() *AppServer {
	r := gin.Default()
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)
	host := fmt.Sprintf(":%s", c.ServerPort)
	return &AppServer{
		usecaseManager: usecaseManager,
		engine:         r,
		host:           host,
	}
}
