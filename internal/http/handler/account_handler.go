package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/internal/parser"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/internal/presenter/json"
)

type AccountHandler struct {
	parser    parser.Parser
	presenter json.JsonPresenter
}

func NewAccountHandler(parser parser.Parser, presenter json.JsonPresenter) *AccountHandler {
	return &AccountHandler{
		parser:    parser,
		presenter: presenter,
	}
}

func (h *AccountHandler) Register(app fiber.Router) {
	// Define your routes here
	app.Get("/accounts/:id", h.GetAccountByID)
	app.Get("/accounts/merchant/:merchantID", h.GetAccountByMerchantID)
}

func (h *AccountHandler) GetAccountByID(c *fiber.Ctx) error {
	id, err := h.parser.ParseUint64(c.Params("id"))
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	account, err := h.presenter.GetAccountByID(c.Context(), id)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, account, "Success", http.StatusOK)
}

func (h *AccountHandler) GetAccountByMerchantID(c *fiber.Ctx) error {
	merchantID, err := h.parser.ParseUint64(c.Params("merchantID"))
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	account, err := h.presenter.GetAccountByMerchantID(c.Context(), merchantID)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, account, "Success", http.StatusOK)
}
