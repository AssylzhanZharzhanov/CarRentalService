package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler) addFeedback(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	advertId := c.Query("advertId")

	var feedback models.Feedback
	if err := c.BindJSON(&feedback); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userObjId, _ := primitive.ObjectIDFromHex(userId)
	feedback.UserId = userObjId

	err = h.service.AddFeedback(c.Request.Context(), feedback, advertId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}


func (h *Handler) getFeedback( c *gin.Context) {
	feedbackId := c.Query("feedbackId")

	feedback, err := h.service.GetFeedbackByUserId(c.Request.Context(), feedbackId)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, feedback)
}

func (h *Handler) updateFeedback( c *gin.Context) {
	feedbackId := c.Query("feedbackId")

	var feedback models.Feedback
	if err := c.BindJSON(&feedback); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.service.UpdateFeedback(c.Request.Context(), feedbackId, feedback)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteFeedback( c *gin.Context) {
	feedbackId := c.Query("feedbackId")

	err := h.service.DeleteFeedback(c.Request.Context(), feedbackId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}