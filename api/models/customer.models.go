package models

type Customer struct {
	ID                 uint   `json:"id"`
	Active             bool   `gorm:"NOT NULL" json:"active"`
	Name               string `gorm:"NOT NULL" json:"name"`
	ShortName          string `gorm:"NOT NULL" json:"shortName"`
	VatNumber          string `gorm:"NOT NULL" json:"vatNumber"`
	VatApply           bool   `gorm:"NOT NULL" json:"vatApply"`
	RegistrationNumber string `gorm:"NOT NULL" json:"registrationNumber"`
	DunsNumber         string `gorm:"NOT NULL" json:"dunsNumber"`
	TaxExempt          bool   `gorm:"NOT NULL" json:"taxExempt"`
	Language           string `gorm:"size:2;NOT NULL" json:"language"`
	Email              string `gorm:"NOT NULL" json:"email"`
	Phone              string `gorm:"NOT NULL" json:"phone"`
	Fax                string `gorm:"NOT NULL" json:"fax"`

	/*
		The following slices are not needed to model the the DB as it is originally.
	*/

	// Addresses 		   []Address `gorm:"NOT NULL"`
	// Contacts           []Contact `gorm:"NOT NULL"`
}
