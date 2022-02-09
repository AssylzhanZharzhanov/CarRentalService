package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get advertisements
// @Security ApiKeyAuth
// @Tags adverts
// @Description Get advertisements
// @ID get-advertisements
// @Accept json
// @Produce json
// @Success 200 {array} object
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/advertisement [get]
func (h *Handler) getAdvertisements(c *gin.Context) {
	adverts, err := h.service.GetAdvertisements(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, notFoundError)
		return
	}

	c.JSON(http.StatusOK, adverts)
}

func (h *Handler) deleteAdvertisement(c *gin.Context) {

}

func (h *Handler) connectAdvertisement(c *gin.Context) {

}