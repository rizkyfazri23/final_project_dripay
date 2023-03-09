package repository

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rizkyfazri23/dripay/model"
)

type TransferRepository interface {
	FindTransferHistory(id int) ([]*model.Transfer, error)
	CreateTransfer(newTransfer *model.TransferInfo) (*model.Transfer, error)
}

type transferRepository struct {
	db *sqlx.DB
}

func (r *transferRepository) FindTransferHistory(id int) ([]*model.Transfer, error) {
	var transferList []*model.Transfer

	query := "SELECT * FROM t_transfer WHERE sender_id = $1"

	err := r.db.Select(&transferList, query, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return transferList, nil
}

func (r *transferRepository) CreateTransfer(newTransfer *model.TransferInfo) (*model.Transfer, error) {
	log.Println(newTransfer)
	tx := r.db.MustBegin()

	var senderId int
	query := "SELECT member_id FROM m_member WHERE username = $1"
	err := tx.Get(&senderId, query, newTransfer.SenderUsername)
	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
	} else {
		log.Println("Get senderId")
	}

	var senderBalance float32
	query = "SELECT wallet_amount FROM m_member WHERE member_id = $1"
	err = tx.Get(&senderBalance, query, senderId)
	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
	} else {
		log.Println("Get SenderBalance")
	}
	if senderBalance < newTransfer.TransferAmount {
		log.Println("Insufficient funds")
		tx.Rollback()
	} else {
		log.Println("Duit cukup")
	}

	//Kirim Uang
	_, err = tx.NamedExec("UPDATE m_member SET wallet_amount = wallet_amount - :Transfer_Amount WHERE member_id = :sender_id", map[string]interface{}{
		"Transfer_Amount": newTransfer.TransferAmount,
		"sender_id":       senderId,
	})
	if err != nil {
		log.Println(err)
		tx.Rollback()
	} else {
		log.Println("funds transferred")
	}

	var receiptId string
	query = "SELECT member_id FROM m_member WHERE username = $1"
	err = tx.Get(&receiptId, query, newTransfer.ReceiptUsername)
	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
	} else {
		log.Println("Get ReceiptId")
	}

	ReceiptId, err := strconv.Atoi(receiptId)

	// Terima Uang
	_, err = tx.NamedExec("UPDATE m_member SET wallet_amount = wallet_amount + :Transfer_Amount WHERE member_id = :receipt_id", map[string]interface{}{
		"Transfer_Amount": newTransfer.TransferAmount,
		"receipt_id":      ReceiptId,
	})
	if err != nil {
		log.Println(err)
		tx.Rollback()
	} else {
		log.Println("funds received")
	}

	var gatewayId string
	query = "SELECT gateway_id FROM m_gateway WHERE gateway_name = $1"
	err = tx.Get(&gatewayId, query, newTransfer.PaymentGateway)
	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
	} else {
		log.Println("Get ReceiptId")
	}
	GatewayId, err := strconv.Atoi(gatewayId)

	TransCode := rand.Intn(100000000)
	query = "INSERT INTO t_transfer (transfer_code, sender_id, receipt_id, transfer_amount, transfer_gateway_id, description) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = r.db.Exec(query, TransCode, senderId, receiptId, newTransfer.TransferAmount, GatewayId, newTransfer.Description)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	} else {
		log.Println("Transfer Created")
	}

	query = "INSERT INTO t_transaction_log (member_id, type_id, amount, transaction_code)  VALUES ($1, $2, $3, $4)"
	_, err = r.db.Exec(query, senderId, 1, newTransfer.TransferAmount, TransCode)
	if err != nil {
		log.Println(err)
		tx.Rollback()
	} else {
		log.Println("Transaction Log Created")
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		tx.Rollback()
	} else {
		log.Println("Commited")
	}

	return &model.Transfer{
		Sender_Id:           senderId,
		Transfer_Amount:     newTransfer.TransferAmount,
		Transfer_Gateway_Id: 2,
		Receipt_Id:          ReceiptId,
		Description:         newTransfer.Description,
	}, nil
}

func NewTransferRepo(db *sqlx.DB) TransferRepository {
	repo := new(transferRepository)
	repo.db = db
	return repo
}
