package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zpix1/avito-test-task/docs"
)

func (h *Handler) GetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	api := router.Group("/api/v1")
	{
		slugs := api.Group("/slugs")
		{
			slugs.POST("/", h.CreateSlug)
			slugs.GET("/get", h.GetUserSlugs)
			slugs.PUT("/update", h.UpdateUserSlugs)
			slugs.DELETE("/:slug_name", h.DeleteSlug)
			slugs.GET("/history", h.GetSlugHistoryCsv)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
