package usecase

import (
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/repository"
)

type SplitUsecase interface {
	Add(newSplit *entity.SplitRequest, member_id int) ([]entity.SplitResponse, error)
}

type splitUsecase struct {
	splitRepo repository.SplitRepository
}

func NewSplitUsecase(splitRepo repository.SplitRepository) SplitUsecase {
	return &splitUsecase{
		splitRepo: splitRepo,
	}
}

func (u *splitUsecase) Add(newSplit *entity.SplitRequest, member_id int) ([]entity.SplitResponse, error) {
	return u.splitRepo.SplitBill(newSplit, member_id)
}
