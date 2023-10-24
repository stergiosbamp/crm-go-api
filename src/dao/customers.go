package dao

import "github.com/stergiosbamp/go-api/src/models"

type CustomerDAO struct {
}

func NewCustomerDAO() *CustomerDAO {
	return &CustomerDAO{}
}

func (cdao *CustomerDAO) GetList() ([]models.Customer, error) {
	var customers []models.Customer
	res := db.Find(&customers)
	return customers, res.Error

}

func (cdao *CustomerDAO) GetById(id uint) (models.Customer, error) {
	var customer models.Customer
	res := db.First(&customer, id)
	return customer, res.Error
}

func (cdao *CustomerDAO) Create(customer *models.Customer) (*models.Customer, error) {
	res := db.Create(customer)
	return customer, res.Error
}

func (cdao *CustomerDAO) Update(customer *models.Customer) (*models.Customer, error) {
	res := db.Save(customer)
	return customer, res.Error
}

func (cdao *CustomerDAO) Delete(id uint) error {
	var customer models.Customer
	// db.Where("id = ?", id).Unscoped().Delete(&customer)
	res := db.Delete(&customer, id).Unscoped()
	return res.Error
}
