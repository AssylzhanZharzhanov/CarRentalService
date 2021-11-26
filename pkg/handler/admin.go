package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) adminSignIn(c *gin.Context) {
	var user struct{
		Phone string `json:"phone"`
	}

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.service.GetUser(c.Request.Context(), user.Phone)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})}
