package handler

import (
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/internal/parser"
	"github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/internal/presenter/json"
)

type MerchantHandler struct {
	parser    parser.Parser
	presenter json.JsonPresenter
}

func NewMerchantHandler(parser parser.Parser, presenter json.JsonPresenter) *MerchantHandler {
	return &MerchantHandler{
		parser:    parser,
		presenter: presenter,
	}
}
