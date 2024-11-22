package Models

import (
	uuid "github.com/satori/go.uuid"
	_ "time"
)

type ReceiptViewModel struct {
	Retailer     string           `json:"retailer"`
	PurchaseDate string           `json:"purchaseDate"`
	PurchaseTime string           `json:"purchaseTime"`
	Total        string           `json:"total"`
	Items        []*ItemViewModel `json:"items"`
}

type ReceiptIdViewModel struct {
	Id uuid.UUID `json:"id"`
}

type ReceiptPointsViewModel struct {
	Points string `json:"points"`
}
