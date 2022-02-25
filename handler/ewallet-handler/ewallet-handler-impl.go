package ewallethandler

import (
	"net/http"

	"github.com/calmantara/go-e-wallet/entity"
	"github.com/calmantara/go-e-wallet/model"
	ewalletservice "github.com/calmantara/go-e-wallet/service/ewallet-service"
	"github.com/gin-gonic/gin"
)

type EWalletHandlerImpl struct {
	ewalletService ewalletservice.EWalletService
}

func NewEWalletHandlerImpl(ewalletService ewalletservice.EWalletService) EWalletHandler {
	return &EWalletHandlerImpl{
		ewalletService: ewalletService,
	}
}

func (e *EWalletHandlerImpl) PostOrganizationPayment(ctx *gin.Context) {
	// bind data
	var commonRequest model.CommonRequest
	if err := ctx.ShouldBind(&commonRequest); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			model.CommonErrorResponse{
				Code:  "99",
				Type:  "BAD_REQUEST",
				Error: err.Error(),
			},
		)
		return
	}

	if err := e.ewalletService.StoreEWalletTransaction(ctx.Copy(), &commonRequest); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			model.CommonErrorResponse{
				Code:  "99",
				Type:  "BAD_REQUEST",
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSONP(
		http.StatusAccepted,
		model.CommonResponse{
			Code: "00",
			Type: "ACCEPTED",
			Data: "transaction successfully",
		},
	)
}
func (e *EWalletHandlerImpl) GetOrganizationPayment(ctx *gin.Context) {
	walletID := ctx.Query("wallet_id")
	if walletID == "" {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			model.CommonErrorResponse{
				Code:  "99",
				Type:  "BAD_REQUEST",
				Error: "wallet id payload is empty",
			},
		)
		return
	}

	var organization *entity.OrganizationModel
	var err error
	if err, organization = e.ewalletService.GetEWalletTransactionByOrganizationID(ctx.Copy(), walletID); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			model.CommonErrorResponse{
				Code:  "99",
				Type:  "BAD_REQUEST",
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSONP(
		http.StatusOK,
		model.CommonResponse{
			Code: "00",
			Type: "ACCEPTED",
			Data: organization,
		},
	)
}
