package common

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"kursarbeit/api/my_jwt"
	. "kursarbeit/config"
	. "kursarbeit/storage/models/user"
)

func PostAuthorize(ctx *gin.Context) {
	var creds Credentials
	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid format"})
		return
	}

	user := Repo.User.GetByLogin(creds.Login)
	fmt.Println(user.Credentials.Login, user.Credentials.Password)
	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Credentials.Password), []byte(creds.Password)) != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = int64(time.Now().Add(720 * time.Hour).Unix()) // week
	claims["uid"] = user.ID
	tokenString, err := token.SignedString(my_jwt.Salt)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to generate token string: " + err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"token": tokenString})
}

func Auth(requiredRole UserRole) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.IndentedJSON(401, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}
		err := my_jwt.ValidateToken(tokenString)
		if err != nil {
			ctx.IndentedJSON(401, gin.H{"error": "Unable to validate token: " + err.Error()})
			ctx.Abort()
			return
		}
		id, err := my_jwt.ExtractID(tokenString)
		if err != nil {
			ctx.IndentedJSON(400, gin.H{"error": "unable to extract id"})
			ctx.Abort()
			return
		}

		user := Repo.User.GetByID(id)
		if user == nil {
			ctx.IndentedJSON(400, gin.H{"error": "user by token not found"})
			ctx.Abort()
			return
		}

		switch requiredRole {
		case ROLE_ADMIN:
			if user.Role < ROLE_ADMIN {
				ctx.IndentedJSON(403, gin.H{"error": "access denied"})
				ctx.Abort()
				return
			}
		case ROLE_MANAGER:
			if user.Role < ROLE_MANAGER {
				ctx.IndentedJSON(403, gin.H{"error": "access denied"})
				ctx.Abort()
				return
			}
		case ROLE_MASTER:
			if user.Role < ROLE_MASTER {
				ctx.IndentedJSON(403, gin.H{"error": "access denied"})
				ctx.Abort()
				return
			}
		case ROLE_CUSTOMER:
			if user.Role != ROLE_CUSTOMER && user.Role != ROLE_ADMIN {
				log.Println(user.Role, " ", ROLE_ADMIN)
				ctx.IndentedJSON(403, gin.H{"error": "access denied"})
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}
