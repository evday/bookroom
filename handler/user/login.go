package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"bookroom/models"
	"bookroom/util"
	"bookroom/pkg/errno"
	"bookroom/pkg/token"
	. "bookroom/handler"
	
)

func Login(c *gin.Context){

	u := LoginRequest{}
	if err := c.Bind(&u);err != nil {
		SendResponse(c,errno.ErrBind,nil)
		return
	}
	user,err := models.GetUser(u.Username)
	if err != nil {
		SendResponse(c,errno.ErrUserNotFound,nil)
		return
	}

	if err := user.Compare(u.Password);err != nil {
		SendResponse(c,errno.ErrPasswordIncorrect,nil)
		return
	}

	t,n,s,err := token.Sign(c,token.Context{ID:user.ID,Username:user.Name,Isadmin:user.IsAdmin},"")
	if err != nil {
		SendResponse(c,errno.ErrToken,nil)
		return
	}

	conn := util.Get()
	defer conn.Close()
	_,err = conn.Do("SET","token:"+fmt.Sprintf("%d",user.ID),t,"EX","360")
	if err != nil {
		SendResponse(c,errno.ErrToken,nil)
		return
	}
	SendResponse(c,nil,models.Token{Token:t,Username:n,Isadmin:s})
}