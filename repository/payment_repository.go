package repository

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/rizkyfazri23/dripay/model/entity"
)

type PaymentRepository interface {
	CreatePayment(payment *entity.PaymentRequest, member_id int) (*entity.Payment, error)
	GetPayment(paymentId int) (*entity.Payment, error)
	GetAllPayment() ([]*entity.Payment, error)
	UpdatePayment(status string, paymentId int) (*entity.Payment, error)
}

type paymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (p *paymentRepository) randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var str string
	for i := 0; i < length; i++ {
		str += string(charset[rand.Intn(len(charset))])
	}
	return str
}

func (p *paymentRepository) CreatePayment(payment *entity.PaymentRequest, member_id int) (*entity.Payment, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return &entity.Payment{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		} else {
			err = tx.Commit()
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}()

	var paymentGatewayId, payment_id int
	var paymentTime time.Time
	var status string
	paymentCode := p.randomString(10)

	err = tx.QueryRow(`SELECT gateway_id FROM m_gateway WHERE gateway_name = $1`, payment.Payment_Gateway).Scan(&paymentGatewayId)
	if err != nil {
		return &entity.Payment{}, err
	}

	query := "INSERT INTO t_payment (payment_code, member_id, payment_amount, payment_gateway_id, description, status, date_time) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING payment_id, status"
	err = tx.QueryRow(query, paymentCode, member_id, payment.Payment_Amount, paymentGatewayId, payment.Description, "menunggu pembayaran", paymentTime).Scan(&payment_id, &status)
	if err != nil {
		return &entity.Payment{}, err
	}

	response := &entity.Payment{
		Id:                 payment_id,
		Payment_Code:       paymentCode,
		Member_Id:          member_id,
		Payment_Amount:     payment.Payment_Amount,
		Payment_Gateway_Id: paymentGatewayId,
		Description:        payment.Description,
		Status:             status,
		Date_Time:          paymentTime,
	}

	return response, nil
}

func (p *paymentRepository) GetPayment(paymentId int) (*entity.Payment, error) {
	var payment entity.Payment
	p.db.QueryRow(`SELECT * FROM t_payment WHERE payment_id = $1`).Scan(&payment)
	return &payment, nil
}

func (p *paymentRepository) GetAllPayment() ([]*entity.Payment, error) {
	var payments []*entity.Payment

	rows, err := p.db.Query(`SELECT * FROM t_payment`)
	if err != nil {
		return []*entity.Payment{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var payment entity.Payment
		if err := rows.Scan(&payment.Id, &payment.Payment_Code, &payment.Member_Id, &payment.Payment_Amount, &payment.Payment_Gateway_Id, &payment.Description, &payment.Status, &payment.Date_Time); err != nil {
			return []*entity.Payment{}, err
		}
		payments = append(payments, &payment)
	}
	if err := rows.Err(); err != nil {
		return []*entity.Payment{}, err
	}

	return payments, nil
}

func (p *paymentRepository) UpdatePayment(status string, paymentId int) (*entity.Payment, error) {
	_, err := p.db.Exec(`UPDATE t_payment SET status = $1 WHERE payment_id = $2`, status, paymentId)
	if err != nil {
		return &entity.Payment{}, err
	}

	payment, err := p.GetPayment(paymentId)
	if err != nil {
		return &entity.Payment{}, err
	}
	return payment, nil
}
