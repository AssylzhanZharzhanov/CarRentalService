package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) searchCarMark(c *gin.Context) {
	value := c.Query("value")
	log.Println(value)
	brands, err := h.service.GetCarModels(c.Request.Context(), value)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, brands)
}

func (h *Handler) getSearchAdverts(c *gin.Context) {
	value := c.Query("value")
	brands, err := h.service.GetAdverts(c.Request.Context(), value)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, brands)
}
