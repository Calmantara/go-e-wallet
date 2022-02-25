package ewallethandler

import "github.com/gin-gonic/gin"

type EWalletHandler interface {
	PostOrganizationPayment(ctx *gin.Context)
	GetOrganizationPayment(ctx *gin.Context)
}
