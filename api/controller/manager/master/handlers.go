package manager

import (
	. "kursarbeit/config"
	. "kursarbeit/storage/models/workshop"
	"kursarbeit/utils/validator"
	"net/mail"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMasters(ctx *gin.Context) {
	masters := Repo.Master.GetAll(ctx)

	ctx.IndentedJSON(200, masters)
}

func GetMasterById(ctx *gin.Context) {
	masterID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || masterID < 1 {
		ctx.IndentedJSON(400, gin.H{"error": "invalid master id"})
		return
	}

	master := Repo.Master.GetByID(masterID)
	if master == nil {
		ctx.IndentedJSON(404, gin.H{"error": "master not found"})
		return
	}

	ctx.IndentedJSON(200, master)
}

func PostMaster(ctx *gin.Context) {
	var master Master
	if err := ctx.ShouldBindJSON(&master); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to parse request body"})
		return
	}

	if err := master.Insert(); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to insert received record"})
		return
	}
}

func PutMaster(ctx *gin.Context) {
	masterID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || masterID < 1 {
		ctx.IndentedJSON(400, gin.H{"error": "invalid id"})
		return
	}

	var masterReq Master
	if err := ctx.ShouldBindJSON(&masterReq); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "unable to parse request body"})
		return
	}

	master := Repo.Master.GetByID(masterID)
	if master == nil {
		ctx.IndentedJSON(404, gin.H{"error": "requested user not found"})
		return
	}

	masterChanged := false

	if masterReq.PersonalInfo != nil {
		if masterReq.PersonalInfo.Name != "" {
			master.PersonalInfo.Name = masterReq.PersonalInfo.Name
			masterChanged = true
		}
		if masterReq.PersonalInfo.Email != "" {
			_, err := mail.ParseAddress(masterReq.PersonalInfo.Email)
			if err != nil {
				ctx.IndentedJSON(400, gin.H{"error": "invalid email format"})
				return
			}
			master.PersonalInfo.Email = masterReq.PersonalInfo.Email
			masterChanged = true
		}
		if masterReq.PersonalInfo.PhoneNumber != "" {
			if validator.CheckPhoneNumber(masterReq.PersonalInfo.PhoneNumber) {
				master.PersonalInfo.PhoneNumber = masterReq.PersonalInfo.PhoneNumber
				masterChanged = true
			} else {
				ctx.IndentedJSON(400, gin.H{"error": "invalid phone format"})
				return
			}
		}
	}

	if masterReq.MasterRankID > 0 && masterReq.MasterRankID < 5 {
		master.MasterRankID = masterReq.MasterRankID
		masterChanged = true
	}

	if masterChanged {
		if err := master.Update(); err != nil {
			ctx.IndentedJSON(400, gin.H{"error": "unable to update received record: " + err.Error()})
		}
	}
}

func DeleteMaster(ctx *gin.Context) {
	masterID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || masterID < 1 {
		ctx.IndentedJSON(400, gin.H{"error": "invalid id"})
		return
	}

	if err := Repo.Master.DeleteByID(masterID); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
	}
}
