package ewalletservice

import (
	"context"

	"github.com/calmantara/go-e-wallet/entity"
	"github.com/calmantara/go-e-wallet/model"
)

type EWalletService interface {
	StoreEWalletTransaction(ctx context.Context, commonRequest *model.CommonRequest) (err error)
	GetEWalletTransactionByOrganizationID(ctx context.Context, walletID string) (err error, organizationEntity *entity.OrganizationModel)
}
