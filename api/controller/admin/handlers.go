package admin

import (
	. "kursarbeit/config"
	. "kursarbeit/storage/models/user"

	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(ctx *gin.Context) {
	users := Repo.User.GetAll(ctx)

	ctx.IndentedJSON(200, users)
}

func GetUserById(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || userID < 1 {
		ctx.IndentedJSON(400, gin.H{"error": "invalid id"})
		return
	}

	user := Repo.User.GetByID(userID)
	if user == nil {
		ctx.IndentedJSON(404, gin.H{"error": "user not found"})
		return
	}

	ctx.IndentedJSON(200, user)
}

func PostUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to parse request body"})
		return
	}

	hashedPassword := []byte(user.Credentials.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(hashedPassword, 15)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to hash password"})
		return
	}

	user.Credentials.Password = string(hashedPassword)

	if err := user.Insert(); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to update received record: " + err.Error()})
		return
	}
}

func PutUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || userID < 1 {
		ctx.IndentedJSON(400, gin.H{"error": "invalid id"})
		return
	}

	var userReq User
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to parse request body"})
		return
	}

	user := Repo.User.GetByID(userID)
	if user == nil {
		ctx.IndentedJSON(404, gin.H{"error": "requested user not found"})
		return
	}

	userChanged := false
	if userReq.Name != "" {
		user.Name = userReq.Name
		userChanged = true
	}

	if userReq.Role != 0 {
		user.Role = userReq.Role
		userChanged = true
	}

	if userReq.Credentials != nil {
		if userReq.Credentials.Login != "" {
			user.Credentials.Login = userReq.Credentials.Login
		}
		if userReq.Credentials.Password != "" {
			newPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Credentials.Password), 15)
			if err != nil {
				ctx.IndentedJSON(400, gin.H{"error": "unable to store password"})
				return
			}
			user.Credentials.Password = string(newPassword)
		}
		userChanged = true
	}

	if userChanged {
		if err := user.Update(); err != nil {
			ctx.IndentedJSON(400, gin.H{"error": "unable to update received record: " + err.Error()})
		}
	}
}

func DeleteUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || userID < 1 {
		ctx.IndentedJSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := Repo.User.DeleteByID(userID); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
	}
}
