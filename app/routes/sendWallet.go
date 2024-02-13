package routes

import (
	"context"
	"ewallet/app/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// type sendWalletRequestDTO struct {
//     ToWallet models.ToWallet `json:"wallet" validate:"required"`
// }

type sendWalletRequestDTO struct {
    ToID     string  `json:"to" bson:"_id,omitempty" validate:"required,hexadecimal,len=24"`
    Amount float32 `json:"amount" bson:"amount" validate:"gt=0"`
}

func sendWallet(c *fiber.Ctx) error {
	println("Handle send wallet")
	
	queryInfo := new(models.ToWallet)
	if err := c.BodyParser(&queryInfo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	validate := validator.New()
    err := validate.Struct(queryInfo)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
			"description": "Ошибка в пользовательском запросе или ошибка перевода",
		})
    }

	walletID := c.Params("walletId")
	err = storage.SendWallet(context.Background(), walletID , queryInfo.ToID, queryInfo.Amount) // если будешь успевать то задай вопрос в ТП о том чтобы добавить проверку на исходщий/входящий поиск кошелька( различные ошибки)

	if err != nil {
		if err.Error() == "insufficient funds" {
			return c.Status(400).JSON(fiber.Map{
				"description": "Ошибка в пользовательском запросе или ошибка перевода",
			})
		} 

		return c.Status(404).JSON(fiber.Map{
			"description": "Исходящий кошелек не найден",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"description": "Перевод успешно проведен",
	})

}
