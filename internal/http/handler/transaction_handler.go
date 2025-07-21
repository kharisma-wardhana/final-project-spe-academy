package handler

import (
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/internal/parser"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/internal/presenter/json"
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
