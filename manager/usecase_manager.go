package manager

import (
	"github.com/rizkyfazri23/dripay/usecase"
)

type UsecaseManager interface {
	GatewayUsecase() usecase.GatewayUsecase
	TransferUsecase() usecase.TransferUsecase
}

type usecaseManager struct {
	repoManager RepoManager
}

func (u *usecaseManager) GatewayUsecase() usecase.GatewayUsecase {
	return usecase.NewGatewayUsecase(u.repoManager.GatewayRepo())
}

func (u *usecaseManager) TransferUsecase() usecase.TransferUsecase {
	return usecase.NewTransferUsecase(u.repoManager.TransferRepo())
}

func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: rm,
	}
}
