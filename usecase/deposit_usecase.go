package usecase

import (
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/repository"
)

type DepositUsecase interface {
	Add(newDeposit *entity.DepositRequest) (*entity.Deposit, error)
}

type depositUsecase struct {
	depositRepo repository.DepositRepository
}

func NewDepositUsecase(depositRepo repository.DepositRepository) DepositUsecase {
	return &depositUsecase{
		depositRepo: depositRepo,
	}
}

func (u *depositUsecase) Add(newDeposit *entity.DepositRequest) (*entity.Deposit, error) {
	return u.depositRepo.MakeDeposit(newDeposit)
}
