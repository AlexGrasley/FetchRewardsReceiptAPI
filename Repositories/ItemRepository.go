package Repositories

import (
	"FetchRewardsReceiptAPI/Entities"
	"FetchRewardsReceiptAPI/Models"
	uuid "github.com/satori/go.uuid"
	"strconv"
)

type ItemRepository struct {
	items map[uuid.UUID]Entities.Item
}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{
		items: make(map[uuid.UUID]Entities.Item),
	}
}

// AddItem Adds a new item to storage. Currently, does not consider duplicates.
// Future enhancements could prevent storing duplicates, but need to consider
// variations in item price between retailers or over time.
func (r *ItemRepository) AddItem(itemViewModel Models.ItemViewModel) (uuid.UUID, error) {
	id := uuid.NewV4()

	price, err := strconv.ParseFloat(itemViewModel.Price, 32)

	if err != nil {
		return uuid.UUID{}, err
	}

	item := Entities.Item{
		Id:          id,
		Description: itemViewModel.Description,
		Price:       price,
	}

	r.items[id] = item
	return id, nil
}

func (r *ItemRepository) GetItems(ids []uuid.UUID) []Entities.Item {
	items := make([]Entities.Item, 0, len(ids))

	for _, id := range ids {
		if item, ok := r.items[id]; ok {
			items = append(items, item)
		}
	}

	return items
}
