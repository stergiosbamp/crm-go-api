package models

type Customer struct {
	ID                 uint  
	Active             bool   	 `gorm:"NOT NULL"`
	Name               string 	 `gorm:"NOT NULL"`
	ShortName          string	 `gorm:"NOT NULL"`
	VatNumber          string	 `gorm:"NOT NULL"`
	VatApply           bool	 	 `gorm:"NOT NULL"`
	RegistrationNumber string	 `gorm:"NOT NULL"`
	DunsNumber         string	 `gorm:"NOT NULL"`
	TaxExempt          bool   	 `gorm:"NOT NULL"`
	Language           string	 `gorm:"size:2;NOT NULL"`
	Email              string	 `gorm:"NOT NULL"`
	Phone              string	 `gorm:"NOT NULL"`
	Fax                string	 `gorm:"NOT NULL"`
	
	/*
	The following slices are not needed to model the the DB as it is originally.
	*/
	
	// Addresses 		   []Address `gorm:"NOT NULL"`
	// Contacts           []Contact `gorm:"NOT NULL"`
}
