package controller

import "fazTudo-API/models"

type CustomerRepository interface {
	FindAll() []*models.Customer
	Find(id int) *models.Customer
	Create(name string) *models.Customer
	Update(id int, name string) *models.Customer
	Delete(id int) bool
}
