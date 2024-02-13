package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func getWallet(c *fiber.Ctx) error {
	println("Handle get wallet")

	walletID := c.Params("walletId")
	walletData, err := storage.FindWalletByID(context.Background(), walletID)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"description": "Указанный кошелек не найден",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"description": "OK",
		"content": fiber.Map{ 
			"id":     walletData.ID,
			"balance": walletData.Balance,
		},
	})

}
