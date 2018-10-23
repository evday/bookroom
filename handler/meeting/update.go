package meeting

import (
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
	. "bookroom/handler"
	"bookroom/pkg/errno"
	"bookroom/models"
)

func UpdateMeetInfo(c *gin.Context){
	recordId,_ := strconv.Atoi(c.Param("id"))

	var m Createmeet

	if err := c.Bind(&m);err != nil{
		SendResponse(c,errno.ErrDatabase,nil)
		return
	}

	u,err := models.GetUser(m.Username)
	if err != nil {
		SendResponse(c,errno.ErrUserNotFound,nil)
		return
	}

	db := models.GetSelfDB()

	records := []models.Record{}

	db.Where("username = ?",u.Name).Find(&records)


	for i:=0;i<len(records);i++ {
		if (records[i].ID == int64(recordId)){
			break
		}else {
			SendResponse(c,errno.ErrNotAllow,nil)
			return
		}
	}

	room,err := models.GetRoom(int64(m.RoomId))
	if err != nil {
		SendResponse(c,errno.ErrRoomNotFound,nil)
		return
	}

	startTime, _:= time.ParseInLocation("2006-01-02 15:04", m.StartTime, time.Local)

	endTime, _ := time.ParseInLocation("2006-01-02 15:04", m.EndTime, time.Local)

	record := models.Record{
		RoomName:room.Name,
		UserName:u.Name,
		Start:startTime,
		End:endTime,
		Theme:m.Theme,
		Member:m.Member,
	}

	record.ID = int64(recordId)

	if err := record.Update();err != nil{
		SendResponse(c,errno.ErrDatabase,nil)
		return
	}

	SendResponse(c,errno.OK,nil)
}
