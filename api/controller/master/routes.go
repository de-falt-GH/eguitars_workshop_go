package master

import (
	"kursarbeit/api/controller/common"

	"github.com/gin-gonic/gin"
)

func SetRoutes(rg *gin.RouterGroup) {
	master := rg.Group("/master").Use(common.Auth(2))

	master.GET("/orders", GetAllOrders)
	master.GET("/orders/:id", GetOrderById)
}
