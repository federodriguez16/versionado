package models

type Versions struct {
	ID       uint `gorm:"primaryKey"`
	productA string
	productB string
	productC string
	productD string
	productE string
}
