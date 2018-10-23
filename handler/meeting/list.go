package meeting

import (
	"github.com/gin-gonic/gin"
	"bookroom/models"
	. "bookroom/handler"
	"bookroom/pkg/errno"
)

func ListMeetInfo(c *gin.Context){
	db := models.GetSelfDB()
	var records []models.Record
	db.Find(&records)
	SendResponse(c,errno.OK,records)
}