package usecase

import (
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/repository"
)

type HistoryUsecase interface {
	GetAll(memberID int) ([]entity.History, error)
	GetAllPayment(memberID int) ([]entity.History, error)
	GetAllTransfer(memberID int) ([]entity.History, error)
	GetAllDeposit(memberID int) ([]entity.History, error)
}

type historyUsecase struct {
	historyRepo repository.HistoryRepository
}

func NewHistoryUsecase(historyRepo repository.HistoryRepository) HistoryUsecase {
	return &historyUsecase{
		historyRepo: historyRepo,
	}
}

func (u *historyUsecase) GetAll(memberID int) ([]entity.History, error) {
	return u.historyRepo.AllHistory(memberID)
}

func (u *historyUsecase) GetAllPayment(memberID int) ([]entity.History, error) {
	return u.historyRepo.PaymentHistory(memberID)
}

func (u *historyUsecase) GetAllTransfer(memberID int) ([]entity.History, error) {
	return u.historyRepo.TransferHistory(memberID)
}

func (u *historyUsecase) GetAllDeposit(memberID int) ([]entity.History, error) {
	return u.historyRepo.DepositHistory(memberID)
}
