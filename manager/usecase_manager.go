package manager

import (
	"github.com/rizkyfazri23/dripay/usecase"
)

type UsecaseManager interface {
	DepositUsecase() usecase.DepositUsecase
	GatewayUsecase() usecase.GatewayUsecase
}

type usecaseManager struct {
	repoManager RepoManager
}

func (u *usecaseManager) DepositUsecase() usecase.DepositUsecase {
	return usecase.NewDepositUsecase(u.repoManager.DepositRepo())
}

func (u *usecaseManager) GatewayUsecase() usecase.GatewayUsecase {
	return usecase.NewGatewayUsecase(u.repoManager.GatewayRepo())
}

func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: rm,
	}
}
