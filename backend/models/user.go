package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Role     string `gorm:"role"`

	Patients []Patient `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
