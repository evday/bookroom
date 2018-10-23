package equipment

import (

	"github.com/gin-gonic/gin"
	"bookroom/models"
	"bookroom/pkg/errno"
	. "bookroom/handler"
)

func Create(c *gin.Context) {

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

	if err := e.Create();err != nil {
		SendResponse(c,errno.ErrDatabase,nil)
		return
	}
	rsp := CreateEquipment{
		Name:r.Name,
		Brand:r.Brand,
		Model:r.Model,
		CreateAt:e.CreateAt.Format("2006-01-02 15:04:05"),
	}
	SendResponse(c,nil,rsp)
}