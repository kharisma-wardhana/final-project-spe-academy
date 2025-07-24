package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/http/auth"
)

func VerifySignature(c *fiber.Ctx, signature auth.ISignature) error {
	if err := signature.VerifySignature(c); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid signature",
		})
	}
	return c.Next()
}
