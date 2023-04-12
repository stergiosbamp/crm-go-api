package models

type Contact struct {
	ID          uint      `json:"id"`
	ContactType string    `gorm:"column:contact_type;NOT NULL;type:enum('Commercial','Finance','Legal','CEO','Other','DPO','Technical')" json:"contactType"`
	FirstName   string    `gorm:"NOT NULL" json:"firstName"`
	LastName    string    `gorm:"NOT NULL" json:"lastName"`
	NickName    string    `gorm:"NOT NULL" json:"nickName"`
	Gender      string    `gorm:"NOT NULL" json:"gender"`
	Birthday    string    `gorm:"NOT NULL; type:date" json:"birthday"`
	Language    string    `gorm:"size:2;NOT NULL" json:"language"`
	JobTitle    string    `gorm:"NOT NULL" json:"jobTitle"`
	Email       string    `gorm:"NOT NULL" json:"email"`
	Skype       string    `gorm:"NOT NULL" json:"skype"`
	PhoneDirect string    `gorm:"NOT NULL" json:"phoneDirect"`
	PhoneOffice string    `gorm:"NOT NULL" json:"phoneOffice"`
	Mobile      string    `gorm:"NOT NULL" json:"mobile"`
	Notes       string    `gorm:"NOT NULL" json:"notes"`
	CustomerID  uint      `gorm:"NOT NULL" json:"customerId"`
	Customer    *Customer `gorm:"NOT NULL;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
	AddressID   *uint     `json:"addressId,omitempty"`
	Address     *Address  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
}
