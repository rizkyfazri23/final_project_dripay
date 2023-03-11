package entity

type Gateway struct {
	Gateway_Id   int    `json:"gateway_id"`
	Gateway_Name string `json:"gateway_name"`
	Status       bool    `json:"status"`
}
