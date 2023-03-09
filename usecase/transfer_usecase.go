package usecase

import (
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/repository"
)

type TransferUsecase interface {
	TransferHistory(id int) ([]*entity.Transfer, error)
	TransferBalance(newTransfer *entity.TransferInfo) (*entity.Transfer, error)
}

type transferUsecase struct {
	transferRepo repository.TransferRepository
}

func NewTransferUsecase(transRepo repository.TransferRepository) TransferUsecase {
	return &transferUsecase{
		transferRepo: transRepo,
	}
}

func (u *transferUsecase) TransferHistory(id int) ([]*entity.Transfer, error) {
	return u.transferRepo.FindTransferHistory(id)
}
func (u *transferUsecase) TransferBalance(newTransfer *entity.TransferInfo) (*entity.Transfer, error) {
	return u.transferRepo.CreateTransfer(newTransfer)
}
