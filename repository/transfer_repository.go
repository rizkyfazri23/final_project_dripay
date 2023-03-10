package repository

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/rizkyfazri23/dripay/model/entity"
)

type TransferRepository interface {
	CreateTransfer(newTransfer *entity.TransferInfo) (entity.Transfer, error)
}

type transferRepository struct {
	db *sql.DB
}

func (r *transferRepository) CreateTransfer(newTransfer *entity.TransferInfo) (entity.Transfer, error) {
	log.Println(newTransfer)

	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return entity.Transfer{}, err
	}

	var senderId int
	query := "SELECT member_id FROM m_member WHERE username = $1"
	err = tx.QueryRow(query, newTransfer.SenderUsername).Scan(&senderId)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return entity.Transfer{}, err
	} else {
		log.Println("Get senderId")
	}

	var senderBalance float32
	query = "SELECT wallet_amount FROM m_member WHERE member_id = $1"
	err = tx.QueryRow(query, senderId).Scan(&senderBalance)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return entity.Transfer{}, err
	} else {
		log.Println("Get SenderBalance")
	}
	if senderBalance < newTransfer.TransferAmount {
		log.Println("Insufficient funds")
		tx.Rollback()
		return entity.Transfer{}, err
	} else {
		log.Println("Duit cukup")
	}

	//Kirim Uang
	_, err = tx.Exec("UPDATE m_member SET wallet_amount = wallet_amount - $1 WHERE member_id = $2", newTransfer.TransferAmount, senderId)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return entity.Transfer{}, err
	} else {
		log.Println("funds transferred")
	}

	var receiptId string
	query = "SELECT member_id FROM m_member WHERE username = $1"
	err = tx.QueryRow(query, newTransfer.ReceiptUsername).Scan(&receiptId)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return entity.Transfer{}, err
	} else {
		log.Println("Get ReceiptId")
	}

	ReceiptId, _ := strconv.Atoi(receiptId)

	// Terima Uang
	_, err = tx.Exec("UPDATE m_member SET wallet_amount = wallet_amount + $1 WHERE member_id = $2", newTransfer.TransferAmount, ReceiptId)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return entity.Transfer{}, err
	} else {
		log.Println("funds received")
	}

	var gatewayId string
	query = "SELECT gateway_id FROM m_gateway WHERE gateway_name = $1"
	err = tx.QueryRow(query, newTransfer.PaymentGateway).Scan(&gatewayId)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return entity.Transfer{}, err
	} else {
		log.Println("Get ReceiptId")
	}
	GatewayId, _ := strconv.Atoi(gatewayId)

	var TransCode string
	var dateTime time.Time
	var transferID int

	query = "INSERT INTO t_transfer (sender_id, receipt_id, transfer_amount, transfer_gateway_id, description) VALUES ($1, $2, $3, $4, $5) RETURNING transfer_id, transfer_code, date_time"
	err = r.db.QueryRow(query, senderId, receiptId, newTransfer.TransferAmount, GatewayId, newTransfer.Description).Scan(&transferID, &TransCode, &dateTime)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return entity.Transfer{}, err
	} else {
		log.Println("Transfer Created")
	}

	query = "INSERT INTO t_transaction_log (member_id, type_id, amount, transaction_code)  VALUES ($1, $2, $3, $4)"
	_, err = r.db.Exec(query, senderId, 1, newTransfer.TransferAmount, TransCode)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return entity.Transfer{}, err
	} else {
		log.Println("Transaction Log Created")
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return entity.Transfer{}, err
	} else {
		log.Println("Commited")
	}

	return entity.Transfer{
		Id:                  transferID,
		Transfer_Code:       TransCode,
		Sender_Id:           senderId,
		Transfer_Amount:     newTransfer.TransferAmount,
		Transfer_Gateway_Id: GatewayId,
		Receipt_Id:          ReceiptId,
		Description:         newTransfer.Description,
		Date_time:           dateTime,
	}, nil
}

func NewTransferRepo(db *sql.DB) TransferRepository {
	repo := new(transferRepository)
	repo.db = db
	return repo
}
