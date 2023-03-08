package manager

import (
	"github.com/rizkyfazri23/dripay/repository"
)

type RepoManager interface {
	DepositRepo() repository.DepositRepository
	GatewayRepo() repository.GatewayRepo
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) DepositRepo() repository.DepositRepository {
	return repository.NewDepositRepository(r.infraManager.DbConn())
}

func (r *repositoryManager) GatewayRepo() repository.GatewayRepo {
	return repository.NewGatewayRepo(r.infraManager.DbConn())
}

func NewRepoManager(manager InfraManager) RepoManager {
	return &repositoryManager{
		infraManager: manager,
	}
}
