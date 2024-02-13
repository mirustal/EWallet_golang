package database

import (
	"context"
	"errors"
	"ewallet/app/models"
	"ewallet/app/queries"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	wallets *mongo.Collection
	states  *mongo.Collection
}

func (d *db) CreateWallet(ctx context.Context) (models.Wallet, error) {
	// Создаем новый кошелек
	wallet := models.Wallet{
		Balance: 100,
		HistoryTransaction: []models.HistoryWallet{},
	}

	// Вставляем данные кошелька в коллекцию
	resultInsert, err := d.wallets.InsertOne(ctx, wallet)
	if err != nil {
		return models.Wallet{}, err
	}

	// Получаем  идентификатор
	oid, ok := resultInsert.InsertedID.(primitive.ObjectID)
	if !ok {
		return models.Wallet{}, err
	}


	wallet.ID = oid.Hex()

	return wallet, nil
}

func (d *db) FindWalletByID(ctx context.Context, walletId string) (models.Wallet, error) {
	wallet := models.Wallet{}

	objIdWallet, _ := primitive.ObjectIDFromHex(walletId)
	filter := bson.D{{"_id", objIdWallet}}
	// Ищем нужный объект в коллекции
	err := d.wallets.FindOne(ctx, filter).Decode(&wallet)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Wallet{}, err
		}
	}
	return wallet, nil
}

func (d *db) SendWallet(ctx context.Context, walletID string, toID string, amount float32) error {
	wallet, err := d.FindWalletByID(ctx, walletID)
	if err != nil {
		return err
	}

	recipientWallet, err := d.FindWalletByID(ctx, toID)
	if err != nil {
		return err
	}
	
	if wallet.Balance <= amount {
		return errors.New("insufficient funds")
	}

	
	currentTime := time.Now().Format(time.RFC3339) 


	newHistoryWallet := models.HistoryWallet{
		FromWalletId: walletID,
		ToWalletId: toID,
		Amount: amount,
		TimeTransaction: currentTime,
	}

	wallet.HistoryTransaction = append(wallet.HistoryTransaction, newHistoryWallet)

	
	wallet.Balance -= amount
	recipientWallet.Balance += amount


	recipientHistoryWallet := models.HistoryWallet{
		FromWalletId: walletID,
		ToWalletId: toID,
		Amount: amount,
		TimeTransaction: currentTime,
	}
	recipientWallet.HistoryTransaction = append(recipientWallet.HistoryTransaction, recipientHistoryWallet)


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
	update := bson.M{
			"balance":            wallet.Balance,
			"historyTransaction": wallet.HistoryTransaction,
	}

	_, err := d.wallets.ReplaceOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (d *db) GetHistoryWallet(ctx context.Context, walletId string) ([]models.HistoryWallet, error){
	wallet, err := d.FindWalletByID(ctx, walletId)
	if err != nil {
			return nil, err
	}
	return wallet.HistoryTransaction, nil
}

func NewStorage(database *mongo.Database) queries.Storage {
	return &db{
		wallets: database.Collection("wallets"),
	}
}
