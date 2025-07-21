package entity

type QRType uint8

type MerchantType uint8

type TransactionType uint8

const (
	QRTypeSticker          QRType          = 1
	QRTypeAPI              QRType          = 2
	MerchantTypeIndividual MerchantType    = 1
	MerchantTypeCorporate  MerchantType    = 2
	TransactionTypePayment TransactionType = 1
	TransactionTypeRefund  TransactionType = 2
)
