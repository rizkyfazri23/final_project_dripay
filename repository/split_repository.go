package repository

import (
	"database/sql"
	"errors"
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

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	var memberUsername string

	err = tx.QueryRow(`SELECT username FROM m_member WHERE member_id = $1`, member_id).Scan(&memberUsername)
	if err != nil {
		return nil, err
	}

	found := false
	for _, member := range newSplit.Member_List {
		if member.Member_Username == memberUsername {
			found = true
			break
		}
	}

	if !found {
		err := errors.New("user not in the member list")
		return nil, err
	}

	memberCount := len(newSplit.Member_List)

	paymentAmount := newSplit.Total_Amount / float32(memberCount)

	for _, member := range newSplit.Member_List {
		var memberID, paymentID int
		var paymentCode string
		var paymentTime time.Time

		err := tx.QueryRow("SELECT member_id FROM m_member WHERE member_username = $1", member.Member_Username).Scan(&memberID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		err = tx.QueryRow("INSERT INTO t_payment (member_id, payment_amount, payment_gateway_id, description, status) VALUES ($1, $2, $3, $4, $5) RETURNING payment_id, payment_code, date_time", memberID, paymentAmount, 1, newSplit.Description, "Success").Scan(&paymentID, paymentCode, paymentTime)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		_, err = tx.Exec("UPDATE m_member SET wallet_amount = wallet_amount - $1 WHERE member_id = $2", paymentAmount, memberID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		_, err = tx.Exec("INSERT INTO t_transaction_log (member_id, transaction_amount, transaction_type) VALUES ($1, $2, $3)", memberID, paymentAmount, "split bill")
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		response = append(response, entity.SplitResponse{
			Payment_Id:      paymentID,
			Payment_Code:    paymentCode,
			Member_Username: member.Member_Username,
			Payment_Amount:  paymentAmount,
			Payment_Gateway: "Dripay",
			Description:     newSplit.Description,
			Status:          "Success",
			Date_Time:       paymentTime,
		})
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewSplitRepository(db *sql.DB) SplitRepository {
	repo := new(splitRepository)
	repo.db = db
	return repo
}
