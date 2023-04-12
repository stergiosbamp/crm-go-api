package models

type Address struct {
	ID              uint      `json:"id"`
	CustomerID      *uint     `json:"customerId,omitempty"`
	Customer        *Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
	Type            string    `gorm:";NOT NULL;type:enum('legal','branch','contact')" json:"type"`
	Address         string    `gorm:"NOT NULL" json:"address"`
	Pobox           string    `gorm:"NOT NULL" json:"pobox"`
	PostalCode      string    `gorm:"NOT NULL" json:"postalCode"`
	City            string    `gorm:"NOT NULL" json:"city"`
	Province        string    `gorm:"NOT NULL" json:"province"`
	Country         string    `gorm:"NOT NULL" json:"country"`
	AttentionPerson *string   `json:"attentionPerson,omitempty"`
}
