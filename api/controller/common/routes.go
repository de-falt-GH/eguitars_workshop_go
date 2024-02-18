package common

import "github.com/gin-gonic/gin"

func SetRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", PostAuthorize)
}
