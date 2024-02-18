package customer

import (
	"kursarbeit/api/controller/common"

	"github.com/gin-gonic/gin"
)

func SetRoutes(rg *gin.RouterGroup) {
	customer := rg.Group("/customer").Use(common.Auth(1))

	customer.GET("/orders", GetAllOrders)
	customer.GET("/orders/:id", GetOrderById)
}
