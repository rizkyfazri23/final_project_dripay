package entity

import "time"

type History struct {
	Id               int       `json:"transaction_log_id"`
	Member_Username  string    `json:"member_username"`
	Transaction_Type string    `json:"transaction_type"`
	Amount           float32   `json:"amount"`
	Date_Time        time.Time `json:"date_time"`
	Status           int       `json:"status"`
	Transaction_Code string    `json:"transaction_code"`
}
