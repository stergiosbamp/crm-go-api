package dao

import (
	"github.com/stergiosbamp/go-api/models"
)

type AddressDAO struct {
	
}

func NewAddressDAO() *AddressDAO {
	return &AddressDAO{}
}

func (adao *AddressDAO) GetList() ([]models.Address, error) {
	var addresses []models.Address
	res := db.Find(&addresses)
	return addresses, res.Error

}

func (adao *AddressDAO) GetById(id uint) (models.Address, error) {
	var address models.Address
	res := db.First(&address, id)
	return address, res.Error
}

func (adao *AddressDAO) Create(address *models.Address) (*models.Address, error) {
	res := db.Create(address)
	return address, res.Error
}

func (adao *AddressDAO) Update(address *models.Address) (*models.Address, error) {
	res := db.Save(address)
	return address, res.Error
}

func (adao *AddressDAO) Delete(id uint) (error) {
	var address models.Address
	// db.Where("id = ?", id).Unscoped().Delete(&Address)
	res := db.Delete(&address, id).Unscoped()
	return res.Error
}

func (adao *AddressDAO) FindAddress(customerId uint, addressType string) (*models.Address, error) {
	var address models.Address
	res := db.Where("customer_id = ? AND type = ?", customerId, addressType).First(&address)
	return &address, res.Error
}
