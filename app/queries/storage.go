package queries

import (
	"context"
	"ewallet/app/models"
)

type Storage interface {
	CreateWallet(ctx context.Context) (models.Wallet, error)
	FindWalletByID(ctx context.Context, walletId string) (models.Wallet, error)
	SendWallet(ctx context.Context, walletId string, toWallet models.ToWallet) error
	UpdateWallet(ctx context.Context, wallet models.Wallet) error
}
