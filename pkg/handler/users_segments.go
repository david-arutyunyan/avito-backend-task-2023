package handler

import (
	avito "avito-backend-task-2023"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary UpdateUserSegments
// @Tags user_segments
// @Description update user's segments
// @ID update-user-segments
// @Accept  json
// @Produce  json
// @Param input body avito.AlteredUserSegments true "user_id and segments to add/delete"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users-segments [post]
func (h *Handler) updateUserSegments(c *gin.Context) {
	var input avito.AlteredUserSegments

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.UsersSegments.UpdateUserSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary GetUserSegments
// @Tags user_segments
// @Description get user's segments
// @ID get-user-segments
// @Accept  json
// @Produce  json
// @Param id path string true "the UUID of a user"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users-segments/{id} [GET]
func (h *Handler) getUserSegments(c *gin.Context) {
	id := c.Param("id")

	lists, err := h.services.UsersSegments.GetUserSegments(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary GetUserSegmentsLogs
// @Tags user_segments
// @Description get user's segments logs
// @ID get-user-segments-logs
// @Accept  json
// @Produce text/csv
// @Param input body avito.CustomDate true "year and month"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users-segments/logs [POST]
func (h *Handler) getUserSegmentsLogs(c *gin.Context) {
	var input avito.CustomDate

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	c.Writer.Header().Add("Content-Disposition", `attachment; filename="logs.csv"`)

	err := h.services.UsersSegments.GetUserSegmentsLogs(c.Writer, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	return
}
