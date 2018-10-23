package meeting

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/gin-gonic/gin"
	. "bookroom/handler"
	"bookroom/models"
	"bookroom/pkg/errno"
)


func SearchMeetInfo(c *gin.Context){
	start := c.Param("start")
	end := c.Param("end")
	roomsId := c.Param("roomid")
	var records []models.Record

	fmt.Println("start  ",start)
	fmt.Println("end  ",end)

	var ids []int64

	if roomsId != "0" && strings.Contains(roomsId,"|"){
		var rms  = strings.Split(roomsId,"|")
		for _,id := range rms{
			intId,_ := strconv.Atoi(id)
			ids = append(ids,int64(intId))
		}
	}else {
		intId,_ := strconv.Atoi(roomsId)
		ids = append(ids,int64(intId))
	}

	db := models.GetSelfDB()

	fmt.Println(ids)

	var rooms []models.Room

	db.Where(ids).Find(&rooms)

	var roomname []string

	for _,obj := range rooms {
		roomname = append(roomname,obj.Name)
	}
	if start != "undefined" && end != "undefined" && roomsId != "0" {	
		db.Where("start >= ? And end <= ? And roomname in (?)",start,end,roomname).Find(&records)
		if len(records) != 0 {
			SendResponse(c,errno.OK,records)
			return
		}else {
			SendResponse(c,errno.NotFound,nil)
			return
		}
	}else if start == "undefined" && end == "undefined" && roomsId != "0" {

		db.Where("roomname in (?)",roomname).Find(&records)
		if len(records) != 0 {
			SendResponse(c,errno.OK,records)
			return
		}else {
			SendResponse(c,errno.NotFound,nil)
			return
		}
	}else if start != "undefined" && end != "undefined" && roomsId == "0" {
		db.Where("start >= ? And end <= ? ",start,end).Find(&records)
		if len(records) != 0 {
			SendResponse(c,errno.OK,records)
			return
		}else {
			SendResponse(c,errno.NotFound,nil)
			return
		}
	}
	SendResponse(c,errno.NotFound,nil)
	return
}