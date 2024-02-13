package models

type Wallet struct {
    ID      string  `json:"id" bson:"_id,omitempty" validate:"required,hexadecimal,len=24"`
    Balance float32 `json:"balance" bson:"balance" validate:"gte=0"`
	HistoryTransaction []HistoryWallet
}

type ToWallet struct {
    ToID     string  `json:"to" bson:"_id,omitempty" validate:"required,hexadecimal,len=24"`
    Amount float32 `json:"amount" bson:"amount" validate:"gt=0"`
}


type HistoryWallet struct{
	TimeTransaction string	`json:"time" bson:"time"`
	FromWalletId string `json:"from" bson:"from validate:"required,hexadecimal,len=24"`
	ToWalletId string `json:"to" bson:"to" validate:"required,hexadecimal,len=24`
	Amount float32 `json:"amount" bson:"amount"`

}