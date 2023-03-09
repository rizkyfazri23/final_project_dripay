package usecase

import (
	"github.com/rizkyfazri23/dripay/model"
	"github.com/rizkyfazri23/dripay/repository"
)

type gatewayUsecase struct {
	gatewayRepo repository.GatewayRepo
}

type GatewayUsecase interface {
	CreateGateway(gateway *model.Gateway) (*model.Gateway, error)
	ReadGateway() ([]*model.Gateway, error)
	UpdateGateway(gateway *model.Gateway) (*model.Gateway, error)
	DeleteGateway(id int) error
}

func NewGatewayUsecase(gatewayRepo repository.GatewayRepo) GatewayUsecase {
	return &gatewayUsecase{
		gatewayRepo: gatewayRepo,
	}
}

func (u *gatewayUsecase) CreateGateway(gateway *model.Gateway) (*model.Gateway, error) {
	return u.gatewayRepo.CreateGateway(gateway)
}

func (u *gatewayUsecase) ReadGateway() ([]*model.Gateway, error) {
	return u.gatewayRepo.ReadGateway()
}

func (u *gatewayUsecase) UpdateGateway(gateway *model.Gateway) (*model.Gateway, error) {
	return u.gatewayRepo.UpdateGateway(gateway)
}

func (u *gatewayUsecase) DeleteGateway(id int) error {
	return u.gatewayRepo.DeleteGateway(id)
}
