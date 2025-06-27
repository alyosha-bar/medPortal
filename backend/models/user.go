package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Role     string `gorm:"role"`
}
