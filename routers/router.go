package routers

import (
	"github.com/gin-gonic/gin"
	"mooblog/api/v1"
	"mooblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		//user模块路由借楼
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		//分类模块的理由借口

		//文章模块的理由借口

	}
	r.Run(utils.HttpPort)
}
