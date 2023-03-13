package usecase

import (
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/repository"
)

type GatewayUsecase interface {
	GetAll() ([]entity.Gateway, error)
	GetOne(id int) (entity.Gateway, error)
	Add(newGateway *entity.Gateway) (entity.Gateway, error)
	Edit(gateway *entity.Gateway) (entity.Gateway, error)
	Remove(id int) error
}

type gatewayUsecase struct {
	gatewayRepo repository.GatewayRepo
}

func (u *gatewayUsecase) GetAll() ([]entity.Gateway, error) {
	return u.gatewayRepo.FindAll()
}

func (u *gatewayUsecase) GetOne(id int) (entity.Gateway, error) {
	return u.gatewayRepo.FindOne(id)
}

func (u *gatewayUsecase) Add(newGateway *entity.Gateway) (entity.Gateway, error) {
	return u.gatewayRepo.Create(newGateway)
}

func (u *gatewayUsecase) Edit(gateway *entity.Gateway) (entity.Gateway, error) {
	return u.gatewayRepo.Update(gateway)
}

func (u *gatewayUsecase) Remove(id int) error {
	return u.gatewayRepo.Delete(id)
}

func NewGatewayUsecase(gatewayRepo repository.GatewayRepo) GatewayUsecase {
	return &gatewayUsecase{
		gatewayRepo: gatewayRepo,
	}
}
