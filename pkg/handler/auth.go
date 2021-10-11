package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/zharzhanov/region/models"
	"net/http"
)

func (h *Handler) verifyCode (c *gin.Context) {
	var code models.InputCode

	if err := c.BindJSON(&code); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, inputError)
		return
	}

	token, err := h.service.VerifyCode(c.Request.Context(), code.Code)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, notFoundError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) generateCode(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, inputError)
		return
	}

	code, err := h.service.SendSMS(c.Request.Context(), user.Phone)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, smsError)
		return
	}

	response := models.GeneratedCode{Code: code}

	c.JSON(http.StatusOK, response)
}


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

