package server

import (
	"os"

	"apsdatabase/controllers"
	"apsdatabase/middleware"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	if len(os.Getenv("GIN_MODE")) == 0 {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// adding middleware server
	router.Use(middleware.CORSMiddleware())
	// health check

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	Routes(router)
	router.Run(":9090") // listen and serve on 0.0.0.0:808
}

func Routes(r *gin.Engine) {
	r.POST("/enquiry", controllers.CreateEnquiry)
	

}
