package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1/recruitment/personaldata")
	SetShowTimeRouters(v1)
	return router
}
