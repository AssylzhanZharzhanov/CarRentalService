package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/zharzhanov/region/models"
	"net/http"
)

func (h *Handler) addRentType(c *gin.Context) {
	var rentType models.RentTypes

	if err := c.BindJSON(&rentType); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bind error")
		return
	}

	err := h.service.AddRentType(c.Request.Context(), rentType)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can not add")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getRentTypes(c *gin.Context) {
	rentTypes, err := h.service.GetRentTypes(c.Request.Context())

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rentTypes)
}

func (h *Handler) deleteRentType(c *gin.Context) {
	value := c.Query("value")
	err := h.service.DeleteRentType(c.Request.Context(), value)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getCategories(c *gin.Context) {
	categories, err := h.service.GetCategories(c.Request.Context())

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *Handler) addCategory(c *gin.Context) {
	var category models.Category

	if err := c.BindJSON(&category); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bind error")
		return
	}

	err := h.service.AddCategory(c.Request.Context(), category)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can not add")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteCategory(c *gin.Context) {
	value := c.Query("value")
	err := h.service.DeleteCategory(c.Request.Context(), value)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getCities(c *gin.Context) {
	cities, err := h.service.GetCities(c.Request.Context())

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cities)
}

func (h *Handler) addCity(c *gin.Context) {
	var city models.City

	if err := c.BindJSON(&city); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bind error")
		return
	}

	err := h.service.AddCity(c.Request.Context(), city)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can not add")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteCity(c *gin.Context) {
	value := c.Query("value")
	err := h.service.DeleteCity(c.Request.Context(), value)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) addPrice(c *gin.Context) {
	var price models.Price

	if err := c.BindJSON(&price); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bind error")
		return
	}

	err := h.service.AddPrice(c.Request.Context(), price)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can not add")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getPrice(c *gin.Context) {
	prices, err := h.service.GetPrices(c.Request.Context())

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, prices)
}

func (h *Handler) deletePrice(c *gin.Context) {
	value := c.Query("value")
	err := h.service.DeletePrices(c.Request.Context(), value)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}