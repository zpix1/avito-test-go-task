package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	entity "github.com/zpix1/avito-test-task/pkg/entities"
	"net/http"
	"strconv"
)

// CreateSlug godoc
//
//	@Summary		Creation of a list
//	@Description	Create a slug by name
//	@Accept			json
//	@Produce		json
//	@Param			slug	body		entity.Slug	true	"Slug object"
//	@Success		201		{integer}	int
//	@Failure		500		{object}	errorMessage
//	@Router			/api/v1/slugs [post]
func (h *Handler) CreateSlug(c *gin.Context) {
	var slug entity.Slug

	if err := c.BindJSON(&slug); err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	slugId, err := h.service.CreateSlug(slug.Name)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": slugId,
	})
}

// UpdateUserSlugs godoc
//
//	@Summary		Update user slugs
//	@Description	Update user slugs by user id and slug names
//	@Accept			json
//	@Produce		json
//	@Param			slug_update	body	entity.SlugUpdate	true	"Slug update object"
//	@Success		204
//	@Failure		400	{object}	errorMessage
//	@Failure		401	{object}	errorMessage
//	@Failure		500	{object}	errorMessage
//	@Router			/api/v1/slugs/update [put]
func (h *Handler) UpdateUserSlugs(c *gin.Context) {
	var slugUpdate entity.SlugUpdate

	if err := c.BindJSON(&slugUpdate); err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	err := h.service.UpdateUserSlugs(slugUpdate.UserId, slugUpdate.AddSlugNames, slugUpdate.DeleteSlugNames)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteSlug godoc
//
//	@Summary		Deletion of a slug
//	@Description	Delete slug by name
//	@Accept			json
//	@Produce		json
//	@Param			slug_name	path	string	true	"Slug name"
//	@Success		204
//	@Failure		500	{object}	errorMessage
//	@Router			/api/v1/slugs/{slug_name} [delete]
func (h *Handler) DeleteSlug(c *gin.Context) {
	slugName := c.Param("slug_name")
	fmt.Println("slug name delete", slugName)

	err := h.service.DeleteSlug(slugName)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// GetUserSlugs godoc
//
//	@Summary		Getting user slugs
//	@Description	Get slugs by user id
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		int	true	"User id"
//	@Success		200		{object}	entity.GetSlugsResponse
//	@Failure		500		{object}	errorMessage
//	@Router			/api/v1/slugs/get [get]
func (h *Handler) GetUserSlugs(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	slugNames, err := h.service.GetUserSlugs(userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"slug_names": slugNames,
	})
}
