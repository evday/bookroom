package user

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"bookroom/models"
	"bookroom/pkg/errno"
	. "bookroom/handler"
	
)

func Delete(c *gin.Context) {
	userId,_ := strconv.Atoi(c.Param("id"))
	if err := models.DeleteUser(int64(userId));err != nil {
		SendResponse(c,errno.ErrDatabase,nil)
		return
	}
	SendResponse(c,nil,nil)
}