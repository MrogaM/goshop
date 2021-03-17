package models

type ShipmentType struct {
	ID          int32
	Name        string
	PaymentType PaymentType
}

type Shipment struct {
	ID   int32
	Type ShipmentType
}
