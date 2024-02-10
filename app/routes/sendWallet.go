package routes

import (
	"context"
	"ewallet/app/models"

	"github.com/gofiber/fiber/v2"

)

type sendWalletRequestDTO struct {
	ToWallet models.ToWallet `json:"wallet"`
}

func sendWallet(c *fiber.Ctx) error {
	println("Handle send wallet")

	queryInfo := new(sendWalletRequestDTO)
	if err := c.BodyParser(queryInfo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	walletID := c.Params("walletId")
	err := storage.SendWallet(context.Background(), walletID, queryInfo.ToWallet)

	if err != nil {
		if err.Error() == "insufficient funds" {
			return c.Status(400).JSON(fiber.Map{
				"description": "нет денег",
			})
		} 

		return c.Status(404).JSON(fiber.Map{
			"description": "Указанный кошелек не найден",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"description": "ОК",
	})

}
