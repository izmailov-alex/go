package controllers

import (
	"example/web-service-avito/initializers"
	"example/web-service-avito/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSegments(c *gin.Context) error {
	var segments []models.Segment
	rows, err := initializers.DB.Query("SELECT * FROM segments")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var seg models.Segment
		if err := rows.Scan(&seg.SegmentID, &seg.SegmentName); err != nil {
			return err
		}
		segments = append(segments, seg)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	c.IndentedJSON(http.StatusOK, segments)
	return err
}

func GetSegmentID(c *gin.Context) error {
	var seg models.Segment
	id := c.Param("id")
	row := initializers.DB.QueryRow("SELECT * FROM segments WHERE segment_id=?", id)
	if err := row.Scan(&seg.SegmentID, &seg.SegmentName); err != nil {
		return err
	}
	c.IndentedJSON(http.StatusOK, seg)
	return nil
}
