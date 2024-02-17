package registry

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reviews-api/internal/domain"
	"reviews-api/internal/domain/registry/services"
	"strconv"
)

func Create(c *gin.Context) {
	var newReview domain.Review
	if err := c.ShouldBindJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing body",
			"error":   err.Error(),
		})
		return
	}

	if err := services.Create(&newReview); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error posting review",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newReview)
}

func Update(c *gin.Context) {
	var review domain.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing body",
			"error":   err.Error(),
		})
		return
	}

	reviewID, err := strconv.ParseInt(c.Param("review_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing review_id",
			"error":   err.Error(),
		})
		return
	}
	review.ReviewID = reviewID

	if err := services.Update(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating review",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, review)
}

func Delete(c *gin.Context) {
	reviewID, err := strconv.ParseInt(c.Param("review_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error serializing review_id",
			"error":   err.Error(),
		})
		return
	}

	if err := services.Delete(reviewID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting review",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}
