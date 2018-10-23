package equipment

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"bookroom/models"
	"bookroom/pkg/errno"
	. "bookroom/handler"
)

func Delete(c *gin.Context) {
	equipmentId,_ := strconv.Atoi(c.Param("id"))
	if err := models.DeleteEquip(int64(equipmentId));err != nil{
		SendResponse(c,errno.ErrDatabase,nil)
		return
	}
	SendResponse(c,errno.OK,nil)
}