package actor

import (
	"crm_serviceV3/dto"
	"crm_serviceV3/entity"
	"github.com/dgrijalva/jwt-go"
)

// struct to request body actor
type ActorBody struct {
	Username string `json:"username" validate:"required,min=1,max=100,alphanum"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=100"`
}

// struct to request update body actor
type UpdateActorBody struct {
	Username string `json:"username" validate:"min=1,max=100,alphanum"`
	Password string `json:"password,omitempty" validate:"min=6,max=100"`
	Verified string `json:"verified" validate:"eq=true|eq=false"`
	Active   string `json:"active" validate:"eq=true|eq=false"`
}

// strict successCreate
type SuccessCreate struct {
	dto.ResponseMeta
	Data ActorBody `json:"data"`
}

// struct FindActor
type FindActor struct {
	dto.ResponseMeta
	Data entity.Actor `json:"data"`
}

// strict find actor all
type FindAllActor struct {
	dto.ResponseMeta
	Page       uint           `json:"page,omitempty"`
	PerPage    uint           `json:"per_page,omitempty"`
	Total      int            `json:"total,omitempty"`
	TotalPages uint           `json:"total_pages,omitempty"`
	Data       []entity.Actor `json:"data"`
}

// struct jwt claims
type CustomClaims struct {
	Role      uint   `json:"role"`
	UserAgent string `json:"user_agent"`
	jwt.StandardClaims
}

// struct succes login
type SuccessLogin struct {
	dto.ResponseMeta
	Data string `json:"data"`
}
