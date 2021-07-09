package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/zharzhanov/region/models"
	"net/http"
)

func (h *Handler) createAdvert(c *gin.Context) {
	var advert models.Advert

	if err := c.BindJSON(&advert); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.service.CreateAdvert(c.Request.Context(), advert)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) getAllAdverts(c *gin.Context) {

	adverts, err := h.service.GetAllAdverts(c.Request.Context())

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, adverts)
}

func (h *Handler) getAdvertById(c *gin.Context) {
	id := c.Param("id")

	advert, err := h.service.GetAdvertById(c.Request.Context(), id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, advert)
}

func (h *Handler) updateAdvert(c *gin.Context) {
	id := c.Param("id")

	var newAdvert models.UpdateAdvertInput
	if err := c.BindJSON(&newAdvert); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.service.UpdateAdvert(c.Request.Context(), id, newAdvert)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteAdvert(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteAdvert(c.Request.Context(), id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
