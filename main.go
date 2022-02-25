package main

import (
	"net/http"
	"time"

	"github.com/calmantara/go-e-wallet/entity"
	ewallethandler "github.com/calmantara/go-e-wallet/handler/ewallet-handler"
	ewalletservice "github.com/calmantara/go-e-wallet/service/ewallet-service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		//define wallet
		wallet1 = entity.WalletModel{
			ID:             0,
			Name:           "wallet1",
			OrganizationID: 0,
			Balance:        0,
			CreatedAt:      time.Now(),
		}
		wallet2 = entity.WalletModel{
			ID:             1,
			Name:           "wallet2",
			OrganizationID: 0,
			Balance:        0,
			CreatedAt:      time.Now(),
		}
		//define org
		organization = entity.OrganizationModel{
			ID:            0,
			Name:          "organization0",
			PIC:           "pic_0",
			PhoneNumber:   "081234567890",
			Email:         "cs@organization0.com",
			CreatedAt:     time.Now(),
			WalletDetails: []entity.WalletModel{wallet1, wallet2},
		}

		//define dependencies injection
		eWalletService = ewalletservice.NewEWalletService(organization)
		eWalletHandler = ewallethandler.NewEWalletHandlerImpl(eWalletService)
	)

	//middleware
	rg := r.Group("ewallet/v1/payment", func(c *gin.Context) {
		v := c.Param("organizationID")
		if v != "organization0" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				"organization is not found",
			)
			return
		}
		c.Next()
	})

	rg.GET("/:organizationID", eWalletHandler.GetOrganizationPayment)
	rg.POST("/:organizationID", eWalletHandler.PostOrganizationPayment)

	r.Run()
}
