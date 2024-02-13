package models

type Wallet struct {
	ID string `json:"id" bson:"_id,omitempty"`
	Balance float32 `json:"balance" bson:"balance"`
	HistoryTransaction []HistoryWallet
}

type ToWallet struct {
	ToID string `json:"to" bson:"_id,omitempty"`
	Amount float32 `json:"amount" bson:"amount"`
}

type HistoryWallet struct{
	TimeTransaction string	`json:"time" bson:"time`
	FromWalletId string `json:"from" bson:"from"`
	ToWalletId string `json:"to" bson:"to"`
	Amount float32 `json:"amount" bson:"amount"`
}