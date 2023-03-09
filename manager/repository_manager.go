package manager

import (
	"github.com/rizkyfazri23/dripay/repository"
)

type RepoManager interface {
	GatewayRepo() repository.GatewayRepo
	TransferRepo() repository.TransferRepository
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) TransferRepo() repository.TransferRepository {
	return repository.NewTransferRepo(r.infraManager.DbConn())
}

func (r *repositoryManager) GatewayRepo() repository.GatewayRepo {
	return repository.NewGatewayRepo(r.infraManager.DbConn())
}

func NewRepoManager(manager InfraManager) RepoManager {
	return &repositoryManager{
		infraManager: manager,
	}
}
