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
	v1Routes := p.engine.Group("/v1")
	p.gatewayController(v1Routes)
	p.memberController(v1Routes)
	p.transferController(v1Routes)
	p.depositController(v1Routes)
	p.historyController(v1Routes)
	p.paymentController(v1Routes)
}

func (p *AppServer) gatewayController(rg *gin.RouterGroup) {
	controller.NewGatewayController(rg, p.usecaseManager.GatewayUsecase())
}

func (p *AppServer) memberController(rg *gin.RouterGroup) {
	controller.NewMemberController(rg, p.usecaseManager.MemberUsecase())
}

func (p *AppServer) transferController(rg *gin.RouterGroup) {
	controller.NewTransferController(rg, p.usecaseManager.TransferUsecase())
}

func (p *AppServer) depositController(rg *gin.RouterGroup) {
	controller.NewDepositController(rg, p.usecaseManager.DepositUsecase())
}

func (p *AppServer) historyController(rg *gin.RouterGroup) {
	controller.NewHistoryController(rg, p.usecaseManager.HistoryUsecase())
}

func (p *AppServer) paymentController(rg *gin.RouterGroup) {
	controller.NewPaymentController(rg, p.usecaseManager.PaymentUsecase())
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
		panic(err)
	}
}

func Server() *AppServer {
	r := gin.Default()
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepoManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)
	host := fmt.Sprintf(":%s", c.ApiPort)
	return &AppServer{
		usecaseManager: usecaseManager,
		engine:         r,
		host:           host,
	}
}
