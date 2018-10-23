package middleware

import (
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
	"bookroom/pkg/token"
	handler "bookroom/handler"
	"bookroom/pkg/errno"
	"bookroom/models"
)

var  allowUrls = []string{"/","/v1/room"}

func Container(s string)bool {
	for _,str := range allowUrls {
		if s == str {
			return true
		}
	}
	return false
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		urlstr := c.Request.RequestURI	
		fmt.Println(urlstr)
		if con,err := token.ParseRequest(c);err != nil {
			handler.SendResponse(c,errno.ErrTokenInvilid,nil)
			c.Abort()
			return
		}else {
			if Container(urlstr)&&c.Request.Method == "GET" || strings.Contains(urlstr,"/v1/meet"){
				c.Next()
			}else {
				user,_ := models.GetUser(con.Username)
				if user.IsAdmin{
					c.Next()
				}else {
					handler.SendResponse(c,errno.ErrPermission,nil)
					c.Abort()
				}
			}
			
			
		}
		c.Next()
	}
}
