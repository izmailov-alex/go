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
	id := c.Param("segment_id")
	row := initializers.DB.QueryRow("SELECT * FROM segments WHERE segment_id=?", id)
	if err := row.Scan(&seg.SegmentID, &seg.SegmentName); err != nil {
		return err
	}
	c.IndentedJSON(http.StatusOK, seg)
	return nil
}

func GetUserSegment(c *gin.Context) error {
	var userSegments []models.Segment
	userID := c.Param("user_id")
	rows, err := initializers.DB.Query("SELECT * FROM user_segments WHERE user_id = ?", userID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var userSeg models.UserSegments
		if err := rows.Scan(&userSeg.UserID, &userSeg.SegmentID); err != nil {
			return err
		}
		var seg models.Segment
		row := initializers.DB.QueryRow("SELECT * FROM segments WHERE segment_id = ?", userSeg.SegmentID)
		if err := row.Scan(&seg.SegmentID, &seg.SegmentName); err != nil {
			return err
		}
		userSegments = append(userSegments, seg)
	}
	c.IndentedJSON(http.StatusOK, userSegments)
	return nil
}
