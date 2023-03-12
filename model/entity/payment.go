package entity

import "time"

type Payment struct {
	Id                 int       `json:"payment_id"`
	Payment_Code       string    `json:"payment_code"`
	Member_Id          int       `json:"member_id"`
	Payment_Amount     float32   `json:"payment_amount"`
	Payment_Gateway_Id int       `json:"payment_gateway_id"`
	Description        string    `json:"description"`
	Status             string    `json:"status"`
	Date_Time          time.Time `json:"date_time"`
}

type PaymentRequest struct {
	Username        string  `json:"member_username"`
	Payment_Amount  float32 `json:"payment_amount"`
	Payment_Gateway string  `json:"payment_gateway"`
	Description     string  `json:"description"`
}

type PaymentRequestStatus struct {
	Status string `json:"status"`
}
