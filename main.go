package main

import (
	"ewallet/app/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization",
	}))

	app.Use(limiter.New(limiter.Config{
		Max:        1, 
		Expiration: 1 * time.Second, 
		LimitReached: func(c *fiber.Ctx) error { 
			return c.Status(fiber.StatusTooManyRequests).SendString("Превышен лимит запросов")
		},
	}))




	routes.Init(app)

	app.Listen(":8080")

}




// token == eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkFkbWluIiwiaWF0IjoiQWRtaW4ifQ.8NqkbR4i2NTzeMA9J8Qnn23yx3nzlO4E8YxJF1XspOU
// var jwtKey = []byte("wallet")

// type CustomClaims struct {
// 	Sub  string `json:"sub"`
// 	Name string `json:"name"`
// 	Iat  string `json:"iat"`
// 	jwt.StandardClaims
// }

// func CreateToken() (string, error) {
// 	claims := CustomClaims{
// 		Sub:  "1234567890",
// 		Name: "Admin",
// 		Iat:  "Admin",
// 		// StandardClaims: jwt.StandardClaims{
// 		// 	ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), 
// 		// },
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }
