package Controllers

import (
	"FetchRewardsReceiptAPI/Models"
	"FetchRewardsReceiptAPI/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _receiptService *Services.ReceiptService

func GetReceiptPoints(c *gin.Context) {
	// Logic to fetch receipts from the database or any data source
	//receipts := []Models.ReceiptViewModel{
	//	// Example data
	//	{Retailer: "Retailer1", PurchaseDate: "2023-10-01", PurchaseTime: "10:00", Total: 100.0, Items: []Models.ItemViewModel{{Description: "Item1", Price: 50.0}}},
	//}
	//c.JSON(http.StatusOK, receipts)
}

// CreateReceipt handles POST requests to create a new receipt
func CreateReceipt(c *gin.Context) {
	var newReceipt Models.ReceiptViewModel

	if err := c.ShouldBindJSON(&newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receiptId, err := _receiptService.ProcessReceipt(&newReceipt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, receiptId)
}

func InitReceiptController(router *gin.Engine, receiptService *Services.ReceiptService) {
	_receiptService = receiptService

	receiptGroup := router.Group("/receipts")
	{
		receiptGroup.POST("/process", CreateReceipt)
	}
}
