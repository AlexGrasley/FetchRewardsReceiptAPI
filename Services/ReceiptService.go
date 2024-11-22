package Services

import (
	"FetchRewardsReceiptAPI/Entities"
	"FetchRewardsReceiptAPI/Models"
	"FetchRewardsReceiptAPI/Repositories"
	uuid "github.com/satori/go.uuid"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ReceiptService struct {
	receiptRepo *Repositories.ReceiptRepository
	itemRepo    *Repositories.ItemRepository
}

func NewReceiptService(repo *Repositories.ReceiptRepository, itemRepo *Repositories.ItemRepository) *ReceiptService {
	return &ReceiptService{
		receiptRepo: repo,
		itemRepo:    itemRepo,
	}
}

func (service *ReceiptService) ProcessReceipt(receipt *Models.ReceiptViewModel) (uuid.UUID, error) {
	receiptId, err := service.receiptRepo.AddReceipt(receipt)
	if err != nil {
		return uuid.UUID{}, err
	}

	service.ProcessReceiptPoints(receiptId)

	return receiptId, err
}

func (service *ReceiptService) GetReceiptPoints(id uuid.UUID) Models.ReceiptPointsViewModel {
	receipt := service.receiptRepo.GetReceipt(id)
	pointsViewModel := Models.ReceiptPointsViewModel{
		Points: strconv.Itoa(receipt.Points),
	}

	return pointsViewModel
}

// ProcessReceiptPoints These rules collectively define how many points should be awarded to a receipt.
// * One point for every alphanumeric character in the retailer name.
// * 50 points if the total is a round dollar amount with no cents.
// * 25 points if the total is a multiple of 0.25.
// * 5 points for every two items on the receipt.
// * If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
// * 6 points if the day in the purchase date is odd.
// * 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func (service *ReceiptService) ProcessReceiptPoints(receiptId uuid.UUID) {
	receipt := service.receiptRepo.GetReceipt(receiptId)
	items := service.itemRepo.GetItems(receipt.ItemIds)

	points := 0

	points += getPointsFromRetailerName(receipt.Retailer)
	points += getPointsFromTotal(receipt.Total)
	points += getPointsFromItems(items)
	points += getPointsFromDate(receipt.PurchaseDate)

	service.receiptRepo.UpdateReceiptPoints(receiptId, points)
}

func processPoints(receipt *Entities.Receipt, items []*Entities.Item) int {
	pointsChannel := make(chan int, 4)

	go func() {
		pointsChannel <- getPointsFromRetailerName(receipt.Retailer)
	}()
	go func() {
		pointsChannel <- getPointsFromTotal(receipt.Total)
	}()
	go func() {
		pointsChannel <- getPointsFromItems(items)
	}()
	go func() {
		pointsChannel <- getPointsFromDate(receipt.PurchaseDate)
	}()

	points := 0
	for i := 0; i < 4; i++ {
		points += <-pointsChannel
	}

	return points
}

func getPointsFromRetailerName(str string) int {
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	result := re.ReplaceAllString(str, "")

	return len(result)
}

func getPointsFromTotal(total float64) int {
	points := 0

	if isRoundDollarAmount(total) {
		points += 50
	}

	if isMultipleOfPoint25(total) {
		points += 25
	}

	return points
}

func getPointsFromItems(items []*Entities.Item) int {
	points := 0
	points += len(items) / 2 * 5

	for _, item := range items {
		descriptLen := len(strings.TrimSpace(item.Description))
		if (descriptLen % 3) == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	return points
}

func getPointsFromDate(date time.Time) int {
	points := 0

	if date.Day()%2 != 0 {
		points += 6
	}

	if date.Hour() >= 14 && date.Hour() < 16 {
		points += 10
	}

	return points
}

func isRoundDollarAmount(value float64) bool {
	return value == math.Floor(value)
}

func isMultipleOfPoint25(value float64) bool {
	return (int(value*100) % 25) == 0
}
