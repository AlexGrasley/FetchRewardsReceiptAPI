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

	purchaseDate, err := GetDateFromString(receiptViewModel.PurchaseDate, receiptViewModel.PurchaseTime)
	if err != nil {
		return uuid.UUID{}, err
	}

	total, err := strconv.ParseFloat(receiptViewModel.Total, 64)
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

	itemIds, err := createAndStoreItems(receiptViewModel)
	if err != nil {
		return uuid.UUID{}, err
	}

	receipt.ItemIds = itemIds

	r.receipts[id] = receipt

	return id, nil
}

func (r *ReceiptRepository) GetReceipt(id uuid.UUID) Entities.Receipt {
	return r.receipts[id]
}

func (r *ReceiptRepository) UpdateReceiptPoints(id uuid.UUID, points int) {
	receipt := r.receipts[id]
	receipt.Points = points
	r.receipts[id] = receipt
}

func createAndStoreItems(receiptViewModel *Models.ReceiptViewModel) ([]uuid.UUID, error) {
	itemIds := make([]uuid.UUID, 0, len(receiptViewModel.Items))

	for _, itemViewModel := range receiptViewModel.Items {
		itemId, err := itemRepo.AddItem(itemViewModel)
		if err != nil {
			return []uuid.UUID{}, err
		}
		itemIds = append(itemIds, itemId)
	}

	return itemIds, nil
}

func GetDateFromString(dateString string, timeString string) (time.Time, error) {
	dateTimeString := dateString + "T" + timeString + ":00Z"

	date, err := time.Parse(time.RFC3339, dateTimeString)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}
