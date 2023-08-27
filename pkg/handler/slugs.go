package handler

import (
	"github.com/gin-gonic/gin"
	entity "github.com/zpix1/avito-test-task/pkg/entities"
	"net/http"
	"strconv"
	"time"
)

// CreateSlug godoc
//
//	@Summary		Create a slug
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

	slugId, err := h.service.CreateSlug(slug.Name, slug.AutoAddPercent)
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

	err := h.service.UpdateUserSlugs(
		slugUpdate.UserId,
		slugUpdate.AddSlugNames,
		slugUpdate.DeleteSlugNames,
		slugUpdate.Ttl,
	)

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteSlug godoc
//
//	@Summary		Delete slug
//	@Description	Delete slug by name
//	@Accept			json
//	@Produce		json
//	@Param			slug_name	path	string	true	"Slug name"
//	@Success		204
//	@Failure		500	{object}	errorMessage
//	@Router			/api/v1/slugs/{slug_name} [delete]
func (h *Handler) DeleteSlug(c *gin.Context) {
	slugName := c.Param("slug_name")

	err := h.service.DeleteSlug(slugName)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// GetUserSlugs godoc
//
//	@Summary		Get user slugs
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

// GetSlugHistoryCsv godoc
//
//	@Summary		Get user slugs history
//	@Description	Get user slugs history in CSV format
//	@Accept			json
//	@Produce		text/csv
//	@Param			user_id	query		int	true	"User id"
//	@Param			start	query		int	true	"Start datetime unixtime (seconds)"
//	@Param			end		query		int	true	"End datetime unixtime (seconds)"
//	@Success		200		{object}	string
//	@Failure		500		{object}	errorMessage
//	@Router			/api/v1/slugs/history [get]
func (h *Handler) GetSlugHistoryCsv(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}
	start, err := strconv.Atoi(c.Query("start"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}
	end, err := strconv.Atoi(c.Query("end"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	csv, err := h.service.GetSlugHistoryCsv(userId,
		time.Unix(int64(start), 0),
		time.Unix(int64(end), 0))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Header("Content-Type", "text/csv; charset=utf-8")

	c.String(http.StatusOK, csv)
}
