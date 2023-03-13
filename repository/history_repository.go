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

	query := (`	
			SELECT 			id,
							member_username, 
							transaction_type, 
							debit, 
							kredit,
							date_time,
							status, 
							transaction_code 
			FROM 			transaction_history
			WHERE			member_id = $1
			ORDER BY		date_time desc;`)
	rows, err := r.db.Query(query, memberID)
	if err != nil {
		log.Fatalln("masuk")

		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	var histories []entity.History

	for rows.Next() {
		var h entity.History
		err := rows.Scan(&h.Id, &h.Member_Username, &h.Transaction_Type, &h.Debit, &h.Kredit, &h.Date_Time, &h.Status, &h.Transaction_Code)
		if err != nil {
			log.Fatalln("masuk2")

			log.Fatal(err)
			return nil, err
		}
		histories = append(histories, h)
	}
	if err := rows.Err(); err != nil {
		log.Fatalln("masuk3")

		log.Fatal(err)
		return nil, err
	}

	return histories, nil
}

func (r *historyRepository) PaymentHistory(memberID int) ([]entity.History, error) {

	query := (`	
			SELECT 			id,
							member_username, 
							transaction_type, 
							debit, 
							date_time,
							status, 
							transaction_code 
			FROM 			transaction_history
			WHERE			member_id = $1 AND transaction_type = 'Payment'
			ORDER BY		date_time desc;`)
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

	query := (`	
			SELECT 			id,
							member_username, 
							transaction_type, 
							debit, 
							kredit,
							date_time,
							status, 
							transaction_code 
			FROM 			transaction_history
			WHERE			member_id = $1 AND transaction_type = 'Transfer'
			ORDER BY		date_time desc;`)
	rows, err := r.db.Query(query, memberID)

	if err != nil {
		log.Fatalln("masuk")

		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	var histories []entity.History

	for rows.Next() {
		var h entity.History
		err := rows.Scan(&h.Id, &h.Member_Username, &h.Transaction_Type, &h.Debit, &h.Kredit, &h.Date_Time, &h.Status, &h.Transaction_Code)
		if err != nil {
			log.Fatalln("masuk2")

			log.Fatal(err)
			return nil, err
		}
		histories = append(histories, h)
	}
	if err := rows.Err(); err != nil {
		log.Fatalln("masuk3")

		log.Fatal(err)
		return nil, err
	}

	return histories, nil
}

func (r *historyRepository) DepositHistory(memberID int) ([]entity.History, error) {

	query := (`	
			SELECT 			id,
							member_username, 
							transaction_type, 
							kredit,
							date_time,
							status, 
							transaction_code 
			FROM 			transaction_history
			WHERE			member_id = $1 AND transaction_type = 'Deposit'
			ORDER BY		date_time desc;`)
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
