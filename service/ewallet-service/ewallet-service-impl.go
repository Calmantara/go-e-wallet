package ewalletservice

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/calmantara/go-e-wallet/entity"
	"github.com/calmantara/go-e-wallet/model"
)

type EWalletServiceImpl struct {
	organization        entity.OrganizationModel
	ewalletTransactions []entity.WalletTransactionModel
}

func NewEWalletService(organization entity.OrganizationModel) EWalletService {
	return &EWalletServiceImpl{
		organization:        organization,
		ewalletTransactions: []entity.WalletTransactionModel{},
	}
}

func (e *EWalletServiceImpl) StoreEWalletTransaction(ctx context.Context, commonRequest *model.CommonRequest) (err error) {
	if valid := e.validateWalletID(commonRequest.WalletID); !valid {
		err = errors.New("wallet is not found")
		return err
	}

	// parsing data
	b, err := json.Marshal(&commonRequest.Data)
	if err != nil {
		err = errors.New("error parsing request data")
		return err
	}

	var data []entity.WalletTransactionModel
	err = json.Unmarshal(b, &data)
	if err != nil {
		err = errors.New("error parsing request data")
		return err
	}

	//check balance
	if commonRequest.Type == entity.SENDING_TYPE {
		if commonRequest.TotalAmount > e.organization.WalletDetails[commonRequest.WalletID].Balance {
			err = errors.New("wallet is insufficient balance")
			return err
		} else {
			tid := len(e.ewalletTransactions)
			for i, val := range data {
				if val.Recipient == "" {
					err = errors.New("transaction data is not valid")
					return err
				}
				tid++
				data[i].ID = tid
			}

			e.organization.WalletDetails[commonRequest.WalletID].Balance -= commonRequest.TotalAmount
		}
	} else if commonRequest.Type == entity.ADDING_TYPE {
		tid := len(e.ewalletTransactions)
		for i := range data {
			tid++
			data[i].ID = tid
		}
		e.organization.WalletDetails[commonRequest.WalletID].Balance += commonRequest.TotalAmount
	}
	// store transaction
	e.ewalletTransactions = append(e.ewalletTransactions, data...)

	return err
}

func (e *EWalletServiceImpl) validateWalletID(walletID int) bool {
	//check wallet exist or not
	isExist := false
	for _, v := range e.organization.WalletDetails {
		if v.ID == walletID {
			isExist = true
			break
		}
	}
	return isExist
}

func (e *EWalletServiceImpl) GetEWalletTransactionByOrganizationID(
	ctx context.Context, walletID string) (err error, organizationEntity *entity.OrganizationModel) {

	organizationEntity = &e.organization
	if walletID != "all" {
		wal, err := strconv.Atoi(walletID)
		if err != nil {
			err = errors.New("wallet is not valid")
			return err, nil
		}

		if valid := e.validateWalletID(wal); !valid {
			err = errors.New("wallet is not found")
			return err, nil
		}
		organizationEntity.WalletDetails = []entity.WalletModel{
			e.organization.WalletDetails[wal],
		}
	}

	log.Println("sending organization detail", organizationEntity.ID)
	return err, organizationEntity
}
