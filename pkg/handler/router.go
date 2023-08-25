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
		lists := api.Group("/slugs")
		{
			lists.POST("/", h.CreateSlug)
			lists.GET("/get", h.GetUserSlugs)
			lists.PUT("/update", h.UpdateUserSlugs)
			lists.DELETE("/:slug_name", h.DeleteSlug)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
