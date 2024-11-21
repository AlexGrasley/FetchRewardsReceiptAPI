package main

import (
	"FetchRewardsReceiptAPI/Controllers"
	"FetchRewardsReceiptAPI/Repositories"
	"FetchRewardsReceiptAPI/Services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	InitDependencies(router)
	router.Run(":8080")
}

func InitDependencies(router *gin.Engine) {
	// initialize dependencies at app startup
	// future enhancement separate global persistent dependencies
	// from transient dependencies
	var itemRepo = Repositories.NewItemRepository()
	var receiptRepo = Repositories.NewReceiptRepository(itemRepo)
	var receiptService = Services.NewReceiptService(receiptRepo)
	Controllers.InitReceiptController(router, receiptService)
}
