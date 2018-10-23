package room

import (
	"strconv"
	"github.com/gin-gonic/gin"
	. "bookroom/handler"
	"bookroom/pkg/errno"
	"bookroom/models"
)

func Create(c *gin.Context){
	var r CreateRoom
	if err := c.Bind(&r);err != nil {
		SendResponse(c,errno.ErrBind,nil)
		return
	}


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
	tx := db.Begin()
	if err := room.Create();err != nil  {
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