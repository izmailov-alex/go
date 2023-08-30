package controllers

import (
	"example/web-service-avito/initializers"
	"example/web-service-avito/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSegment(c *gin.Context) error {
	var seg models.Segment
	name := c.Param("segment_name")
	row := initializers.DB.QueryRow("SELECT * FROM segments WHERE segment_name = ?", name)
	if err := row.Scan(&seg.SegmentID, &seg.SegmentName); err != nil {
		return err
	}
	_, err := initializers.DB.Exec("DELETE FROM user_segments WHERE segment_id = ?", seg.SegmentID)
	if err != nil {
		return err
	}
	_, err = initializers.DB.Exec("DELETE FROM segments WHERE segment_id = ?", seg.SegmentID)
	if err != nil {
		return err
	}
	c.IndentedJSON(http.StatusOK, seg)
	return nil
}
