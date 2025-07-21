package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/parser"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/presenter/json"
	usecase_merchant "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/merchant"
	usecase_transaction "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/transaction"
)

type MerchantHandler struct {
	parser             parser.Parser
	presenter          json.JsonPresenter
	accountUseCase     usecase_merchant.IMerchantUseCase
	transactionUseCase usecase_transaction.ITransactionUseCase
}

func NewMerchantHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	accountUseCase usecase_merchant.IMerchantUseCase,
	transactionUseCase usecase_transaction.ITransactionUseCase,
) *MerchantHandler {
	return &MerchantHandler{parser, presenter, accountUseCase, transactionUseCase}
}

func (h *MerchantHandler) Register(app fiber.Router) {
	// Define your routes here
	app.Get("/merchants/:id", h.GetMerchantByID)
	app.Post("/merchants", h.CreateMerchant)
	app.Put("/merchants/:id", h.UpdateMerchant)
	app.Delete("/merchants/:id", h.DeleteMerchant)
	app.Get("/merchants/:id/transactions", h.GetMerchantTransactions)
}

func (h *MerchantHandler) GetMerchantByID(c *fiber.Ctx) error {
	// Implementation for getting a merchant by ID
	return nil
}

func (h *MerchantHandler) CreateMerchant(c *fiber.Ctx) error {
	// Implementation for creating a new merchant
	return nil
}

func (h *MerchantHandler) UpdateMerchant(c *fiber.Ctx) error {
	// Implementation for updating a merchant
	return nil
}

func (h *MerchantHandler) DeleteMerchant(c *fiber.Ctx) error {
	// Implementation for deleting a merchant
	return nil
}

func (h *MerchantHandler) GetMerchantTransactions(c *fiber.Ctx) error {
	// Implementation for getting transactions of a merchant
	return nil
}
