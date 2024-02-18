package manager

import (
	"kursarbeit/api/controller/common"
	. "kursarbeit/api/controller/manager/customer"
	. "kursarbeit/api/controller/manager/master"
	. "kursarbeit/api/controller/manager/order"

	"github.com/gin-gonic/gin"
)

func SetRoutes(rg *gin.RouterGroup) {
	manager := rg.Group("/manager").Use(common.Auth(3))

	// customer controllers
	manager.GET("/customers", GetCustomers)
	manager.GET("/customers/:id", GetCustomerById)

	manager.POST("/customers", PostCustomer)

	manager.PUT("/customers/:id", PutCustomer) // PUT Customer unembedded(own) fields

	manager.DELETE("/customers/:id", DeleteCustomer)

	// manager.PUT("/customers/:id/personal", PutCustomerPersonalInfo)

	// master controllers
	manager.GET("/masters", GetAllMasters)
	manager.GET("/masters/:id", GetMasterById)

	manager.POST("/masters", PostMaster)

	manager.PUT("/masters/:id", PutMaster) // PUT Master unembedded(own) fields

	manager.DELETE("/masters/:id", DeleteMaster)

	// Order controllers
	manager.GET("/orders", GetAllOrders)
	manager.GET("/orders/:id", GetOrderById)

	manager.POST("/orders", PostOrder)

	manager.PUT("/orders/:id", PutOrder) // PUT Order unembedded(own) fields

	manager.DELETE("/orders/:id", DeleteOrder)

	manager.GET("/components", GetAllComponents)
}
