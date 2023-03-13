package usecase

import (
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/repository"
)

type PaymentTypeUsecase interface {
	AddType(newType *entity.TransactionTypeInput) (entity.TransactionType, error)
	GetAll() ([]entity.TransactionType, error)
	GetOne(typeID int) (entity.TransactionType, error)
	Edit(typeID int, typeEdit *entity.TransactionTypeInput) (entity.TransactionType, error)
	Remove(typeID int) error
}

type paymentTypeUsecase struct {
	typeRepo repository.PaymentTypeRepo
}

func (u *paymentTypeUsecase) AddType(newType *entity.TransactionTypeInput) (entity.TransactionType, error) {
	return u.typeRepo.CreateType(newType)
}
func (u *paymentTypeUsecase) GetAll() ([]entity.TransactionType, error) {
	return u.typeRepo.ReadAllType()
}
func (u *paymentTypeUsecase) GetOne(typeID int) (entity.TransactionType, error) {
	return u.typeRepo.ReadTypeById(typeID)
}
func (u *paymentTypeUsecase) Edit(typeID int, typeEdit *entity.TransactionTypeInput) (entity.TransactionType, error) {
	return u.typeRepo.UpdateType(typeID, typeEdit)
}
func (u *paymentTypeUsecase) Remove(typeID int) error {
	return u.typeRepo.DeleteType(typeID)
}

func newPaymentTypeUsecase(typesRepo repository.PaymentTypeRepo) PaymentTypeUsecase {
	return &paymentTypeUsecase{
		typeRepo: typesRepo,
	}
}
