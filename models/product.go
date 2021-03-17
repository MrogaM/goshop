package models

type Buyable interface {
	Generate(...interface{})
}

type Product struct {
	ID        int32
	Name      string
	Category  string
	Price     float32
	Sku       string
	Quantity  int32
	Available bool
	Shipment  Shipment
}