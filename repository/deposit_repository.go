package repository

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rizkyfazri23/dripay/model/entity"
)

type DepositRepository interface {
	MakeDeposit(newDeposit *entity.DepositRequest) (*entity.Deposit, error)
}

type depositRepository struct {
	db *sqlx.DB
}

func (d *depositRepository) MakeDeposit(newDeposit *entity.DepositRequest) (*entity.Deposit, error) {
	tx, err := d.db.Beginx()
	if err != nil {
		return &entity.Deposit{}, err
	}

	var memberID int

	err = tx.Get(&memberID, `SELECT member_id FROM m_member WHERE username = $1`, newDeposit.Member_Username)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return &entity.Deposit{}, err
	}

	var depositGatewayID int

	err = tx.Get(&depositGatewayID, `SELECT gateway_id FROM m_gateway WHERE gateway_name = $1`, newDeposit.Deposit_Gateway)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return &entity.Deposit{}, err
	}

	var transactionCode string
	var dateTime time.Time

	query := `INSERT INTO t_deposit (member_id, deposit_amount, deposit_gateway_id, description) VALUES ($1, $2, $3, $4) RETURNING deposit_code, date_time`
	row := tx.QueryRow(query, memberID, newDeposit.Deposit_Amount, depositGatewayID, newDeposit.Description)

	err = row.Scan(&transactionCode, &dateTime)

	if err != nil {
		tx.Rollback()
		log.Println(err)
		return &entity.Deposit{}, err
	}

	var transactionID int

	err = tx.Get(&transactionID, `SELECT type_id FROM m_transaction_type WHERE type_name = $1`, "Deposit")
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return &entity.Deposit{}, err
	}

	query = `INSERT INTO t_transaction_log (member_id, type_id, amount, status, transaction_code) VALUES ($1, $2, $3, $4, $5)`
	row = tx.QueryRow(query, memberID, transactionID, newDeposit.Deposit_Amount, "SUCCESS", transactionCode)

	if err != nil {
		tx.Rollback()
		log.Println(err)
		return &entity.Deposit{}, err
	}

	query = `UPDATE m_member SET wallet_amount = wallet_amount + $1 WHERE username = $2`
	_, err = tx.Exec(query, newDeposit.Deposit_Amount, newDeposit.Member_Username)

	if err != nil {
		tx.Rollback()
		return &entity.Deposit{}, err
	}

	err = tx.Commit()
	if err != nil {
		return &entity.Deposit{}, err
	}

	deposit := &entity.Deposit{
		Deposit_Code:       transactionCode,
		Member_Id:          memberID,
		Deposit_Amount:     newDeposit.Deposit_Amount,
		Deposit_Gateway_Id: depositGatewayID,
		Description:        newDeposit.Description,
		Date_time:          dateTime,
	}
	return deposit, nil
}

func NewDepositRepository(db *sqlx.DB) DepositRepository {
	repo := new(depositRepository)
	repo.db = db
	return repo
}
