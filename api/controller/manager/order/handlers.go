package manager

import (
	. "kursarbeit/config"
	. "kursarbeit/storage/models/workshop"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllComponents(ctx *gin.Context) {
	components := Repo.Order.GetAllComponents()

	ctx.IndentedJSON(200, components)
}

func GetAllOrders(ctx *gin.Context) {
	orders := Repo.Order.GetAll(ctx)

	ctx.IndentedJSON(200, orders)
}

func GetOrderById(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || orderID < 1 {
		ctx.IndentedJSON(400, gin.H{"error": "invalid id"})
		return
	}

	order := Repo.Order.GetByID(orderID)
	if order == nil {
		ctx.IndentedJSON(404, gin.H{"error": "order not found"})
		return
	}

	ctx.IndentedJSON(200, order)
}

func PostOrder(ctx *gin.Context) {
	var order Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to parse request body"})
		return
	}

	if err := order.Insert(); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to insert received record"})
		return
	}
}

func PutOrder(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || orderID < 1 {
		ctx.IndentedJSON(400, gin.H{"error": "invalid id"})
		return
	}

	var orderReq Order
	if err := ctx.ShouldBindJSON(&orderReq); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to parse request body"})
		return
	}

	order := Repo.Order.GetByID(orderID)
	if order == nil {
		ctx.IndentedJSON(404, gin.H{"error": "requested order not found"})
		return
	}

	orderChanged := false
	if orderReq.MasterID != 0 {
		order.MasterID = orderReq.MasterID
		orderChanged = true
	}
	if orderReq.OrderStatusID != 0 {
		order.OrderStatusID = orderReq.OrderStatusID
		orderChanged = true
	}

	if !orderChanged {
		ctx.IndentedJSON(400, gin.H{"error": "order wasn't updated"})
		return
	}

	order.Update()
}

func DeleteOrder(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || orderID < 1 {
		ctx.IndentedJSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := Repo.Order.DeleteByID(orderID); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
	}
}
