package handler

import (
	avito "avito-backend-task-2023"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary CreateSegment
// @Tags segment
// @Description create a new segment
// @ID create-segment
// @Accept  json
// @Produce  json
// @Param input body avito.Segment true "segment info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segment [post]
func (h *Handler) createSegment(c *gin.Context) {
	var input avito.Segment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Segment.CreateSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary DeleteSegment
// @Tags segment
// @Description delete segment by name
// @ID delete-segment
// @Accept  json
// @Produce  json
// @Param input body avito.Segment true "segment name"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /segment [delete]
func (h *Handler) deleteSegment(c *gin.Context) {
	var input avito.Segment

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Segment.DeleteSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
