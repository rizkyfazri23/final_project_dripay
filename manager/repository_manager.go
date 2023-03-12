package manager

import "github.com/rizkyfazri23/dripay/repository"

type RepoManager interface {
	GatewayRepo() repository.GatewayRepo
	MemberRepo() repository.MemberRepo
	HistoryRepo() repository.HistoryRepository
	DepositRepo() repository.DepositRepository
	TransferRepo() repository.TransferRepository
	PaymentRepo() repository.PaymentRepository
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) GatewayRepo() repository.GatewayRepo {
	return repository.NewGatewayRepository(r.infraManager.DbConn())
}

func (r *repositoryManager) MemberRepo() repository.MemberRepo {
	return repository.NewMemberRepository(r.infraManager.DbConn())
}

func (r *repositoryManager) HistoryRepo() repository.HistoryRepository {
	return repository.NewHistoryRepository(r.infraManager.DbConn())
}

func (r *repositoryManager) DepositRepo() repository.DepositRepository {
	return repository.NewDepositRepository(r.infraManager.DbConn())
}

func (r *repositoryManager) TransferRepo() repository.TransferRepository {
	return repository.NewTransferRepo(r.infraManager.DbConn())
}

func (r *repositoryManager) PaymentRepo() repository.PaymentRepository {
	return repository.NewPaymentRepository(r.infraManager.DbConn())
}

func NewRepoManager(manager InfraManager) RepoManager {
	return &repositoryManager{
		infraManager: manager,
	}
}
