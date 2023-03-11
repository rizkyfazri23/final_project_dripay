package usecase

import (
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/repository"
)

type MemberUsecase interface {
	GetAll() ([]entity.Member, error)
	GetOne(id int) (entity.Member, error)
	Add(newMember *entity.Member) (entity.Member, error)
	Edit(member *entity.Member) (entity.Member, error)
	Remove(id int) error
	LoginCheck(username string, password string) (string, error)
}

type memberUsecase struct {
	memberRepo repository.MemberRepo
}

func (u *memberUsecase) GetAll() ([]entity.Member, error) {
	return u.memberRepo.FindAll()
}

func (u *memberUsecase) GetOne(id int) (entity.Member, error) {
	return u.memberRepo.FindOne(id)
}

func (u *memberUsecase) Add(newMember *entity.Member) (entity.Member, error) {
	return u.memberRepo.Create(newMember)
}

func (u *memberUsecase) Edit(member *entity.Member) (entity.Member, error) {
	return u.memberRepo.Update(member)
}

func (u *memberUsecase) Remove(id int) error {
	return u.memberRepo.Delete(id)
}

func (u *memberUsecase) LoginCheck(username string, password string) (string, error) {
	return u.memberRepo.LoginCheck(username, password)
}

func NewMemberUsecase(memberRepo repository.MemberRepo) MemberUsecase {
	return &memberUsecase{
		memberRepo: memberRepo,
	}
}
