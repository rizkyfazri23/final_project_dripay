package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/rizkyfazri23/dripay/model/entity"
)

type SplitRepository interface {
	SplitBill(newSplit *entity.SplitRequest, member_id int) ([]entity.SplitResponse, error)
}

type splitRepository struct {
	db *sql.DB
}

func (s *splitRepository) SplitBill(newSplit *entity.SplitRequest, member_id int) ([]entity.SplitResponse, error) {
	var response []entity.SplitResponse
	var memberUsername string

	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	memberCount := len(newSplit.Member_List) + 1
	paymentAmount := newSplit.Total_Amount / float32(memberCount)

	err = tx.QueryRow(`SELECT username FROM m_member WHERE member_id = $1`, member_id).Scan(&memberUsername)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	newMember := entity.MemberList{
		Member_Username: memberUsername,
	}

	newSplit.Member_List = append(newSplit.Member_List, newMember)

	for _, member := range newSplit.Member_List {
		var memberID, paymentID int
		var paymentCode string
		var paymentTime time.Time
		var walletAmount float32

		err := tx.QueryRow("SELECT member_id, wallet_amount FROM m_member WHERE username = $1", member.Member_Username).Scan(&memberID, &walletAmount)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return []entity.SplitResponse{}, err
		}

		if walletAmount < paymentAmount {
			err := fmt.Errorf("insufficient funds for member %s", member.Member_Username)
			tx.Rollback()
			return []entity.SplitResponse{}, err
		}

		err = tx.QueryRow("INSERT INTO t_payment (member_id, payment_amount, payment_gateway_id, description, status) VALUES ($1, $2, $3, $4, $5) RETURNING payment_id, payment_code, date_time", memberID, paymentAmount, 1, newSplit.Description, "Success").Scan(&paymentID, &paymentCode, &paymentTime)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return []entity.SplitResponse{}, err
		}
		_, err = tx.Exec("UPDATE m_member SET wallet_amount = wallet_amount - $1 WHERE member_id = $2", paymentAmount, memberID)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return []entity.SplitResponse{}, err
		}

		_, err = tx.Exec("INSERT INTO t_transaction_log (member_id, amount, type_id, transaction_code) VALUES ($1, $2, $3, $4)", memberID, paymentAmount, 1, paymentCode)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return []entity.SplitResponse{}, err
		}

		response = append(response, entity.SplitResponse{
			Payment_Code:    paymentCode,
			Member_Username: member.Member_Username,
			Payment_Amount:  paymentAmount,
			Payment_Gateway: "Dripay",
			Description:     newSplit.Description,
			Status:          "Success",
		})
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return []entity.SplitResponse{}, err
	}

	return response, nil
}

func NewSplitRepository(db *sql.DB) SplitRepository {
	repo := new(splitRepository)
	repo.db = db
	return repo
}
