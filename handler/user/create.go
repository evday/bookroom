package user

import (

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log/lager"
	"github.com/lexkong/log"
	"bookroom/util"
	"bookroom/pkg/errno"
	"bookroom/models"
	. "bookroom/handler"
)

func Create(c *gin.Context) {
	log.Info("User Create function called.",lager.Data{"X-Request-Id":util.GetReqID(c)})

	var r CreateUser
	if err := c.Bind(&r);err != nil {
		SendResponse(c,errno.ErrBind,nil)
		return
	}

	if r.Password != r.Repassword {
		SendResponse(c,errno.ErrNotMatch,nil)
		return
	}

	u := models.User{
		Name:r.Username,
		Password:r.Password,
		IsAdmin:r.IsAdmin,
	}

	if err := u.Validate(); err != nil {

		log.Error("Validate the data.", err)
		SendResponse(c, &errno.Errno{Code: errno.ErrValidation.Code, Message: err.Error()}, nil)
		return
	}

	if err := u.Create();err != nil {
		SendResponse(c,errno.ErrDatabase,nil)
		return
	}

	rsp := CreateUser{
		Username:r.Username,
		CreateAt:u.CreateAt.Format("2006-01-02 15:04:05"),
		IsAdmin:u.IsAdmin,
	}

	SendResponse(c,nil,rsp)
}