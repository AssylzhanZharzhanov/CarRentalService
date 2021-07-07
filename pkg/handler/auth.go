package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/zharzhanov/region/models"
	"net/http"
)

func (h *Handler) signUp (c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.service.SignUp(c.Request.Context(), user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) signIn(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	token, err := h.service.SignIn(c.Request.Context(), user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, token)
}
