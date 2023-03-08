package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/rizkyfazri23/dripay/model"
)

type DepositRepository interface {
	FindDepositHistory(member_id int) ([]*model.Deposit, error)
	FindDepositByCode(deposit_code int) (*model.Deposit, error)
	FindDepositByGateway(payment_gateway_id int) (*model.Deposit, error)
	CreateDeposit(newDeposit *model.Deposit) (*model.Deposit, error)
}

type depositRepository struct {
	db *sqlx.DB
}

func NewDepositRepository(db *sqlx.DB) DepositRepository {
	repo := new(depositRepository)
	repo.db = db
	return repo
}

func (r *depositRepository) FindDepositHistory(member_id int) ([]*model.Deposit, error) {
	var deposits []*model.Deposit

	query := `SELECT * FROM t_deposit WHERE member_id = $1`

	err := r.db.Select(&deposits, query, member_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return deposits, nil
}

func (r *depositRepository) FindDepositByCode(deposit_code int) (*model.Deposit, error) {
	var deposit *model.Deposit

	query := `SELECT * FROM t_deposit WHERE deposit_code = $2 LIMIT 1`

	err := r.db.Select(&deposit, query, deposit_code)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return deposit, nil
}

func (r *depositRepository) FindDepositByGateway(payment_gateway_id int) (*model.Deposit, error) {
	var deposit *model.Deposit

	query := `SELECT * FROM t_deposit WHERE payment_gateway_id = $2 LIMIT 1`

	err := r.db.Select(&deposit, query, payment_gateway_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return deposit, nil
}

func (r *depositRepository) CreateDeposit(newDeposit *model.Deposit) (*model.Deposit, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := `
		INSERT INTO t_deposit (
			deposit_code, member_id, deposit_amount, payment_gateway_id, description, date_time
		) VALUES (
			$1, $2, $3, $4, $5, $6
		) RETURNING *
	`

	var deposit *model.Deposit
	err = tx.Get(&deposit, query,
		newDeposit.Deposit_Code,
		newDeposit.Member_Id,
		newDeposit.Deposit_Amount,
		newDeposit.Payment_Gateway_Id,
		newDeposit.Description,
		newDeposit.Date_time,
	)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return deposit, nil
}
