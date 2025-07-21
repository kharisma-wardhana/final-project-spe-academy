package handler

import (
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/parser"
	"github.com/kharisma-wardhana/final-project-spe-academy/internal/presenter/json"
)

type TransactionHandler struct {
	parser    parser.Parser
	presenter json.JsonPresenter
}

func NewTransactionHandler(parser parser.Parser, presenter json.JsonPresenter) *TransactionHandler {
	return &TransactionHandler{
		parser:    parser,
		presenter: presenter,
	}
}
