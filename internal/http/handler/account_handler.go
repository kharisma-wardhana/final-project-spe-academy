package handler

import (
	"net/http"

	"github.com/kharisma-wardhana/final-project-spe-academy/internal/parser"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/presenter/json"
	usecase_account "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/account"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/account/entity"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	parser    parser.Parser
	presenter json.JsonPresenter
	usecase   usecase_account.IAccountUseCase
}

func NewAccountHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	usecase usecase_account.IAccountUseCase,
) *AccountHandler {
	return &AccountHandler{
		parser:    parser,
		presenter: presenter,
		usecase:   usecase,
	}
}

func (h *AccountHandler) Register(app fiber.Router) {
	// Define your routes here
	app.Get("/accounts/:id", h.GetAccountByID)
	app.Get("/accounts/:id/merchants", h.GetAccountMerchants)
	app.Post("/accounts", h.CreateAccount)
	app.Put("/accounts/:id", h.UpdateAccount)
	app.Delete("/accounts/:id", h.DeleteAccount)
}

func (h *AccountHandler) GetAccountByID(c *fiber.Ctx) error {
	id, err := h.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	account, err := h.usecase.GetAccountByID(c.Context(), id)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, account, "Account successfully retrieved", http.StatusOK)
}

func (h *AccountHandler) GetAccountMerchants(c *fiber.Ctx) error {
	id, err := h.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	account, err := h.usecase.GetAccountByMerchantID(c.Context(), id)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, account, "Merchants in account sucessfully retrieved", http.StatusOK)
}

func (h *AccountHandler) CreateAccount(c *fiber.Ctx) error {
	var accountRequest *entity.AccountRequest

	err := h.parser.ParserBodyRequest(c, &accountRequest)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	accountResponse, err := h.usecase.CreateAccount(c.Context(), accountRequest)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, accountResponse, "Account created successfully", http.StatusCreated)
}

func (h *AccountHandler) UpdateAccount(c *fiber.Ctx) error {
	id, err := h.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	var accountRequest *entity.AccountRequest
	err = h.parser.ParserBodyRequest(c, &accountRequest)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	accountResponse, err := h.usecase.UpdateAccount(c.Context(), id, accountRequest)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, accountResponse, "Account updated successfully", http.StatusOK)
}

func (h *AccountHandler) DeleteAccount(c *fiber.Ctx) error {
	id, err := h.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	err = h.usecase.DeleteAccount(c.Context(), id)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, nil, "Account deleted successfully", http.StatusNoContent)
}
