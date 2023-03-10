package usecase

import (
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/repository"
)

type HistoryUsecase interface {
	GetAll() ([]entity.History, error)
	GetAllPayment() ([]entity.History, error)
	GetAllTransfer() ([]entity.History, error)
	GetAllDeposit() ([]entity.History, error)
}

type historyUsecase struct {
	historyRepo repository.HistoryRepository
}

func NewHistoryUsecase(historyRepo repository.HistoryRepository) HistoryUsecase {
	return &historyUsecase{
		historyRepo: historyRepo,
	}
}

func (u *historyUsecase) GetAll() ([]entity.History, error) {
	return u.historyRepo.AllHistory()
}

func (u *historyUsecase) GetAllPayment() ([]entity.History, error) {
	return u.historyRepo.PaymentHistory()
}

func (u *historyUsecase) GetAllTransfer() ([]entity.History, error) {
	return u.historyRepo.TransferHistory()
}

func (u *historyUsecase) GetAllDeposit() ([]entity.History, error) {
	return u.historyRepo.TransferHistory()
}
