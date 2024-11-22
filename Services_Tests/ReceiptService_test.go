package Services_Tests

import (
	"FetchRewardsReceiptAPI/Entities"
	"FetchRewardsReceiptAPI/Models"
	"FetchRewardsReceiptAPI/Repositories"
	"FetchRewardsReceiptAPI/Services"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	itemRepo    *Repositories.ItemRepository
	receiptRepo *Repositories.ReceiptRepository
	service     *Services.ReceiptService
)

func TestMain(m *testing.M) {
	// Setup code
	itemRepo = Repositories.NewItemRepository()
	receiptRepo = Repositories.NewReceiptRepository(itemRepo)
	service = Services.NewReceiptService(receiptRepo, itemRepo)

	// Run tests
	code := m.Run()

	// Exit with the result of m.Run()
	os.Exit(code)
}

func TestProcessReceipt(t *testing.T) {
	receiptViewModel := &Models.ReceiptViewModel{
		Retailer:     "Test Store",
		PurchaseDate: "2021-01-01",
		PurchaseTime: "12:00",
		Items:        make([]*Models.ItemViewModel, 0),
		Total:        "20.99",
	}

	id, err := service.ProcessReceipt(receiptViewModel)
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.UUID{}, id)
}

func TestGetReceiptPoints(t *testing.T) {
	receiptViewModel := &Models.ReceiptViewModel{
		Retailer:     "Test Store",
		PurchaseDate: "2021-01-01",
		PurchaseTime: "12:00",
		Items:        make([]*Models.ItemViewModel, 0),
		Total:        "20.99",
	}

	id, _ := service.ProcessReceipt(receiptViewModel)
	pointsViewModel := service.GetReceiptPoints(id)

	assert.Equal(t, "15", pointsViewModel.Points)
}

func TestGetPointsFromRetailerName(t *testing.T) {
	points := Services.GetPointsFromRetailerName("Test Store")
	assert.Equal(t, 9, points)
}

func TestGetPointsFromTotal(t *testing.T) {
	points := Services.GetPointsFromTotal(25)
	assert.Equal(t, 75, points)
}

func TestGetPointsFromItems(t *testing.T) {
	// 0 points
	var item0 = &Entities.Item{
		Description: "Test Item 0",
		Price:       10.99,
	}

	// 3 points
	var item1 = &Entities.Item{
		Description: "ABC",
		Price:       14.99,
	}

	// 0 points
	var item2 = &Entities.Item{
		Description: "Test Item 2",
		Price:       18.99,
	}

	// 5 points, length is 3.
	items := []*Entities.Item{item0, item1, item2}

	points := Services.GetPointsFromItems(items)

	assert.Equal(t, 8, points)
}

func TestGetPointsFromDate(t *testing.T) {
	// 6 points, date is odd
	date1, _ := time.Parse(time.RFC3339, "2021-01-01T12:00:00Z")

	// 0 points, date is even
	date2, _ := time.Parse(time.RFC3339, "2021-01-02T11:00:00Z")

	// 10 points, time is between 2 and 4 pm
	date3, _ := time.Parse(time.RFC3339, "2021-01-02T15:00:00Z")

	// 0 points, time is not between 2 and 4 pm
	date4, _ := time.Parse(time.RFC3339, "2021-01-04T11:00:00Z")

	//16 points, date is odd and time is between 2 and 4 pm
	date5, _ := time.Parse(time.RFC3339, "2021-01-05T15:00:00Z")

	points1 := Services.GetPointsFromDate(date1)
	points2 := Services.GetPointsFromDate(date2)
	points3 := Services.GetPointsFromDate(date3)
	points4 := Services.GetPointsFromDate(date4)
	points5 := Services.GetPointsFromDate(date5)

	assert.Equal(t, 6, points1)
	assert.Equal(t, 0, points2)
	assert.Equal(t, 10, points3)
	assert.Equal(t, 0, points4)
	assert.Equal(t, 16, points5)
}
