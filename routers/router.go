package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter()*gin.Engine{
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())


	apiv1 := router.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/user",)
	}

	return router
}
