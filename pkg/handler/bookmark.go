package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) addBookmark(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	advertId := c.Query("advertId")
	log.Printf("advertId: %s", advertId)

	err = h.service.AddUserBookmark(c.Request.Context(), userId, advertId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSONP(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getBookmarks(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	adverts, err := h.service.GetUserBookmarks(c.Request.Context(), userId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSONP(http.StatusOK, adverts)
}

