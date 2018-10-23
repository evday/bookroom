package equipment

import (
	"github.com/gin-gonic/gin"
	. "bookroom/handler"
	"strconv"
	"bookroom/pkg/errno"
	"bookroom/models"
)

func Update(c *gin.Context) {
	equipmentId,_ := strconv.Atoi(c.Param("id"))
	var r CreateEquipment

	if err := c.Bind(&r);err != nil {
		SendResponse(c,errno.ErrBind,nil)
		return
	}
	e := models.Equipment{
		Name:r.Name,
		Brand:r.Brand,
		PModel:r.Model,
		Num:r.Store,
	}

	e.ID = int64(equipmentId)

	if err := e.Update();err != nil {
		SendResponse(c,errno.ErrDatabase,nil)
		return
	}
	SendResponse(c,errno.OK,nil)

}