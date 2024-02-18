package admin

import (
	"kursarbeit/api/controller/common"

	"github.com/gin-gonic/gin"
)

func SetRoutes(rg *gin.RouterGroup) {
	admin := rg.Group("/admin").Use(common.Auth(4))

	admin.GET("/users", GetAllUsers)
	admin.GET("/users/:id", GetUserById)

	admin.POST("/users", PostUser)

	admin.PUT("/users/:id", PutUser) // PUT User unembedded fields

	admin.DELETE("/users/:id", DeleteUser)
}
