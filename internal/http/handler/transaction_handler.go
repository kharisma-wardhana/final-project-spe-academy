package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/parser"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/presenter/json"
	usecase_transaction "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/transaction"
)

type TransactionHandler struct {
	parser    parser.Parser
	presenter json.JsonPresenter
	usecase   usecase_transaction.ITransactionUseCase
}

func NewTransactionHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	usecase usecase_transaction.ITransactionUseCase,
) *TransactionHandler {
	return &TransactionHandler{
		parser:    parser,
		presenter: presenter,
		usecase:   usecase,
	}
}

func (h *TransactionHandler) Register(app fiber.Router) {
	// Define your routes here
	app.Get("/transactions/:id", h.GetTransactionByID)
	app.Post("/transactions", h.CreateTransaction)
}

func (h *TransactionHandler) GetTransactionByID(c *fiber.Ctx) error {
	// Implementation for getting a transaction by ID
	return nil
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	// Implementation for creating a new transaction
	return nil
}
