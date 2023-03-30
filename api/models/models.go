package models

import (
	"time"
)

type Customer struct {
	ID                 uint   `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Active             int    `gorm:"column:active;NOT NULL"`
	Name               string `gorm:"column:name;NOT NULL"`
	ShortName          string `gorm:"column:short_name;NOT NULL"`
	VatNumber          string `gorm:"column:vat_number;NOT NULL"`
	VatApply           string `gorm:"column:vat_apply;NOT NULL"`
	RegistrationNumber string `gorm:"column:registration_number;NOT NULL"`
	DunsNumber         string `gorm:"column:duns_number;NOT NULL"`
	TaxExempt          int    `gorm:"column:tax_exempt;NOT NULL"`
	Language           string `gorm:"column:language;NOT NULL"`
	Email              string `gorm:"column:email;NOT NULL"`
	Phone              string `gorm:"column:phone;NOT NULL"`
	Fax                string `gorm:"column:fax;NOT NULL"`
}

type Contacts struct {
	ID          uint      `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	// ContactType UnSupport `gorm:"column:contact_type;NOT NULL"`
	FirstName   string    `gorm:"column:first_name;NOT NULL"`
	LastName    string    `gorm:"column:last_name;NOT NULL"`
	NickName    string    `gorm:"column:nick_name;NOT NULL"`
	Gender      string    `gorm:"column:gender;NOT NULL"`
	Birthday    time.Time `gorm:"column:birthday;NOT NULL"`
	Language    string    `gorm:"column:language;NOT NULL"`
	JobTitle    string    `gorm:"column:job_title;NOT NULL"`
	Email       string    `gorm:"column:email;NOT NULL"`
	Skype       string    `gorm:"column:skype;NOT NULL"`
	PhoneDirect string    `gorm:"column:phone_direct;NOT NULL"`
	PhoneOffice string    `gorm:"column:phone_office;NOT NULL"`
	Mobile      string    `gorm:"column:mobile;NOT NULL"`
	Notes       string    `gorm:"column:notes;NOT NULL"`
	CustomerID  uint      `gorm:"column:customer_id;NOT NULL"`
	AddressID   uint      `gorm:"column:address_id"`
}

type Address struct {
	ID              uint      `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CustomerID      uint      `gorm:"column:customer_id"`
	// Type            UnSupport `gorm:"column:type;NOT NULL"`
	Address         string    `gorm:"column:address;NOT NULL"`
	Pobox           int64     `gorm:"column:pobox;NOT NULL"`
	PostalCode      int64     `gorm:"column:postal_code;NOT NULL"`
	City            string    `gorm:"column:city;NOT NULL"`
	Province        string    `gorm:"column:province;NOT NULL"`
	Country         string    `gorm:"column:country;NOT NULL"`
	AttentionPerson string    `gorm:"column:attention_person"`
}
