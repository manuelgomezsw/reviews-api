package router

import (
	"github.com/gin-gonic/gin"
	"reviews-api/internal/infraestructure/controller/registry"
	"reviews-api/internal/infraestructure/controller/search"
)

func mapURLs(router *gin.Engine) {
	registryURLs(router)
	searchURLs(router)
}

func registryURLs(router *gin.Engine) {
	router.POST("/reviews", registry.Create)
	router.PUT("/reviews/:review_id", registry.Update)
	router.DELETE("/reviews/:review_id", registry.Delete)
}

func searchURLs(router *gin.Engine) {
	router.GET("/reviews/:review_id", search.GetByID)
}
