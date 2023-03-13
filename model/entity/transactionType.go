package entity

type TransactionType struct {
	TypeId   int    `json:"type_id"`
	TypeName string `json:"type_name"`
}

type TransactionTypeInput struct {
	TypeName string `json:"type_name"`
}
