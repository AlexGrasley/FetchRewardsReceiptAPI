package Services

import (
	"FetchRewardsReceiptAPI/Models"
	"FetchRewardsReceiptAPI/Repositories"
	uuid "github.com/satori/go.uuid"
)

type ReceiptService struct {
	receiptRepo *Repositories.ReceiptRepository
}

func NewReceiptService(repo *Repositories.ReceiptRepository) *ReceiptService {
	return &ReceiptService{
		receiptRepo: repo,
	}
}

func (service *ReceiptService) ProcessReceipt(receipt *Models.ReceiptViewModel) (uuid.UUID, error) {

	receiptId, err := service.receiptRepo.AddReceipt(receipt)

	if err != nil {
		return uuid.UUID{}, err
	}

	return receiptId, err
}
