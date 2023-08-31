package main

import (
	"example/web-service-avito/controllers"
	"example/web-service-avito/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	router.GET("/segments/:segment_id", GinHandler(controllers.GetSegmentID))
	router.GET("/user_segments/:user_id", GinHandler(controllers.GetUserSegment))
	router.GET("/tables", GinHandler(controllers.GetTables))

	router.POST("/segments/:segment_name", GinHandler(controllers.PostSegments))
	router.POST("/users/:user_id", GinHandler(controllers.PostUsers))

	router.PUT("/user_segments/:user_id", GinHandler(controllers.PutUserSegments))

	router.DELETE("/segments/:segment_name", GinHandler(controllers.DeleteSegment))

	router.Run("0.0.0.0:8081")
}
