package entity

import "time"

type Deposit struct {
	Id                 int       `json:"deposit_id"`
	Deposit_Code       string    `json:"deposit_code"`
	Member_Id          int       `json:"member_id"`
	Deposit_Amount     float32   `json:"deposit_amount"`
	Deposit_Gateway_Id int       `json:"deposit_gateway_id"`
	Description        string    `json:"description"`
	Date_Time          time.Time `json:"date_time"`
}

type DepositRequest struct {
	Member_Username string  `json:"member_username"`
	Deposit_Amount  float32 `json:"deposit_amount"`
	Deposit_Gateway string  `json:"deposit_gateway"`
	Description     string  `json:"description"`
}
