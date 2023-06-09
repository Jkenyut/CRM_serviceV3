package customer

import (
	"crm_serviceV3/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// struct route customer
type RouterCustomerStruct struct {
	customerRequestHandler RequestHandlerCustomerStruct
}

// func new router
func NewRouter(
	dbCrud *gorm.DB,
) RouterCustomerStruct {
	return RouterCustomerStruct{
		customerRequestHandler: RequestHandler(
			dbCrud,
		),
	}
}

// func handle route customer
func (r RouterCustomerStruct) Handle(router *gin.Engine) {
	basepath := "v1/customer"
	customerRouter := router.Group(basepath, middleware.Auth)

	customerRouter.POST("/register",
		r.customerRequestHandler.CreateCustomer,
	)

	customerRouter.GET("/:id", middleware.CustomerBulk,
		r.customerRequestHandler.GetCustomerById,
	)
	customerRouter.GET("", middleware.CustomerBulk,
		r.customerRequestHandler.GetAllCustomer,
	)

	customerRouter.PUT("/:id",
		r.customerRequestHandler.UpdateCustomerById,
	)
	customerRouter.DELETE("/:id",
		r.customerRequestHandler.DeleteCustomerById,
	)
}
