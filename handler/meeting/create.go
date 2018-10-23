package meeting

import (
	"time"
	"github.com/gin-gonic/gin"
	. "bookroom/handler"
	"bookroom/pkg/errno"
	"bookroom/models"
)

func CreateMeet(c *gin.Context){
	var m Createmeet
	db := models.GetSelfDB()
	
	if err := c.Bind(&m);err != nil {
		SendResponse(c,errno.ErrBind,nil)
		return
	}
	u,err := models.GetUser(m.Username)
	if err != nil {
		SendResponse(c,errno.ErrUserNotFound,nil)
		return
	}

	room,err := models.GetRoom(int64(m.RoomId))
	if err != nil {
		SendResponse(c,errno.ErrRoomNotFound,nil)
		return
	}


	createDate,_ := time.ParseInLocation("2006-01-02", m.Date, time.Local)

	startTime, _:= time.ParseInLocation("2006-01-02 15:04", m.StartTime, time.Local)

	endTime, _ := time.ParseInLocation("2006-01-02 15:04", m.EndTime, time.Local)

	currentTime:=time.Now().Format("2006-01-02")
	currenttime,_ := time.ParseInLocation("2006-01-02", currentTime, time.Local)

	minTime,_ := time.ParseInLocation("2006-01-02 15:04",m.Date+" "+"09:00", time.Local)
	maxTime,_ := time.ParseInLocation("2006-01-02 15:04",m.Date+" "+"21:00", time.Local)

	if createDate.Before(currenttime) || (startTime.Before(minTime) && endTime.After(maxTime)) {
		SendResponse(c,errno.ErrOutRange,nil)
		return
	}

	//查询该会议室下所有的会议记录
	var records []models.Record
	db.Where("roomname in (?)",room.Name).Find(&records)

	for index,_ := range records{
		if (records[index].Start.Before(startTime)||records[index].Start.Equal(startTime)) && (records[index].End.After(endTime)||records[index].End.Equal(endTime)){
			SendResponse(c,errno.ErrExist,nil)
			return
		}
	}

	meet := models.Record{
		RoomName:room.Name,
		UserName:u.Name,
		Start:startTime,
		End:endTime,
		Theme:m.Theme,
		Member:m.Member,
	}

	if err := meet.Create();err != nil {
		SendResponse(c,errno.ErrDatabase,nil)
		return
	}

	SendResponse(c,errno.OK,nil)
}