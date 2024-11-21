package Repositories

import (
	"FetchRewardsReceiptAPI/Entities"
	"FetchRewardsReceiptAPI/Models"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"time"
)

var itemRepo *ItemRepository

type ReceiptRepository struct {
	receipts map[uuid.UUID]Entities.Receipt
}

func NewReceiptRepository(iRepo *ItemRepository) *ReceiptRepository {
	itemRepo = iRepo

	return &ReceiptRepository{
		receipts: make(map[uuid.UUID]Entities.Receipt),
	}
}

func (r *ReceiptRepository) AddReceipt(receiptViewModel *Models.ReceiptViewModel) (uuid.UUID, error) {
	id := uuid.NewV4()
	dateTimeString := receiptViewModel.PurchaseDate + "T" + receiptViewModel.PurchaseTime + ":00Z"

	purchaseDate, err := time.Parse(time.RFC3339, dateTimeString)
	if err != nil {
		return uuid.UUID{}, err
	}

	total, err := strconv.ParseFloat(receiptViewModel.Total, 32)

	if err != nil {
		return uuid.UUID{}, err
	}

	receipt := Entities.Receipt{
		Id:           id,
		Retailer:     receiptViewModel.Retailer,
		PurchaseDate: purchaseDate,
		Total:        total,
		ItemIds:      []uuid.UUID{},
	}

	for _, itemViewModel := range receiptViewModel.Items {
		itemId, err := itemRepo.AddItem(itemViewModel)
		if err != nil {
			return uuid.UUID{}, err
		}
		receipt.ItemIds = append(receipt.ItemIds, itemId)
	}

	r.receipts[id] = receipt

	return id, nil
}
