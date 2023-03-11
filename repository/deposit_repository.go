package repository

import (
	"database/sql"
	"time"

	"github.com/rizkyfazri23/dripay/model/entity"
)

type DepositRepository interface {
	MakeDeposit(newDeposit *entity.DepositRequest, member_id int) (entity.Deposit, error)
}

type depositRepository struct {
	db *sql.DB
}

// func (d *depositRepository) MakeDeposit(newDeposit *entity.DepositRequest) (entity.Deposit, error) {
// 	tx, err := d.db.Begin()
// 	if err != nil {
// 		log.Println(err)
// 		return entity.Deposit{}, err
// 	} else {
// 		log.Println("25")
// 	}

// 	var memberID int

// 	err = tx.QueryRow(`SELECT member_id FROM m_member WHERE username = $1`, newDeposit.Member_Username).Scan(&memberID)
// 	if err != nil {
// 		tx.Rollback()
// 		log.Println(err)
// 		return entity.Deposit{}, err
// 	} else {
// 		log.Println("36")
// 	}

// 	var depositGatewayID int

// 	err = tx.QueryRow(`SELECT gateway_id FROM m_gateway WHERE gateway_name = $1`, newDeposit.Deposit_Gateway).Scan(&depositGatewayID)
// 	if err != nil {
// 		tx.Rollback()
// 		log.Println(err)
// 		return entity.Deposit{}, err
// 	} else {
// 		log.Println("47")
// 	}

// 	var transactionCode string
// 	var dateTime time.Time

// 	query := `INSERT INTO t_deposit (member_id, deposit_amount, deposit_gateway_id, description) VALUES ($1, $2, $3, $4) RETURNING deposit_code, date_time`
// 	row := tx.QueryRow(query, memberID, newDeposit.Deposit_Amount, depositGatewayID, newDeposit.Description)

// 	err = row.Scan(&transactionCode, &dateTime)

// 	if err != nil {
// 		tx.Rollback()
// 		log.Println(err)
// 		return entity.Deposit{}, err
// 	} else {
// 		log.Println("62")
// 	}

// 	var transactionID int

// 	err = tx.QueryRow(`SELECT type_id FROM m_transaction_type WHERE type_name = $1`, "Deposit").Scan(&transactionID)
// 	if err != nil {
// 		tx.Rollback()
// 		log.Println(err)
// 		return entity.Deposit{}, err
// 	} else {
// 		log.Println("74")
// 	}

// 	query = `UPDATE m_member SET wallet_amount = wallet_amount + $1 WHERE username = $2`
// 	_ = tx.QueryRow(query, newDeposit.Deposit_Amount, newDeposit.Member_Username)

// 	if err != nil {
// 		tx.Rollback()
// 		log.Println(err)
// 		return entity.Deposit{}, err
// 	} else {
// 		log.Println("85")
// 	}

// 	query = `INSERT INTO t_transaction_log (member_id, type_id, amount, status, transaction_code) VALUES ($1, $2, $3, $4, $5)`
// 	_ = tx.QueryRow(query, memberID, transactionID, newDeposit.Deposit_Amount, 1, transactionCode)

// 	if err != nil {
// 		tx.Rollback()
// 		log.Println(err)
// 		return entity.Deposit{}, err
// 	} else {
// 		log.Println("96")
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		log.Println(err)
// 		tx.Rollback()
// 		return entity.Deposit{}, err
// 	} else {
// 		log.Println("103")
// 	}

// 	deposit := &entity.Deposit{
// 		Deposit_Code:       transactionCode,
// 		Member_Id:          memberID,
// 		Deposit_Amount:     newDeposit.Deposit_Amount,
// 		Deposit_Gateway_Id: depositGatewayID,
// 		Description:        newDeposit.Description,
// 		Date_Time:          dateTime,
// 	}
// 	return *deposit, nil
// }

func (d *depositRepository) MakeDeposit(newDeposit *entity.DepositRequest, member_id int) (entity.Deposit, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return entity.Deposit{}, err
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

	var depositGatewayID, transactionID, depositID int
	var transactionCode string
	var dateTime time.Time

	err = tx.QueryRow(`SELECT gateway_id FROM m_gateway WHERE gateway_name = $1`, newDeposit.Deposit_Gateway).Scan(&depositGatewayID)
	if err != nil {
		return entity.Deposit{}, err
	}

	query := `INSERT INTO t_deposit (member_id, deposit_amount, deposit_gateway_id, description) VALUES ($1, $2, $3, $4) RETURNING deposit_id, deposit_code, date_time`
	err = tx.QueryRow(query, member_id, newDeposit.Deposit_Amount, depositGatewayID, newDeposit.Description).Scan(&depositID, &transactionCode, &dateTime)
	if err != nil {
		return entity.Deposit{}, err
	}

	err = tx.QueryRow(`SELECT type_id FROM m_transaction_type WHERE type_name = $1`, "Deposit").Scan(&transactionID)
	if err != nil {
		return entity.Deposit{}, err
	}

	query = `UPDATE m_member SET wallet_amount = wallet_amount + $1 WHERE member_id = $2`
	_, err = tx.Exec(query, newDeposit.Deposit_Amount, member_id)
	if err != nil {
		return entity.Deposit{}, err
	}

	query = `INSERT INTO t_transaction_log (member_id, type_id, amount, status, transaction_code) VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.Exec(query, member_id, transactionID, newDeposit.Deposit_Amount, 1, transactionCode)
	if err != nil {
		return entity.Deposit{}, err
	}

	deposit := &entity.Deposit{
		Id:                 depositID,
		Deposit_Code:       transactionCode,
		Member_Id:          member_id,
		Deposit_Amount:     newDeposit.Deposit_Amount,
		Deposit_Gateway_Id: depositGatewayID,
		Description:        newDeposit.Description,
		Date_Time:          dateTime,
	}

	return *deposit, nil
}

func NewDepositRepository(db *sql.DB) DepositRepository {
	repo := new(depositRepository)
	repo.db = db
	return repo
}
