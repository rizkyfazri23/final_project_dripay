package usecase

import (
	"github.com/rizkyfazri23/dripay/model"
	"github.com/rizkyfazri23/dripay/repository"
)

type TransferUsecase interface {
	TransferHistory(id int) ([]*model.Transfer, error)
	TransferBalance(newTransfer *model.TransferInfo) (*model.Transfer, error)
}

type transferUsecase struct {
	transferRepo repository.TransferRepository
}

func NewTransferUsecase(transRepo repository.TransferRepository) TransferUsecase {
	return &transferUsecase{
		transferRepo: transRepo,
	}
}

func (u *transferUsecase) TransferHistory(id int) ([]*model.Transfer, error) {
	return u.transferRepo.FindTransferHistory(id)
}
func (u *transferUsecase) TransferBalance(newTransfer *model.TransferInfo) (*model.Transfer, error) {
	return u.transferRepo.CreateTransfer(newTransfer)
}
