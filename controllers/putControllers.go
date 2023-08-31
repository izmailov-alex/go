package controllers

import (
	"example/web-service-avito/initializers"
	"example/web-service-avito/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutUserSegments(c *gin.Context) error {
	var request, response models.Put
	var segmentsAdd, segmentsDelete []models.Segment
	userID := c.Param("user_id")

	// Call BindJSON to bind the received JSON to
	// request.
	if err := c.BindJSON(&request); err != nil {
		return err
	}

	//Checking if the segments are correct
	for _, segmentName := range request.SegmentAdd {
		var seg models.Segment
		row := initializers.DB.QueryRow("SELECT * FROM segments WHERE segment_name = ?", segmentName)
		if err := row.Scan(&seg.SegmentID, &seg.SegmentName); err != nil {
			return err
		}
		segmentsAdd = append(segmentsAdd, seg)
	}
	for _, segmentName := range request.SegmentDelete {
		var seg models.Segment
		row := initializers.DB.QueryRow("SELECT * FROM segments WHERE segment_name = ?", segmentName)
		if err := row.Scan(&seg.SegmentID, &seg.SegmentName); err != nil {
			return err
		}
		segmentsDelete = append(segmentsDelete, seg)
	}

	//Add or remove correct segments
	for _, seg := range segmentsAdd {
		_, err := initializers.DB.Exec("INSERT INTO user_segments (user_id, segment_id) VALUES (?, ?)", userID, seg.SegmentID)
		if err != nil {
			return err
		}
		response.SegmentAdd = append(response.SegmentAdd, seg.SegmentName)
	}
	for _, seg := range segmentsDelete {
		_, err := initializers.DB.Exec("DELETE FROM user_segments WHERE (user_id = ? and segment_id = ?)", userID, seg.SegmentID)
		if err != nil {
			return err
		}
		response.SegmentDelete = append(response.SegmentDelete, seg.SegmentName)
	}
	c.IndentedJSON(http.StatusOK, response)
	return nil
}
