package Repositories

import (
	"FetchRewardsReceiptAPI/Models"
	"FetchRewardsReceiptAPI/Repositories"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddReceipt(t *testing.T) {
	itemRepo := Repositories.NewItemRepository()
	repo := Repositories.NewReceiptRepository(itemRepo)

	receiptViewModel := &Models.ReceiptViewModel{
		Retailer:     "Test Store",
		PurchaseDate: "2021-01-01",
		PurchaseTime: "12:00",
		Items:        make([]*Models.ItemViewModel, 0),
		Total:        "20.99",
	}

	id, err := repo.AddReceipt(receiptViewModel)
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.UUID{}, id)

	receipt := repo.GetReceipt(id)
	assert.Equal(t, "Test Store", receipt.Retailer)
	assert.Equal(t, "2021-01-01 12:00:00 +0000 UTC", receipt.PurchaseDate.String())
	assert.Len(t, receipt.ItemIds, 0)
	assert.Equal(t, 20.99, receipt.Total)
}
