package equipment

import (
	"github.com/gin-gonic/gin"
	"bookroom/models"
	. "bookroom/handler"
	"bookroom/pkg/errno"
	"fmt"

)

func ListEquipment(c *gin.Context){
	db := models.GetSelfDB()
	var equips []models.Equipment
	keyword := c.Param("keyword")
	if keyword != ""{
		db.Where("equipname LIKE ?",fmt.Sprintf("%%%s%%",keyword)).Order("id desc").Find(&equips)
	}else {
		db.Find(&equips)
	}
	SendResponse(c,errno.OK,equips)
}