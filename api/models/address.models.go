package models

type Address struct {
	ID              uint
	CustomerID      *uint
	Customer        *Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Type            string   `gorm:";NOT NULL;type:enum('legal','branch','contact')"`
	Address         string   `gorm:"NOT NULL"`
	Pobox           string   `gorm:"NOT NULL"`
	PostalCode      string   `gorm:"NOT NULL"`
	City            string   `gorm:"NOT NULL"`
	Province        string   `gorm:"NOT NULL"`
	Country         string   `gorm:"NOT NULL"`
	AttentionPerson *string
}
