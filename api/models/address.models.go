package models

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
