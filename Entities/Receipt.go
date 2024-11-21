package Entities

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Receipt struct {
	Id           uuid.UUID
	Retailer     string
	PurchaseDate time.Time
	Total        float64
	ItemIds      []uuid.UUID
	Points       int
}
