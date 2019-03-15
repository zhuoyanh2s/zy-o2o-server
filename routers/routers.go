package routers

import (
	"github.com/gin-gonic/gin"
	"zy-o2o-service/middleware/jwt"
	"zy-o2o-service/views"
)

var Engine *gin.Engine

func Setup() {

	Engine = gin.New()
	Engine.Use(gin.Logger())
	Engine.Use(gin.Recovery())
	apiv1 := Engine.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/user", views.GetUser)
		apiv1.POST("/user", views.CreatedUser)
		apiv1.PUT("/user:id", views.UpdateUser)
		apiv1.DELETE("/user:id", views.DeleteUser)
	}

}
