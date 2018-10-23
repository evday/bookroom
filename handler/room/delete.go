package room

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"bookroom/models"
	"bookroom/pkg/errno"
	. "bookroom/handler"
)

func DeleteRoom(c *gin.Context){
	roomId,_ := strconv.Atoi(c.Param("id"))

	fmt.Println(roomId)
	db := models.GetSelfDB()
	tx := db.Begin()
	if err := models.DeleteRoom(int64(roomId));err != nil {
		SendResponse(c,errno.ErrDatabase,nil)
		tx.Rollback()
		return
	}
	if err := models.DeleteEquiRoom(int64(roomId));err != nil {
		SendResponse(c,errno.ErrDatabase,nil)
		tx.Rollback()
		return
	}
	tx.Commit()
	SendResponse(c,errno.OK,nil)
}