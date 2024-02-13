package queries

import (
	"context"
	"ewallet/app/models"
)

type Storage interface {
	CreateWallet(ctx context.Context) (models.Wallet, error)
	FindWalletByID(ctx context.Context, walletId string) (models.Wallet, error)
	SendWallet(ctx context.Context, walletId string, toID string, amount float32) error
	UpdateWallet(ctx context.Context, wallet models.Wallet) error
	GetHistoryWallet(ctx context.Context, walletId string) ([]models.HistoryWallet, error)
}
