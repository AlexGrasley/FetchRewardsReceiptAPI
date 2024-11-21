package Entities

import (
	uuid "github.com/satori/go.uuid"
	"syscall"
)

type Item struct {
	Id          uuid.UUID
	ReceiptId   syscall.GUID
	Description string
	Price       float64
}
