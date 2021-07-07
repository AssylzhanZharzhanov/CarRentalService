package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/zharzhanov/region/models"
	"net/http"
)

func (h *Handler) createAdvert (c *gin.Context) {
	var advert models.Advert

	if err := c.BindJSON(&advert); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.service.CreateAdvert(c.Request.Context(), advert)

	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) getAllAdverts (c *gin.Context) {


}

func (h *Handler) getAdvertById (c *gin.Context) {

}

func (h *Handler) updateAdvert (c *gin.Context) {

}

func (h *Handler) deleteAdvert (c *gin.Context) {

}