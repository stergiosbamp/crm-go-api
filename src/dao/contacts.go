package dao

import "github.com/stergiosbamp/go-api/src/models"

type ContactDAO struct {
}

func NewContactDAO() *ContactDAO {
	return &ContactDAO{}
}

func (cdao *ContactDAO) GetList() ([]models.Contact, error) {
	var contacts []models.Contact
	res := db.Find(&contacts)
	return contacts, res.Error

}

func (cdao *ContactDAO) GetById(id uint) (models.Contact, error) {
	var contact models.Contact
	res := db.First(&contact, id)
	return contact, res.Error
}

func (cdao *ContactDAO) Create(contact *models.Contact) (*models.Contact, error) {
	res := db.Create(contact)
	return contact, res.Error
}

func (cdao *ContactDAO) Update(contact *models.Contact) (*models.Contact, error) {
	res := db.Save(contact)
	return contact, res.Error
}

func (cdao *ContactDAO) Delete(id uint) error {
	var contact models.Contact
	// db.Where("id = ?", id).Unscoped().Delete(&contact)
	res := db.Delete(&contact, id).Unscoped()
	return res.Error
}
