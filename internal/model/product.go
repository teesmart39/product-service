package model

type Product struct {
	ID          string  `gorm:"primaryKey"`
	Name        string
	Price       float32
	Description string
}
