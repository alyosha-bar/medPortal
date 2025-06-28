package models

type Patient struct {
	ID        uint   `gorm:"primaryKey"`
	Firstname string `gorm:"firstname"`
	Lastname  string `gorm:"lastname"`
	Age       uint   `gorm:"age"`
	Gender    string `gorm:"gender"`

	DoctorID *uint // Use *uint for nullable foreign key (if a patient can exist without an assigned doctor)
	// If a patient MUST have a doctor, use uint (non-nullable) and remove 'OnDelete:SET NULL' if you want RESTRICT
	Doctor User `gorm:"foreignKey:DoctorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	MedicalNotes string
}
