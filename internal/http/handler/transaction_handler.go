package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/parser"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/presenter/json"
	usecase_transaction "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/transaction"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/transaction/entity"
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
	id, err := h.parser.ParserMerchantID(c)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}
	transaction, err := h.usecase.GetTransactionsByRefID(c.Context(), id)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, transaction, "Transaction successfully retrieved", http.StatusOK)
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var req entity.TransactionRequest
	err := h.parser.ParserBodyRequest(c, &req)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	transaction, err := h.usecase.CreateTransaction(c.Context(), &req)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, transaction, "Transaction successfully created", http.StatusCreated)
}
