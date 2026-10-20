package models

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	PurchaseID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	UserID            uuid.UUID      `json:"user_id" gorm:"type:uuid;not null;index"`
	User              User           `json:"user" gorm:"foreignKey:UserID"`
	ProductID         uint           `json:"product_id"`
	Product           Product        `json:"product"`
	Quantity          uint           `json:"quantity"`
	Courier           string         `json:"courier"`
	DestinationID     uint           `json:"destination_id"`
	TotalWeight       uint           `json:"total_weight"`
	Fee               float64        `json:"fee"`
	Amount            float64        `json:"amount"` // total harga barang pesanan
	Total             float64        `json:"total"`  // amount + fee
	EstimatedDelivery uint           `json:"estimated_delivery"`
	Waybill           *string        `json:"waybill"` // akan diisi worker saat on_delivery
	Status            PurchaseStatus `json:"status" gorm:"type:purchase_status;not null;default:pending;index"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	PaidAt            *time.Time     `json:"paid_at"`
	CancelledAt       *time.Time     `json:"cancelled_at"`
}

type PurchaseStatus string

const (
	Pending    PurchaseStatus = "pending"
	Paid       PurchaseStatus = "paid"
	OnDelivery PurchaseStatus = "on_delivery"
	Cancelled  PurchaseStatus = "cancelled"
)

func (p *PurchaseStatus) Scan(value any) error {
	*p = PurchaseStatus(value.([]byte))
	return nil
}

func (p PurchaseStatus) Value() (driver.Value, error) {
	return string(p), nil
}

func (p *Purchase) BeforeCreate(tx *gorm.DB) error {
	if p.PurchaseID == uuid.Nil {
		p.PurchaseID = uuid.New()
	}

	return nil
}
