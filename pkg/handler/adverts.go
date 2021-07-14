package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
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

type Filter struct {
	City string `json:"city" bson:"city,omitempty"`
	Category string `json:"category" bson:"category,omitempty"`
	RentType string `json:"rent_type" bson:"rent_type,omitempty"`
	Price int `json:"price" bson:"price,omitempty"`
}

func (h *Handler) getAllAdverts(c *gin.Context) {

	query := bson.M{}

	if city := c.Query("city"); city != "" {
		query["city"] = city
	}

	if category := c.Query("category"); category != "" {
		query["category"] = category
	}

	if rentType := c.Query("rent_type"); rentType != "" {
		query["rent_type"] = rentType
	}

	if title := c.Query("title"); title != "" {
		query["title"] = bson.M{

		}
	}

	if c.Query("minPrice") != "" && c.Query("maxPrice") != ""  {

		minPrice, _ := strconv.Atoi(c.Query("minPrice"))
		maxPrice, _ := strconv.Atoi(c.Query("maxPrice"))

		query["price"] = bson.M{"$gte": minPrice, "$lte": maxPrice}
	}

	adverts, err := h.service.GetAllAdverts(c.Request.Context(), query)

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
