package controllers

import (
	"example/web-service-avito/initializers"
	"example/web-service-avito/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// postSegments adds an segment from JSON received in the request header.
func PostSegments(c *gin.Context) error {
	var newSegment models.Segment
	name := c.Param("segment_name")
	reslult, err := initializers.DB.Exec("INSERT INTO segments (segment_name) VALUES (?)", name)
	if err != nil {
		return err
	}
	id, err := reslult.LastInsertId()
	if err != nil {
		return err
	}
	newSegment.SegmentID = strconv.FormatInt(id, 10)
	newSegment.SegmentName = name
	c.IndentedJSON(http.StatusCreated, newSegment)
	return nil
}

func PostUsers(c *gin.Context) error {
	var newUsers models.User
	userID := c.Param("user_id")
	_, err := initializers.DB.Exec("INSERT INTO users (user_id) VALUES (?)", userID)
	if err != nil {
		return err
	}
	newUsers.UserID = userID
	c.IndentedJSON(http.StatusCreated, newUsers)
	return nil
}
