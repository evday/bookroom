package meeting

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"bookroom/models"
	. "bookroom/handler"
	"bookroom/pkg/errno"
)

func DeleteMeet(c *gin.Context){
	meetId,_ := strconv.Atoi(c.Param("id"))
	username := c.Param("username")

	u,err := models.GetUser(username)
	if err != nil {
		SendResponse(c,errno.ErrUserNotFound,nil)
		return
	}

	db := models.GetSelfDB()

	records := []models.Record{}

	db.Where("username = ?",u.Name).Find(&records)


	for i:=0;i<len(records);i++ {
		if (records[i].ID == int64(meetId)){
			break
		}else {
			SendResponse(c,errno.ErrNotAllow,nil)
			return
		}
	}

	if err := models.DeleteMeetInfo(int64(meetId),u.Name);err != nil{
		SendResponse(c,errno.ErrDatabase,nil)
		return
	}
	SendResponse(c,errno.OK,nil)
}