package customer

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

	customerID := Repo.User.IDtoCustomerID(userID)
	if customerID == nil {
		ctx.IndentedJSON(404, gin.H{"error": "user is not registered as master"})
		return
	}

	orders := Repo.Order.GetByCustomerID(int(*customerID))
	ctx.IndentedJSON(200, orders)
}

func GetOrderById(ctx *gin.Context) {
	userID, err := my_jwt.ExtractID(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to extract id"})
		return
	}

	customerID := Repo.User.IDtoCustomerID(userID)
	if customerID == nil {
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

	if order.CustomerID != *customerID {
		ctx.IndentedJSON(403, gin.H{"error": "order customer id doesn't match"})
		return
	}

	ctx.IndentedJSON(200, order)
}
