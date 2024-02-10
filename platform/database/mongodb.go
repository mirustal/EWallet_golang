package database

import (
	"context"
	"errors"
	"ewallet/app/models"
	"ewallet/app/queries"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	wallets *mongo.Collection
	states  *mongo.Collection
}

func (d *db) CreateWallet(ctx context.Context) (models.Wallet, error) {
	// Создаем новый кошелек с начальным балансом
	wallet := models.Wallet{
		Balance: 100,
	}

	// Вставляем данные кошелька в коллекцию
	resultInsert, err := d.wallets.InsertOne(ctx, wallet)
	if err != nil {
		return models.Wallet{}, err
	}

	// Получаем вставленный идентификатор
	oid, ok := resultInsert.InsertedID.(primitive.ObjectID)
	if !ok {
		return models.Wallet{}, err
	}

	// Присваиваем идентификатор кошелька
	wallet.ID = oid.Hex()

	return wallet, nil
}

func (d *db) FindWalletByID(ctx context.Context, walletId string) (models.Wallet, error) {
	wallet := models.Wallet{}

	objIdWallet, _ := primitive.ObjectIDFromHex(walletId)
	filter := bson.D{{"_id", objIdWallet}}

	err := d.wallets.FindOne(ctx, filter).Decode(&wallet)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Wallet{}, err
		}
	}
	return wallet, nil
}

func (d *db) SendWallet(ctx context.Context, walletID string, toWallet models.ToWallet) error {
	wallet, err := d.FindWalletByID(ctx, walletID)
	if err != nil {
		return err
	}

	recipientWallet, err := d.FindWalletByID(ctx, toWallet.ID)
	if err != nil {
		return err
	}

	if wallet.Balance < toWallet.Amount {
		return errors.New("insufficient funds")
	}

	wallet.Balance -= toWallet.Amount
	recipientWallet.Balance += toWallet.Amount

	if err := d.UpdateWallet(ctx, recipientWallet); err != nil {
		return err
	}
	if err := d.UpdateWallet(ctx, wallet); err != nil {
		return err
	}

	return nil
}

func (d *db) UpdateWallet(ctx context.Context, wallet models.Wallet) error {
	objectId, _ := primitive.ObjectIDFromHex(wallet.ID)
	filter := bson.M{"_id": objectId}
	update := bson.M{"balance": wallet.Balance}

	_, err := d.wallets.ReplaceOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func NewStorage(database *mongo.Database) queries.Storage {
	return &db{
		wallets: database.Collection("wallets"),
	}
}
