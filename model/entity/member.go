package entity

// import (
// 	"golang.org/x/crypto/bcrypt"
// 	_ "github.com/lib/pq"
// ) 

type Member struct {
	Member_Id      int     `json:"member_id"`
	Username       string  `json:"username"`
	Password       string  `json:"password"`
	Email_Address  string  `json:"email_address"`
	Contact_Number string  `json:"contact_number"`
	Wallet_Amount  float64 `json:"wallet_amount"`
	Status         bool    `json:"status"`
}

type MemberLogin struct {
	Member_Id int 	`json:"member_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}


