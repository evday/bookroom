package router

import (
	"net/http"
	user "bookroom/handler/user"
	equip "bookroom/handler/equipment"
	room "bookroom/handler/room"
	meet "bookroom/handler/meeting"
	"github.com/gin-gonic/gin"
	"bookroom/handler/sd"
	"bookroom/router/middleware"

	
	
)

func Load(g *gin.Engine,mv ...gin.HandlerFunc)*gin.Engine {
	g.Use(gin.Recovery())
	g.Use(mv...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound,"The incorrect API route.")
	})

	g.POST("/",user.Login)
	g.POST("/login",user.Login)
	u := g.Group("/v1")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("/user",user.Create)
		u.GET("/user",user.List)
		u.DELETE("/user/:id",user.Delete)
		u.POST("/equip",equip.Create)
		u.GET("/equip",equip.ListEquipment)
		u.GET("/equip/:keyword",equip.ListEquipment)
		u.DELETE("/equip/:id",equip.Delete)
		u.PUT("/equip/:id",equip.Update)
		u.POST("/room",room.Create)
		u.GET("/room",room.ListRoom)
		u.DELETE("/room/:id",room.DeleteRoom)
		u.PUT("/room/:id",room.UpdateRoom)
		u.GET("/room/:keyword",room.ListRoom)
		u.POST("/meet",meet.CreateMeet)
		u.GET("/meet",meet.ListMeetInfo)
		u.DELETE("/meet/:id/:username",meet.DeleteMeet)
		u.PUT("/meet/:id",meet.UpdateMeetInfo)
		u.GET("/meet/:start/:end/:roomid",meet.SearchMeetInfo)
	}

	svcd := g.Group("/sd")
	{
		svcd.GET("/health",sd.HealthCheck)
	}

	return g
}