package models

import (
	"time"
)

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

type Contact struct {
	ID          uint     
	ContactType string `gorm:"column:contact_type;NOT NULL;type:enum('Commercial','Finance','Legal','CEO','Other','DPO','Technical')"`
	FirstName   string   `gorm:"NOT NULL"`
	LastName    string   `gorm:"NOT NULL"`
	NickName    string   `gorm:"NOT NULL"`
	Gender      string   `gorm:"NOT NULL"`
	Birthday    time.Time`gorm:"NOT NULL"`
	Language    string   `gorm:"size:2;NOT NULL"`
	JobTitle    string   `gorm:"NOT NULL"`
	Email       string   `gorm:"NOT NULL"`
	Skype       string   `gorm:"NOT NULL"`
	PhoneDirect string   `gorm:"NOT NULL"`
	PhoneOffice string   `gorm:"NOT NULL"`
	Mobile      string   `gorm:"NOT NULL"`
	Notes       string   `gorm:"NOT NULL"`
	CustomerID  uint     `gorm:"NOT NULL"`
	Customer 	Customer `gorm:"NOT NULL"`
	AddressID   uint    
	Address 	Address 
}

type Address struct {
	ID              uint    
	CustomerID      uint	`gorm:"NOT NULL"`
	Customer 		Customer`gorm:"NOT NULL"`
	Type            string 	`gorm:";NOT NULL;type:enum('legal','branch','contact')"`
	Address         string	`gorm:"NOT NULL"`
	Pobox           string 	`gorm:"NOT NULL"`
	PostalCode      string 	`gorm:"NOT NULL"`
	City            string	`gorm:"NOT NULL"`
	Province        string	`gorm:"NOT NULL"`
	Country         string	`gorm:"NOT NULL"`
	AttentionPerson string	
}
