package model

import "github.com/calmantara/go-e-wallet/entity"

type CommonResponse struct {
	Code string      `json:"response_code"`
	Type string      `json:"response_type"`
	Data interface{} `json:"data"`
}

type CommonErrorResponse struct {
	Code  string      `json:"response_code"`
	Type  string      `json:"response_type"`
	Error interface{} `json:"error_payload"`
}

type CommonRequest struct {
	WalletID    int                             `json:"wallet_id"`
	Type        entity.TransactionType          `json:"transaction_type"`
	TotalAmount int                             `json:"total_amount"`
	Data        []entity.WalletTransactionModel `json:"data"`
}
