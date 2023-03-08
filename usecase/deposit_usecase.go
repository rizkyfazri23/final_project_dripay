package usecase

import (
	"github.com/rizkyfazri23/dripay/model"
	"github.com/rizkyfazri23/dripay/repository"
)

type DepositUsecase interface {
	GetAll(member_id int) ([]*model.Deposit, error)
	GetByCode(deposit_code int) (*model.Deposit, error)
	GetByGateway(payment_gateway_id int) (*model.Deposit, error)
	Add(newDeposit *model.Deposit) (*model.Deposit, error)
}

type depositUsecase struct {
	depositRepo repository.DepositRepository
}

func NewDepositUsecase(depositRepo repository.DepositRepository) DepositUsecase {
	return &depositUsecase{
		depositRepo: depositRepo,
	}
}

func (u *depositUsecase) GetAll(member_id int) ([]*model.Deposit, error) {
	return u.depositRepo.FindDepositHistory(member_id)
}

func (u *depositUsecase) GetByCode(deposit_code int) (*model.Deposit, error) {
	return u.depositRepo.FindDepositByCode(deposit_code)
}

func (u *depositUsecase) GetByGateway(payment_gateway_id int) (*model.Deposit, error) {
	return u.depositRepo.FindDepositByGateway(payment_gateway_id)
}

func (u *depositUsecase) Add(newDeposit *model.Deposit) (*model.Deposit, error) {
	return u.depositRepo.CreateDeposit(newDeposit)
}
