package routes

import (
	"context"

	"ewallet/app/queries"
	"ewallet/platform/database"

	"strings"

	"github.com/dgrijalva/jwt-go"
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

	walletGroup.Post("/", authorize, createWallet)
	walletGroup.Get("/:walletId", getWallet)
	walletGroup.Post("/:walletId/send", authorize, sendWallet)
	walletGroup.Get("/:walletId/history", getHistory)
}




var jwtKey = []byte("wallet")

func authorize(ctx *fiber.Ctx) error {

	authHeader := ctx.Get("Authorization")
	if authHeader == "" {

		return ctx.Status(fiber.StatusUnauthorized).SendString("Missing or invalid JWT token")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return ctx.Status(fiber.StatusUnauthorized).SendString("Invalid JWT token format")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).SendString("Invalid JWT token")
	}

	return ctx.Next()
}
