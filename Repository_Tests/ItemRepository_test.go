package Repositories

import (
	"FetchRewardsReceiptAPI/Models"
	"FetchRewardsReceiptAPI/Repositories"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddItem(t *testing.T) {
	repo := Repositories.NewItemRepository()

	itemViewModel := &Models.ItemViewModel{
		Description: "Test Item",
		Price:       "10.99",
	}

	id, err := repo.AddItem(itemViewModel)
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.UUID{}, id)

	items := repo.GetItems([]uuid.UUID{id})
	item := items[0]
	assert.Equal(t, "Test Item", item.Description)
	assert.Equal(t, 10.99, item.Price)
}
