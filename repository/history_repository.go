package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/rizkyfazri23/dripay/model/entity"
)

type HistoryRepository interface {
	AllHistory() ([]*entity.History, error)
	PaymentHistory() ([]*entity.History, error)
	TransferHistory() ([]*entity.History, error)
	DepositHistory() ([]*entity.History, error)
}

type historyRepository struct {
	db *sqlx.DB
}

func (r *historyRepository) AllHistory() ([]*entity.History, error) {

	rows, err := r.db.Query(`SELECT 
							t.transaction_log_id as Id,
							m.username as Member_Username,
							ty.type_name as Transaction_Type,
							t.amount as Amount,
							t.date_time as Date_Time,
							t.status as Status,
							t.transaction_code as Transaction_Code
							FROM 
							t_transaction_log t
							INNER JOIN m_transaction_type ty ON t.type_id = ty.type_id
							INNER JOIN m_member m ON t.member_id = m.member_id`)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	var histories []*entity.History

	for rows.Next() {
		var h *entity.History
		err := rows.Scan(&h.Id, &h.Member_Username, &h.Transaction_Type, &h.Amount, &h.Date_Time, &h.Status, &h.Transaction_Code)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		histories = append(histories, h)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return histories, nil
}

func (r *historyRepository) PaymentHistory() ([]*entity.History, error) {

	rows, err := r.db.Query(`SELECT 
							t.transaction_log_id as Id,
							m.username as Member_Username,
							ty.type_name as Transaction_Type,
							t.amount as Amount,
							t.date_time as Date_Time,
							t.status as Status,
							t.transaction_code as Transaction_Code
							FROM 
							t_transaction_log t
							INNER JOIN m_transaction_type ty ON t.type_id = ty.type_id
							INNER JOIN m_member m ON t.member_id = m.member_id
							WHERE ty.type_name = 'Payment'`)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	var histories []*entity.History

	for rows.Next() {
		var h *entity.History
		err := rows.Scan(&h.Id, &h.Member_Username, &h.Transaction_Type, &h.Amount, &h.Date_Time, &h.Status, &h.Transaction_Code)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		histories = append(histories, h)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return histories, nil
}

func (r *historyRepository) TransferHistory() ([]*entity.History, error) {

	rows, err := r.db.Query(`SELECT 
							t.transaction_log_id as Id,
							m.username as Member_Username,
							ty.type_name as Transaction_Type,
							t.amount as Amount,
							t.date_time as Date_Time,
							t.status as Status,
							t.transaction_code as Transaction_Code
							FROM 
							t_transaction_log t
							INNER JOIN m_transaction_type ty ON t.type_id = ty.type_id
							INNER JOIN m_member m ON t.member_id = m.member_id
							WHERE ty.type_name = 'Transfer'`)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	var histories []*entity.History

	for rows.Next() {
		var h *entity.History
		err := rows.Scan(&h.Id, &h.Member_Username, &h.Transaction_Type, &h.Amount, &h.Date_Time, &h.Status, &h.Transaction_Code)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		histories = append(histories, h)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return histories, nil
}

func (r *historyRepository) DepositHistory() ([]*entity.History, error) {

	rows, err := r.db.Query(`SELECT 
							t.transaction_log_id as Id,
							m.username as Member_Username,
							ty.type_name as Transaction_Type,
							t.amount as Amount,
							t.date_time as Date_Time,
							t.status as Status,
							t.transaction_code as Transaction_Code
							FROM 
							t_transaction_log t
							INNER JOIN m_transaction_type ty ON t.type_id = ty.type_id
							INNER JOIN m_member m ON t.member_id = m.member_id
							WHERE ty.type_name = 'Deposit'`)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	var histories []*entity.History

	for rows.Next() {
		var h *entity.History
		err := rows.Scan(&h.Id, &h.Member_Username, &h.Transaction_Type, &h.Amount, &h.Date_Time, &h.Status, &h.Transaction_Code)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		histories = append(histories, h)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return histories, nil
}

func NewHistoryRepository(db *sqlx.DB) HistoryRepository {
	repo := new(historyRepository)
	repo.db = db
	return repo
}
