package workshop

import (
	"fmt"
	"time"
)

type OrderStatus struct {
	ID          uint
	Description string
	Orders      []Order
}

type OrderType struct {
	ID          uint
	Description string
}

type Guitar struct {
	ID           uint
	Name         string
	Condition    string
	SerialNumber string
}

type RequiredComponents struct {
	ID          uint
	OrderID     uint
	ComponentID uint
	Component   Component
}

type Component struct {
	ID           uint
	Type         string
	Manufacturer string
	Name         string
	Quantity     int
}

type Order struct {
	ID            uint
	CustomerID    uint
	Customer      *Customer  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MasterID      uint       `gorm:"foreignKey:ID"`
	Master        *Master    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderTypeID   uint       `gorm:"foreignKey:ID"`
	OrderType     *OrderType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderStatusID uint       `gorm:"foreignKey:ID"`
	// OrderStatus        *OrderStatus         `gorm:"foreignKey:OrderStatusID;default:1;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GuitarID           *uint                `gorm:"foreignKey:ID"`
	Guitar             *Guitar              `gorm:"default:null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RequiredComponents []RequiredComponents `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Description        string
	Price              int
	CreatedAt          time.Time
}

func (order *Order) Insert() error {
	db := repo.GetDB()
	if err := db.Create(order).Error; err != nil {
		return err
	}

	return nil
}

func (order *Order) Update() error {
	fmt.Println(order.OrderStatusID)
	if err := repo.GetDB().Save(order).Error; err != nil {
		return err
	}
	fmt.Println(order.OrderStatusID)

	return nil
}

func (order *Order) Delete() error {
	db := repo.GetDB()
	if err := db.Delete(order).Error; err != nil {
		return err
	}

	return nil
}
