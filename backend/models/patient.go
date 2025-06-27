package models

type Patient struct {
	ID        uint   `gorm:"primaryKey"`
	Firstname string `gorm:"firstname"`
	Lastname  string `gorm:"lastname"`
	Age       uint   `gorm:"age"`
	Gender    string `gorm:"gender"`
}
