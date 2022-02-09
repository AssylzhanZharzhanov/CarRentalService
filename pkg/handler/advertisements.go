package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/zharzhanov/region/models"
	"net/http"
)

func (h *Handler) addAdvertisement(c *gin.Context) {
	var advertisement models.AdvertisementInput

	if err := c.BindJSON(&advertisement); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, inputError)
		return
	}

	err := h.service.CreateAdvertisement(c.Request.Context(), advertisement)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, cannotCreateError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getAdvertisements(c *gin.Context) {
	adverts, err := h.service.GetAdvertisements(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, notFoundError)
		return
	}

	c.JSON(http.StatusOK, adverts)
}

func (h *Handler) getAdvertisementByID(c *gin.Context) {

}

func (h *Handler) updateAdvertisement(c *gin.Context) {

}

func (h *Handler) deleteAdvertisement(c *gin.Context) {

}

func (h *Handler) connectAdvertisement(c *gin.Context) {

}

func (h *Handler) updateConnectedAdvertisement(c *gin.Context) {

}