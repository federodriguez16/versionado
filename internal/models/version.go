package models

type Versions struct {
	ID       uint `gorm:"primaryKey"`
	ProductA string
	ProductB string
	ProductC string
	ProductD string
	ProductE string
}
