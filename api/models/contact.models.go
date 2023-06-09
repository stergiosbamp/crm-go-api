package models

type Contact struct {
	ID          uint
	ContactType string   `gorm:"column:contact_type;NOT NULL;type:enum('Commercial','Finance','Legal','CEO','Other','DPO','Technical')"`
	FirstName   string   `gorm:"NOT NULL"`
	LastName    string   `gorm:"NOT NULL"`
	NickName    string   `gorm:"NOT NULL"`
	Gender      string   `gorm:"NOT NULL"`
	Birthday    string   `gorm:"NOT NULL; type:date"`
	Language    string   `gorm:"size:2;NOT NULL"`
	JobTitle    string   `gorm:"NOT NULL"`
	Email       string   `gorm:"NOT NULL"`
	Skype       string   `gorm:"NOT NULL"`
	PhoneDirect string   `gorm:"NOT NULL"`
	PhoneOffice string   `gorm:"NOT NULL"`
	Mobile      string   `gorm:"NOT NULL"`
	Notes       string   `gorm:"NOT NULL"`
	CustomerID  uint     `gorm:"NOT NULL"`
	Customer    Customer `gorm:"NOT NULL;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AddressID   *uint
	Address     *Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
