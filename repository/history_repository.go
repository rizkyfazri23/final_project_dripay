package repository

import (
	"database/sql"
	"log"

	"github.com/rizkyfazri23/dripay/model/entity"
)

type HistoryRepository interface {
	AllHistory(memberID int) ([]entity.History, error)
	PaymentHistory(memberID int) ([]entity.History, error)
	TransferHistory(memberID int) ([]entity.History, error)
	DepositHistory(memberID int) ([]entity.History, error)
}

type historyRepository struct {
	db *sql.DB
}

func (r *historyRepository) AllHistory(memberID int) ([]entity.History, error) {

	query := (`				SELECT 
										t.transaction_log_id as Id,
										m.username as Member_Username,
										ty.type_name as Transaction_Type,
										CASE 
											WHEN t.status = 0 THEN t.amount
											ELSE 0
										END AS debit,
										CASE 
											WHEN t.status = 1 THEN t.amount
											ELSE 0
										END AS kredit,
										t.date_time as Date_Time,
										t.status as Status,
										t.transaction_code as Transaction_Code
							FROM		 
										t_transaction_log t
							INNER JOIN 					
										m_transaction_type ty ON t.type_id = ty.type_id
							INNER JOIN 					
										m_member m ON t.member_id = m.member_id
							WHERE 		m.member_id = $1
							ORDER BY	date_time desc`)
	rows, err := r.db.Query(query, memberID)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	var histories []entity.History

	for rows.Next() {
		var h entity.History
		err := rows.Scan(&h.Id, &h.Member_Username, &h.Transaction_Type, &h.Debit, &h.Debit, &h.Date_Time, &h.Status, &h.Transaction_Code)
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

func (r *historyRepository) PaymentHistory(memberID int) ([]entity.History, error) {

	query := (`				SELECT 
										t.transaction_log_id as Id,
										m.username as Member_Username,
										ty.type_name as Transaction_Type,
										CASE 
											WHEN t.status = 0 THEN t.amount
											ELSE 0
										END AS debit,
										t.date_time as Date_Time,
										t.status as Status,
										t.transaction_code as Transaction_Code
							FROM		 
										t_transaction_log t
							INNER JOIN 					
										m_transaction_type ty ON t.type_id = ty.type_id
							INNER JOIN 					
										m_member m ON t.member_id = m.member_id
							WHERE 		m.member_id = $1 AND ty.type_name = 'Payment'
							ORDER BY	date_time desc`)
	rows, err := r.db.Query(query, memberID)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	var histories []entity.History

	for rows.Next() {
		var h entity.History
		err := rows.Scan(&h.Id, &h.Member_Username, &h.Transaction_Type, &h.Debit, &h.Date_Time, &h.Status, &h.Transaction_Code)
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

func (r *historyRepository) TransferHistory(memberID int) ([]entity.History, error) {

	query := (`				SELECT 
										t.transaction_log_id as Id,
										m.username as Member_Username,
										ty.type_name as Transaction_Type,
										CASE 
											WHEN t.status = 0 THEN t.amount
											ELSE 0
										END AS debit,
										CASE 
											WHEN t.status = 1 THEN t.amount
											ELSE 0
										END AS kredit,
										t.date_time as Date_Time,
										t.status as Status,
										t.transaction_code as Transaction_Code
							FROM		 
										t_transaction_log t
							INNER JOIN 					
										m_transaction_type ty ON t.type_id = ty.type_id
							INNER JOIN 					
										m_member m ON t.member_id = m.member_id
							WHERE 		m.member_id = $1 AND ty.type_name = 'Transfer'
							ORDER BY	date_time desc`)
	rows, err := r.db.Query(query, memberID)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	var histories []entity.History

	for rows.Next() {
		var h entity.History
		err := rows.Scan(&h.Id, &h.Member_Username, &h.Transaction_Type, &h.Debit, &h.Kredit, &h.Date_Time, &h.Status, &h.Transaction_Code)
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

func (r *historyRepository) DepositHistory(memberID int) ([]entity.History, error) {

	query := (`				SELECT 
										t.transaction_log_id as Id,
										m.username as Member_Username,
										ty.type_name as Transaction_Type,
										CASE 
											WHEN t.status = 1 THEN t.amount
											ELSE 0
										END AS kredit,
										t.date_time as Date_Time,
										t.status as Status,
										t.transaction_code as Transaction_Code
							FROM		 
										t_transaction_log t
							INNER JOIN 					
										m_transaction_type ty ON t.type_id = ty.type_id
							INNER JOIN 					
										m_member m ON t.member_id = m.member_id
							WHERE 		m.member_id = $1 AND ty.type_name = 'Deposit'
							ORDER BY	date_time desc`)
	rows, err := r.db.Query(query, memberID)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	var histories []entity.History

	for rows.Next() {
		var h entity.History
		err := rows.Scan(&h.Id, &h.Member_Username, &h.Transaction_Type, &h.Kredit, &h.Date_Time, &h.Status, &h.Transaction_Code)
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

func NewHistoryRepository(db *sql.DB) HistoryRepository {
	repo := new(historyRepository)
	repo.db = db
	return repo
}
