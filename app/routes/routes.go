package routes

import (
	"context"
	"ewallet/app/queries"
	"ewallet/platform/database"

	"github.com/gofiber/fiber/v2"
)

var storage queries.Storage


func Init(app *fiber.App) {
	db, err := database.NewClient(context.Background())
	if err != nil {
		panic(err)
	}

	storage = database.NewStorage(db)
	
	walletGroup := app.Group("/api/v1/wallet")

	walletGroup.Post("/", createWallet)
	walletGroup.Get("/:walletId", getWallet)
	walletGroup.Post("/:walletId/send", sendWallet)

}