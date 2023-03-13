package usecase

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/jung-kurt/gofpdf"
	"github.com/rizkyfazri23/dripay/model/entity"
	"github.com/rizkyfazri23/dripay/repository"
)

type HistoryUsecase interface {
	GetAll(memberID int) ([]entity.History, error)
	GetAllPayment(memberID int) ([]entity.History, error)
	GetAllTransfer(memberID int) ([]entity.History, error)
	GetAllDeposit(memberID int) ([]entity.History, error)
	ExportPDF(histories []entity.History) ([]byte, error)
}

type historyUsecase struct {
	historyRepo repository.HistoryRepository
}

func NewHistoryUsecase(historyRepo repository.HistoryRepository) HistoryUsecase {
	return &historyUsecase{
		historyRepo: historyRepo,
	}
}

func (u *historyUsecase) GetAll(memberID int) ([]entity.History, error) {
	return u.historyRepo.AllHistory(memberID)
}

func (u *historyUsecase) GetAllPayment(memberID int) ([]entity.History, error) {
	return u.historyRepo.PaymentHistory(memberID)
}

func (u *historyUsecase) GetAllTransfer(memberID int) ([]entity.History, error) {
	return u.historyRepo.TransferHistory(memberID)
}

func (u *historyUsecase) GetAllDeposit(memberID int) ([]entity.History, error) {
	return u.historyRepo.DepositHistory(memberID)
}

func (u *historyUsecase) ExportPDF(histories []entity.History) ([]byte, error) {
	var pdf = gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Transaction History")
	pdf.Ln(12)

	// Set table headers
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(10, 10, "No", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Username", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Type", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Debit", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Kredit", "1", 0, "C", false, 0, "")
	pdf.CellFormat(45, 10, "Date Time", "1", 0, "C", false, 0, "")
	pdf.CellFormat(50, 10, "Status", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Set table data
	pdf.SetFont("Arial", "", 12)
	no := 0
	for _, h := range histories {
		debitIDR := fmt.Sprintf("Rp %.2f", h.Debit)
		kreditIDR := fmt.Sprintf("Rp %.2f", h.Kredit)

		no++
		pdf.CellFormat(10, 10, strconv.Itoa(no), "1", 0, "R", false, 0, "")
		pdf.CellFormat(40, 10, h.Member_Username, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, h.Transaction_Type, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, debitIDR, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, kreditIDR, "1", 0, "", false, 0, "")
		pdf.CellFormat(45, 10, h.Date_Time.Format("2006-01-02 15:04:05"), "1", 0, "", false, 0, "")
		pdf.CellFormat(50, 10, h.Status, "1", 0, "", false, 0, "")
		pdf.Ln(-1)
	}

	buf := new(bytes.Buffer)
	err := pdf.Output(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
