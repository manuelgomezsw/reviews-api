package search

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reviews-api/internal/domain/search/services"
	"strconv"
)

func GetByID(c *gin.Context) {
	reviewID, err := strconv.ParseInt(c.Param("review_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing review_id",
			"error":   err.Error(),
		})
		return
	}

	review, err := services.GetByID(reviewID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting review",
			"error":   err.Error(),
		})
		return
	}

	if review.ReviewID == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, review)
}
