package entity

import "time"

type TransactionType string
type TransactionStatus string

const ADDING_TYPE TransactionType = "ADDING"
const SENDING_TYPE TransactionType = "SENDING"

const ACTIVE_STATUS TransactionStatus = "ACTIVE"
const SUCCESS_STATUS TransactionStatus = "SUCCESS"
const CANCELED_STATUS TransactionStatus = "CANCELED"
const EXPIRED_STATUS TransactionStatus = "EXPIRED"

type WalletTransactionModel struct {
	ID        int               `json:"id"`
	WalletID  int               `json:"wallet_id"`
	Type      TransactionType   `json:"type"`
	Recipient string            `json:"recipient"`
	Amount    int               `json:"amount"`
	Status    TransactionStatus `json:"status"`
	ExpiredAt time.Time         `json:"expired_at,omitempty"`
	CreatedAt time.Time         `json:"-,omitempty"`
	UpdatedAt time.Time         `json:"-,omitempty"`
	DeletedAt time.Time         `json:"-,omitempty"`
}
