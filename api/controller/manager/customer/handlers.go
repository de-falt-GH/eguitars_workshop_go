package manager

import (
	. "kursarbeit/config"
	. "kursarbeit/storage/models/workshop"
	"kursarbeit/utils/validator"
	"net/mail"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCustomers(ctx *gin.Context) {
	orders := Repo.Customer.GetAll(ctx)

	ctx.IndentedJSON(200, orders)
}

func GetCustomerById(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || ID < 1 {
		ctx.IndentedJSON(400, gin.H{"error": "invalid customer id"})
		return
	}

	customer := Repo.Customer.GetByID(ID)
	if customer == nil {
		ctx.IndentedJSON(404, gin.H{"error": "customer not found"})
		return
	}

	ctx.IndentedJSON(200, customer)
}

func PostCustomer(ctx *gin.Context) {
	var customer Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to parse request body"})
		return
	}

	if err := customer.Insert(); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to insert received record"})
		return
	}

	ctx.IndentedJSON(200, gin.H{"message": "customer posted successfully"})
}

func PutCustomer(ctx *gin.Context) {
	customerID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || customerID < 1 {
		ctx.IndentedJSON(400, gin.H{"error": "invalid id"})
		return
	}

	var customerReq Customer
	if err := ctx.ShouldBindJSON(&customerReq); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to parse request body"})
		return
	}

	customer := Repo.Customer.GetByID(customerID)
	if customer == nil {
		ctx.IndentedJSON(404, gin.H{"error": "requested user not found"})
		return
	}

	customerChanged := false
	if customerReq.Notes != "" {
		customer.Notes = customerReq.Notes
		customerChanged = true
	}

	if customerReq.PersonalInfo != nil {
		if customerReq.PersonalInfo.Name != "" {
			customer.PersonalInfo.Name = customerReq.PersonalInfo.Name
			customerChanged = true
		}
		if customerReq.PersonalInfo.Email != "" {
			_, err := mail.ParseAddress(customerReq.PersonalInfo.Email)
			if err != nil {
				ctx.IndentedJSON(400, gin.H{"error": "invalid email format"})
				return
			}
			customer.PersonalInfo.Email = customerReq.PersonalInfo.Email
			customerChanged = true
		}
		if customerReq.PersonalInfo.PhoneNumber != "" {
			if validator.CheckPhoneNumber(customerReq.PersonalInfo.PhoneNumber) {
				customer.PersonalInfo.PhoneNumber = customerReq.PersonalInfo.PhoneNumber
				customerChanged = true
			} else {
				ctx.IndentedJSON(400, gin.H{"error": "invalid phone format"})
				return
			}
		}
	}

	if customerChanged {
		if err := customer.Update(); err != nil {
			ctx.IndentedJSON(400, gin.H{"error": "unable to update received record: " + err.Error()})
		}
	}

}

func DeleteCustomer(ctx *gin.Context) {
	customerID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || customerID < 1 {
		ctx.IndentedJSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := Repo.Customer.DeleteByID(customerID); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
	}
}
