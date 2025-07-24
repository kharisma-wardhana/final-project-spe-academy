package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/parser"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/presenter/json"
	usecase_merchant "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/merchant"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/merchant/entity"
	usecase_qr "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/qr"
	qrEntity "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/qr/entity"
	usecase_transaction "github.com/kharisma-wardhana/final-project-spe-academy/internal/usecase/transaction"
)

type MerchantHandler struct {
	parser             parser.Parser
	presenter          json.JsonPresenter
	merchantUseCase    usecase_merchant.IMerchantUseCase
	transactionUseCase usecase_transaction.ITransactionUseCase
	qrUseCase          usecase_qr.IQRUseCase
}

func NewMerchantHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	merchantUseCase usecase_merchant.IMerchantUseCase,
	transactionUseCase usecase_transaction.ITransactionUseCase,
	qrUseCase usecase_qr.IQRUseCase,
) *MerchantHandler {
	return &MerchantHandler{parser, presenter, merchantUseCase, transactionUseCase, qrUseCase}
}

func (h *MerchantHandler) Register(app fiber.Router) {
	// Define your routes here
	app.Get("/merchants/:id", h.GetMerchantByID)
	app.Post("/merchants", h.CreateMerchant)
	app.Put("/merchants/:id", h.UpdateMerchant)
	app.Delete("/merchants/:id", h.DeleteMerchant)
	app.Get("/merchants/:id/transactions", h.GetMerchantTransactions)
	app.Post("/merchants/:id/qr", h.CreateQRForMerchant)
}

func (h *MerchantHandler) GetMerchantByID(c *fiber.Ctx) error {
	id, err := h.parser.ParserMerchantID(c)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	merchant, err := h.merchantUseCase.GetMerchantByMID(c.Context(), id)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, merchant, "Merchant successfully retrieved", http.StatusOK)
}

func (h *MerchantHandler) CreateMerchant(c *fiber.Ctx) error {
	var req entity.MerchantRequest
	if err := h.parser.ParserBodyRequest(c, &req); err != nil {
		return h.presenter.BuildError(c, err)
	}

	merchant, err := h.merchantUseCase.CreateMerchant(c.Context(), &req)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, merchant, "Merchant successfully created", http.StatusCreated)
}

func (h *MerchantHandler) UpdateMerchant(c *fiber.Ctx) error {
	id, err := h.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}
	var req entity.MerchantRequest
	if err := h.parser.ParserBodyRequest(c, &req); err != nil {
		return h.presenter.BuildError(c, err)
	}

	merchant, err := h.merchantUseCase.UpdateMerchant(c.Context(), id, &req)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, merchant, "Merchant successfully updated", http.StatusOK)
}

func (h *MerchantHandler) DeleteMerchant(c *fiber.Ctx) error {
	id, err := h.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	if err := h.merchantUseCase.DeleteMerchantByID(c.Context(), id); err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, nil, "Merchant successfully deleted", http.StatusOK)
}

func (h *MerchantHandler) GetMerchantTransactions(c *fiber.Ctx) error {
	id, err := h.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	transactions, err := h.transactionUseCase.GetTransactionsByMerchantID(c.Context(), id)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, transactions, "Merchant transactions successfully retrieved", http.StatusOK)
}

func (h *MerchantHandler) CreateQRForMerchant(c *fiber.Ctx) error {
	id, err := h.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	var req qrEntity.QRRequest
	if err := h.parser.ParserBodyRequest(c, &req); err != nil {
		return h.presenter.BuildError(c, err)
	}
	req.MerchantID = id
	qr, err := h.qrUseCase.GenerateQR(c.Context(), req)
	if err != nil {
		return h.presenter.BuildError(c, err)
	}

	return h.presenter.BuildSuccess(c, qr, "QR code successfully created", http.StatusCreated)
}
