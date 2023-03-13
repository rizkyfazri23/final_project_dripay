package usecase

import (
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/repository"
)

type PaymentUsecase interface {
	CreatePayment(payment *entity.PaymentRequest, member_id int) (*entity.Payment, error)
	GetPayment(paymentId int) (*entity.Payment, error)
	GetAllPayment() ([]*entity.Payment, error)
	UpdatePayment(status string, paymentId int) (*entity.Payment, error)
}

type paymentUsecase struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
	}
}

func (u *paymentUsecase) CreatePayment(payment *entity.PaymentRequest, member_id int) (*entity.Payment, error) {
	return u.paymentRepo.CreatePayment(payment, member_id)
}

func (u *paymentUsecase) GetPayment(paymentId int) (*entity.Payment, error) {
	return u.paymentRepo.GetPayment(paymentId)
}

func (u *paymentUsecase) GetAllPayment() ([]*entity.Payment, error) {
	return u.paymentRepo.GetAllPayment()
}

func (u *paymentUsecase) UpdatePayment(status string, paymentId int) (*entity.Payment, error) {
	return u.paymentRepo.UpdatePayment(status, paymentId)
}
