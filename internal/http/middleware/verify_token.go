package middleware

import (
	apperr "github.com/kharisma-wardhana/final-project-spe-academy/error"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/http/auth"

	"github.com/gofiber/fiber/v2"
)

func VerifyJWTToken(c *fiber.Ctx) error {
	if err := auth.VerifyToken(c); err != nil {
		return c.Status(apperr.ErrInvalidToken().HTTPCode).JSON(apperr.ErrInvalidToken())
	}

	return c.Next()
}
