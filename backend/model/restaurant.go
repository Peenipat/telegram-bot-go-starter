package model

import "time"

type MenuItem struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Category    string
	Price       float64 `gorm:"not null"`
	IsAvailable bool    `gorm:"default:true"`
}

type Order struct {
	ID            uint       `gorm:"primaryKey"`
	OrderTime     time.Time  `gorm:"not null;default:current_timestamp"`
	TotalAmount   float64    `gorm:"not null"`
	TableNumber   *int
	PaymentMethod string
	Status        string     `gorm:"default:completed"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	Feedbacks  []Feedback  `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID         uint    `gorm:"primaryKey"`
	OrderID    uint    `gorm:"not null"`
	MenuItemID uint    `gorm:"not null"`
	Quantity   int     `gorm:"not null"`
	Price      float64 `gorm:"not null"`

	MenuItem MenuItem `gorm:"foreignKey:MenuItemID"`
}

type Reservation struct {
	ID           uint      `gorm:"primaryKey"`
	CustomerName string    `gorm:"not null"`
	Phone        string
	TableNumber  int
	ReservedTime time.Time `gorm:"not null"`
	Status       string    `gorm:"default:pending"`
}

type Feedback struct {
	ID           uint      `gorm:"primaryKey"`
	OrderID      *uint
	CustomerName string
	Comment      string
	Rating       int       `gorm:"check:rating >= 1 AND rating <= 5"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
}

type Ingredient struct {
	ID         uint    `gorm:"primaryKey"`
	Name       string  `gorm:"not null"`
	Unit       string  `gorm:"not null"` // e.g., kg, pcs
	StockLevel float64 `gorm:"default:0"`
}

type StockMovement struct {
	ID           uint      `gorm:"primaryKey"`
	IngredientID uint      `gorm:"not null"`
	Amount       float64   `gorm:"not null"`
	Type         string    `gorm:"not null;check:type IN ('in','out')"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`

	Ingredient Ingredient `gorm:"foreignKey:IngredientID"`
}
