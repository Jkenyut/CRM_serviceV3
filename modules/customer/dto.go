package customer

import (
	"crm_serviceV3/dto"
	"crm_serviceV3/entity"
)

// struct request body customer
type CustomerBody struct {
	FirstName string `json:"firstname" validate:"required,min=1,max=100,alpha"`
	LastName  string `json:"lastname" validate:"min=1,max=100,alpha"`
	Email     string `json:"email" validate:"required,email"`
	Avatar    string `json:"avatar" validate:"min=1,max=250,alphanumunicode"`
}

// struct request update body customer
type UpdateCustomerBody struct {
	FirstName string `json:"firstname" validate:"required,min=1,max=100,alpha"`
	LastName  string `json:"lastname" validate:"min=1,max=100,alpha"`
	Avatar    string `json:"avatar" validate:"min=1,max=250,alphanumunicode"`
}

// struct success
type SuccessCreate struct {
	dto.ResponseMeta
	Data CustomerBody `json:"data"`
}

// struct find customer
type FindCustomer struct {
	dto.ResponseMeta
	Data entity.Customer `json:"data"`
}

// struct find all customer
type FindAllCustomer struct {
	dto.ResponseMeta
	Page       uint              `json:"page,omitempty"`
	PerPage    uint              `json:"per_page,omitempty"`
	Total      int               `json:"total,omitempty"`
	TotalPages uint              `json:"total_pages,omitempty"`
	Data       []entity.Customer `json:"data"`
}
