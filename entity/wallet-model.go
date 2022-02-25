package entity

import "time"

type WalletModel struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	OrganizationID int       `json:"organization_id"`
	Balance        int       `json:"balance"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	DeletedAt      time.Time `json:"-"`

	TransactionDetails *[]WalletModel `json:"wallet_details,omitempty"`
}
