package entity

import "time"

type History struct {
	Id               int       `json:"transaction_log_id"`
	Member_Username  string    `json:"member_username"`
	Transaction_Type string    `json:"transaction_type"`
	Kredit           float32   `json:"kredit"`
	Debit            float32   `json:"debit"`
	Date_Time        time.Time `json:"date_time"`
	Status           string    `json:"status"`
	Transaction_Code string    `json:"transaction_code"`
}
