package routes

import "github.com/gin-gonic/gin"

func initRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId")
}
