package room

import (
	"fmt"

	"strconv"
	"github.com/gin-gonic/gin"
	"bookroom/models"
	. "bookroom/handler"
	"bookroom/pkg/errno"
)



func ListRoom(c *gin.Context){
	db := models.GetSelfDB()
	var rooms []models.Room
	keyword := c.Param("keyword")
	var r []CreateRoom

	if keyword != ""{
		db.Where("roomname LIKE ?",fmt.Sprintf("%%%s%%",keyword)).Order("id desc").Find(&rooms)
		for i:=0;i<len(rooms);i++{
			var e []models.EquiRoom
			
			db.Where("room_id = ?",rooms[i].ID).Find(&e)
			var mm []interface{}
			for j:=0;j<len(e);j++{
				mp := make(map[string]string)
				mp["name"] = e[j].EquipName
				num := strconv.Itoa(int(e[j].EquiNum))
				mp["num"] = num
				mm = append(mm,mp)
			}
			
			room :=  CreateRoom{
				ID:rooms[i].ID,
				Name:rooms[i].Name,
				Capacity:rooms[i].Capacity,
				Location:rooms[i].Location,
				State:rooms[i].State,
				Equipment:mm,
			}
			r = append(r,room)
		}
	}else {
		db.Find(&rooms)
		for i:=0;i<len(rooms);i++{
			var e []models.EquiRoom
			
			db.Where("room_id = ?",rooms[i].ID).Find(&e)
			var mm []interface{}
			for j:=0;j<len(e);j++{
				mp := make(map[string]string)
				mp["name"] = e[j].EquipName
				num := strconv.Itoa(int(e[j].EquiNum))
				mp["num"] = num
				mm = append(mm,mp)
			}
			
			room :=  CreateRoom{
				ID:rooms[i].ID,
				Name:rooms[i].Name,
				Capacity:rooms[i].Capacity,
				Location:rooms[i].Location,
				State:rooms[i].State,
				Equipment:mm,
			}
			r = append(r,room)
		}
	}
	SendResponse(c,errno.OK,r)
}