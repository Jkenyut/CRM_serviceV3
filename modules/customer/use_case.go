package customer

import (
	"crm_serviceV3/entity"
	"crm_serviceV3/repository"
)

// interface customer usecase
type UseCaseCustomerInterface interface {
	CreateCustomer(customer CustomerBody) (entity.Customer, error)
	GetCustomerById(id uint) (entity.Customer, error)
	GetAllCustomer(page uint, username string) (uint, uint, int, uint, []entity.Customer, error)
	UpdateCustomerById(id uint, customer UpdateCustomerBody) (entity.Customer, error)
	DeleteCustomerById(id uint) error
}

// struct customer usecase
type customerUseCaseStruct struct {
	customerRepository repository.CustomerRepoInterface
}

// method create customer
func (uc customerUseCaseStruct) CreateCustomer(customer CustomerBody) (entity.Customer, error) {

	NewCustomer := entity.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
	}

	createCustomer, err := uc.customerRepository.CreateCustomer(&NewCustomer)
	if err != nil {
		return NewCustomer, err
	}
	return createCustomer, nil
}

// method get customer by id
func (uc customerUseCaseStruct) GetCustomerById(id uint) (entity.Customer, error) {
	var customer entity.Customer
	customer, err := uc.customerRepository.GetCustomerById(id)
	return customer, err
}

// method get all customer
func (uc customerUseCaseStruct) GetAllCustomer(page uint, username string) (uint, uint, int, uint, []entity.Customer, error) {
	var customer []entity.Customer
	page, perPage, total, totalPages, customer, err := uc.customerRepository.GetAllCustomer(page, username)
	return page, perPage, total, totalPages, customer, err
}

// method update customer by id
func (uc customerUseCaseStruct) UpdateCustomerById(id uint, customer UpdateCustomerBody) (entity.Customer, error) {

	newCustomer := entity.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Avatar:    customer.Avatar,
	}

	updatedCustomer, err := uc.customerRepository.UpdateCustomerById(id, &newCustomer)
	if err != nil {
		return newCustomer, err
	}

	return updatedCustomer, nil
}

// method delete customer by id
func (uc customerUseCaseStruct) DeleteCustomerById(id uint) error {
	err := uc.customerRepository.DeleteCustomerById(id)
	return err
}
