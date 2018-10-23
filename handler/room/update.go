package room

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"bookroom/pkg/errno"
	"bookroom/models"
	. "bookroom/handler"
)

func UpdateRoom(c *gin.Context){
	roomId,_ := strconv.Atoi(c.Param("id"))
	
	var r CreateRoom
	if err := c.Bind(&r);err != nil{
		SendResponse(c,errno.ErrBind,nil)
		return
	}
	fmt.Println(r.Equipment)

	var names []string 
	for _,value := range r.Equipment{
		j := value.(map[string]interface{})
		names = append(names,j["name"].(string))
	}

	db := models.GetSelfDB()
	var equipment []models.Equipment
	err := db.Where("equipname in (?)",names).Find(&equipment).Error
	if err != nil {
		SendResponse(c,errno.ErrEquipNotFound,nil)
		return
	}

	room := models.Room{
		Name:r.Name,
		Location:r.Location,
		Capacity:r.Capacity,
		State:r.State,
	}

	room.ID = int64(roomId)

	tx := db.Begin()

	if err := room.Update();err != nil {
		SendResponse(c,errno.ErrDatabase,nil)
		tx.Rollback()
		return
	}

	//先全部删除后重新创建
	if err := models.DeleteEquiRoom(int64(roomId));err != nil {
		SendResponse(c,errno.ErrDatabase,nil)
		tx.Rollback()
		return
	}

	for _,value := range r.Equipment{
		j := value.(map[string]interface{})
		name := j["name"].(string)
		num := j["num"].(string)
		i,_ := strconv.Atoi(num)
		err := models.CreateEquiRoom(room.ID,name,int64(i))
		if err != nil {
			SendResponse(c,errno.ErrDatabase,nil)
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	SendResponse(c,errno.OK,nil)

}