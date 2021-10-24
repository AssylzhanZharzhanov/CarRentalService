package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) searchCarMark(c *gin.Context) {
	brand := c.Query("brand")

	brands, err := h.service.GetCarModels(c.Request.Context(), brand)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, brands)
}
