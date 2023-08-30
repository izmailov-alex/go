package main

import (
	"example/web-service-avito/controllers"
	"example/web-service-avito/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type segment struct {
	SegmentID   string `json:"segment_id"`
	SegmentName string `json:"segment_name"`
}

func GinHandler(myhandler func(c *gin.Context) error) (ginhandler func(c *gin.Context)) {
	return func(c *gin.Context) {
		if err := myhandler(c); err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		}
	}
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.GET("/segments", GinHandler(controllers.GetSegments))
	router.GET("/segments/:id", GinHandler(controllers.GetSegmentID))
	router.POST("/segments", GinHandler(controllers.PostSegments))

	router.Run("localhost:8080")
}
