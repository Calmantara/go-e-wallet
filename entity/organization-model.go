package entity

import "time"

type OrganizationModel struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	PIC         string    `json:"pic"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   time.Time `json:"-"`

	WalletDetails []WalletModel `json:"wallet_details"`
}
