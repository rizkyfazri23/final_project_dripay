package usecase

import (
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/repository"
)

type gatewayUsecase struct {
	gatewayRepo repository.GatewayRepo
}

type GatewayUsecase interface {
	CreateGateway(gateway *entity.Gateway) (*entity.Gateway, error)
	ReadGateway() ([]*entity.Gateway, error)
	UpdateGateway(gateway *entity.Gateway) (*entity.Gateway, error)
	DeleteGateway(id int) error
}

func NewGatewayUsecase(gatewayRepo repository.GatewayRepo) GatewayUsecase {
	return &gatewayUsecase{
		gatewayRepo: gatewayRepo,
	}
}

func (u *gatewayUsecase) CreateGateway(gateway *entity.Gateway) (*entity.Gateway, error) {
	return u.gatewayRepo.CreateGateway(gateway)
}

func (u *gatewayUsecase) ReadGateway() ([]*entity.Gateway, error) {
	return u.gatewayRepo.ReadGateway()
}

func (u *gatewayUsecase) UpdateGateway(gateway *entity.Gateway) (*entity.Gateway, error) {
	return u.gatewayRepo.UpdateGateway(gateway)
}

func (u *gatewayUsecase) DeleteGateway(id int) error {
	return u.gatewayRepo.DeleteGateway(id)
}
