package entity

import "time"

type memberList struct {
	Member_Username string
}

type SplitRequest struct {
	Member_List  []memberList `json:"member_list"`
	Total_Amount float32      `json:"total_amount"`
	Description  string       `json:"description"`
}

type SplitResponse struct {
	Payment_Id      int       `json:"payment_id"`
	Payment_Code    string    `json:"payment_code"`
	Member_Username string    `json:"member_username"`
	Payment_Amount  float32   `json:"payment_amount"`
	Payment_Gateway string    `json:"payment_gateway"`
	Description     string    `json:"description"`
	Status          string    `json:"status"`
	Date_Time       time.Time `json:"date_time"`
}
