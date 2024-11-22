package Controllers

import (
	"FetchRewardsReceiptAPI/Models"
	"FetchRewardsReceiptAPI/Services"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

var _receiptService *Services.ReceiptService

func GetReceiptPoints(c *gin.Context) {
	receiptId := c.Param("id")

	id, err := uuid.FromString(receiptId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	points := _receiptService.GetReceiptPoints(id)
	c.JSON(http.StatusOK, points)
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
		receiptGroup.GET("/:id/points", GetReceiptPoints)
	}
}
