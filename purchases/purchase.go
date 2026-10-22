package purchases

import "github.com/google/uuid"

type ItemOrder struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  uint `json:"quantity" validate:"required"`
}

type PurchaseRequest struct {
	Items             []ItemOrder `json:"items" validate:"required,min=1,dive"`
	Courier           string      `json:"courier" validate:"required,validCourier"`
	DestinationID     uint        `json:"destination_id" validate:"required"`
	UserID            uuid.UUID
	ProductID         uint
	ProductName       string
	Weight            uint
	Price             float64 // harga per barang * quantity
	Fee               float64 // ongkir
	Amount            float64 // * price
	Total             float64 // fee + amount
	EstimatedDelivery uint
}

type UpdatePurchaseRequest struct {
	Status string `json:"status" validate:"required"`
}
