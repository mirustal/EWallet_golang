package models

type Wallet struct {
	ID string `json:"id" bson:"_id,omitempty"`
	Balance float32 `json:"balance" bson:"balance"`
}

type ToWallet struct {
	ID string `json:"id" bson:"_id,omitempty"`
	Amount float32 `json:"amount" bson:"amount"`
}