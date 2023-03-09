package entity

type Transfer struct {
	Id                  int     `json:"transfer_id"`
	TransCode           int     `json:"transfer_code"`
	Sender_Id           int     `json:"sender_id"`
	Transfer_Amount     float32 `json:"transfer_amount"`
	Transfer_Gateway_Id int     `json:"transfer_gateaway_id"`
	Receipt_Id          int     `json:"receipt_id"`
	Description         string  `json:"description"`
	DateTime            string  `json:"date_time"`
}

type TransferInfo struct {
	SenderUsername  string  `json:"sender_username"`
	ReceiptUsername string  `json:"receipt_username"`
	TransferAmount  float32 `json:"transfer_amount"`
	PaymentGateway  string  `json:"payment_gateway"`
	Description     string  `json:"Description"`
}
