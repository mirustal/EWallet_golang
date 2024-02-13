package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
)


func createWallet(c *fiber.Ctx) error {
	println("Handle create wallet")

	walletData, err := storage.CreateWallet(context.Background())

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Ошибка в запросе",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"description": "Кошелек создан",
		"content": fiber.Map{
			"id":     walletData.ID,
			"amount": walletData.Balance,
		},
	})
}