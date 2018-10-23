package user

import (
	. "bookroom/handler"
	"bookroom/pkg/errno"
	"bookroom/service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var r ListRequest

	if err := c.ShouldBind(&r);err != nil {
		SendResponse(c,errno.ErrBind,nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}