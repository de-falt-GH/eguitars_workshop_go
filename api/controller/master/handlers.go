package master

import (
	"kursarbeit/api/my_jwt"
	. "kursarbeit/config"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(ctx *gin.Context) {
	userID, err := my_jwt.ExtractID(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to extract id"})
		return
	}

	masterID := Repo.User.IDtoMasterID(userID)
	if masterID == nil {
		ctx.IndentedJSON(403, gin.H{"error": "user is not registered as master"})
		return
	}

	orders := Repo.Order.GetByMasterID(int(*masterID))
	ctx.IndentedJSON(200, orders)
}

func GetOrderById(ctx *gin.Context) {
	userID, err := my_jwt.ExtractID(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to extract id"})
		return
	}

	masterID := Repo.User.IDtoMasterID(userID)
	if masterID == nil {
		ctx.IndentedJSON(404, gin.H{"error": "user is not registered as master"})
		return
	}

	orderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "invalid order id"})
		return
	}

	order := Repo.Order.GetByID(orderID)
	if order == nil {
		ctx.IndentedJSON(404, gin.H{"error": "order not found"})
		return
	}

	if order.MasterID != *masterID {
		ctx.IndentedJSON(403, gin.H{"error": "order master id doesn't match"})
		return
	}

	ctx.IndentedJSON(200, order)
}
