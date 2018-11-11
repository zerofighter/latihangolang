package routers

import (
	"github.com/gin-gonic/gin"
	"../controllers"
)

func SetShowTimeRouters(router *gin.RouterGroup) {
	router.GET("/", controllers.GetShowTimes)
	router.GET("/:id", controllers.GetShowTimeById)
	router.POST("/", controllers.Update)
	//router.Handle("/showtimes/{id}", authorization.AuthMiddleware(controllers.GetShowTimeById)).Methods("GET")

}
