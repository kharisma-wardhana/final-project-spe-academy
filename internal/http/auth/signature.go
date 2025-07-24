package auth

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/parser"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/repository/mysql/entity"
)

type ISignature interface {
	VerifySignature(c *fiber.Ctx) error
}

type Signature struct {
	// Add any dependencies needed for signature verification here
	parser       parser.Parser
	accountRepo  mysql.IAccountRepository
	merchantRepo mysql.IMerchantRepository
}

func NewSignature(parser parser.Parser, accountRepo mysql.IAccountRepository, merchantRepo mysql.IMerchantRepository) ISignature {
	return &Signature{
		parser:       parser,
		accountRepo:  accountRepo,
		merchantRepo: merchantRepo,
	}
}

func (u *Signature) VerifySignature(c *fiber.Ctx) error {
	clientId := c.Get("X-Client-ID")
	if clientId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "X-Client-ID header is missing",
		})
	}

	signatureString := c.Get("X-Signature")
	if signatureString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "X-Signature header is missing",
		})
	}

	account, err := u.accountRepo.FindByClientID(c.Context(), clientId)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid account",
		})
	}

	if !isValidSignature(account, signatureString) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid signature",
		})
	}

	return c.Next()
}

func isValidSignature(account *entity.AccountEntity, signature string) bool {
	// Check Signature
	data := account.ClientID + account.ClientSecret + account.PublicKey
	hash := sha256.Sum256([]byte(data))
	expectedSignature := hex.EncodeToString(hash[:])
	return expectedSignature == signature
}
