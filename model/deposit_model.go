package model

import "time"

type Deposit struct {
	Id                 int       `json:"deposit_id"`
	Deposit_Code       string    `json:"deposit_code"`
	Member_Id          int       `json:"member_id"`
	Deposit_Amount     float32   `json:"deposit_amount"`
	Payment_Gateway_Id int       `json:"payment_gateway_id"`
	Description        string    `json:"description"`
	Date_time          time.Time `json:"date_time"`
}

type DepositRequest struct {
	Member_Id          int    `json:"member_id"`
	Deposit_Code       string `json:"deposit_code"`
	Payment_Gateway_Id int    `json:"payment_gateway_id"`
}
