package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func BasicAuthMiddleware(username string, password string) fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			username: password,
		},
		Authorizer: func(user, pass string) bool {
			fmt.Println("Step 1: authenticate")
			fmt.Println("env.BASIC_AUTH_USERNAME", username)
			fmt.Println("env.BASIC_AUTH_PASSWORD", password)

			if user == username && pass == password {
				return true
			}
			fmt.Println("Step 1: authenticate >> Failed")
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		},
	})
}
