package handler

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Create Advert
// @Tags adverts
// @Description create advert
// @ID create-advert
// @Accept mpfd
// @Produce json
// @Param input body models.AdvertInput true "advert body"
// @Success 200 {integer} integer
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/adverts [post]
func (h *Handler) createAdvert(c *gin.Context) {
	advert := models.AdvertInput{}
	if err := c.ShouldBind(&advert); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, inputError)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	files := form.File["images[]"]

	var fileNames []string
	var imageUrls []string
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		log.Println(filename)
		fileNames = append(fileNames, filename)
		imageUrls = append(imageUrls, staticFileHost + filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}

	id, err := h.service.CreateAdvert(c.Request.Context(), advert, imageUrls)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, createObjectError)
		return
	}

	c.JSON(http.StatusOK, id)
}

type Filter struct {
	City     string `json:"city" bson:"city,omitempty"`
	Category string `json:"category" bson:"category,omitempty"`
	RentType string `json:"rent_type" bson:"rent_type,omitempty"`
	Price    int    `json:"price" bson:"price,omitempty"`
}

// @Summary Get all adverts
// @Tags adverts
// @Description get all adverts
// @ID get-all-adverts
// @Accept json
// @Produce json
// @Param city path string true "City"
// @Param category path string true "Category"
// @Param rent_type path string true "Rent type"
// @Param minPrice path string true "Minimum price"
// @Param maxPrice path string true "Maximum price"
// @Param title path string true "Title"
// @Success 200 {array} models.Advert
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/adverts [get]
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
		query["title_search"] = bson.M{"$regex": "[" + strings.ToLower(title) + "]"}
	}

	if c.Query("minPrice") != "" && c.Query("maxPrice") != "" {

		minPrice, _ := strconv.Atoi(c.Query("minPrice"))
		maxPrice, _ := strconv.Atoi(c.Query("maxPrice"))

		query["price"] = bson.M{"$gte": minPrice, "$lte": maxPrice}
	}

	adverts, err := h.service.GetAllAdverts(c.Request.Context(), query)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, notFoundError)
		return
	}

	c.JSON(http.StatusOK, adverts)
}

func (h *Handler) getAdvertById(c *gin.Context) {
	id := c.Param("id")

	advert, err := h.service.GetAdvertById(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, notFoundError)
		return
	}

	c.JSON(http.StatusOK, advert)
}

func (h *Handler) updateAdvert(c *gin.Context) {
	id := c.Param("id")

	var newAdvert models.UpdateAdvertInput
	if err := c.BindJSON(&newAdvert); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, inputError)
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
