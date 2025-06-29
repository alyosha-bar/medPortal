package models

type Patient struct {
	ID        uint   `gorm:"primaryKey"`
	Firstname string `gorm:"firstname"`
	Lastname  string `gorm:"lastname"`
	Age       uint   `gorm:"age"`
	Gender    string `gorm:"gender"`

	DoctorID *uint
	Doctor   User `gorm:"foreignKey:DoctorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	MedicalNotes string
}
