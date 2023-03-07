package model

type Gateway struct {
	Gateway_Id   int    `json:"gateway_id"`
	Gateway_Name string `json:"gateway_name"`
	Type         string `json:"type"`
	Status       int    `json:"status"`
}
