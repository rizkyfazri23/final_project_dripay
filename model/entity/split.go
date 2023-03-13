package entity

type MemberList struct {
	Member_Username string `json:"member_username"`
}

type SplitRequest struct {
	Member_List  []MemberList `json:"member_list"`
	Total_Amount float32      `json:"total_amount"`
	Description  string       `json:"description"`
}

type SplitResponse struct {
	Payment_Code    string  `json:"payment_code"`
	Member_Username string  `json:"member_username"`
	Payment_Amount  float32 `json:"payment_amount"`
	Payment_Gateway string  `json:"payment_gateway"`
	Description     string  `json:"description"`
	Status          string  `json:"status"`
}
