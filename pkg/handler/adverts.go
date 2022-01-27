package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type Filter struct {
	City     string `json:"city" bson:"city,omitempty"`
	Category string `json:"category" bson:"category,omitempty"`
	RentType string `json:"rent_type" bson:"rent_type,omitempty"`
	Price    int    `json:"price" bson:"price,omitempty"`
}

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
		log.Printf("file: %v", file)
		dst := path.Join("./static", filename)
		fileNames = append(fileNames, filename)
		imageUrls = append(imageUrls, staticFileHost + filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	_, err = h.service.CreateAdvert(c.Request.Context(), advert, imageUrls, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, createObjectError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// @Summary Get all adverts
// @Tags adverts
// @Description Get all adverts
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

	if status := c.Query("status"); status != "" {
		query["status"] = status
	}

	//array
	if rentType := c.Query("rent_type"); rentType != "" {
		query["rent_type"] = rentType
	}

	if title := c.Query("title"); title != "" {
		query["$text"] = bson.M{"$search": fmt.Sprintf("\"%s\"", strings.ToLower(title))}
	}

	if c.Query("minPrice") != "" && c.Query("maxPrice") != "" {
		minPrice, _ := strconv.Atoi(c.Query("minPrice"))
		maxPrice, _ := strconv.Atoi(c.Query("maxPrice"))
		query["price"] = bson.M{"$gte": minPrice, "$lte": maxPrice}
	} else if c.Query("minPrice") != "" &&  c.Query("maxPrice") == "" {
		minPrice, _ := strconv.Atoi(c.Query("minPrice"))
		query["price"] = bson.M{"$gte": minPrice}
	} else if c.Query("minPrice") == "" &&  c.Query("maxPrice") != "" {
		maxPrice, _ := strconv.Atoi(c.Query("maxPrice"))
		query["price"] = bson.M{"$lte": maxPrice}
	}

	adverts, err := h.service.GetAllAdverts(c.Request.Context(), query)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, adverts)
}

// @Summary Get advert by id
// @Tags adverts
// @Description Get advert by id
// @ID get-advert-by-id
// @Accept json
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {array} models.Advert
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/adverts/{id} [get]
func (h *Handler) getAdvertById(c *gin.Context) {
	id := c.Param("id")

	advert, err := h.service.GetAdvertById(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, notFoundError)
		return
	}

	c.JSON(http.StatusOK, advert)
}

// @Summary Get user adverts
// @Tags adverts
// @Description Get user adverts
// @ID get-user-adverts
// @Accept json
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {array} models.Advert
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/adverts/{id} [get]
func (h *Handler) getUserAdverts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	adverts, err := h.service.Adverts.GetMyAdverts(c.Request.Context(), userId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	c.JSON(http.StatusOK, adverts)
}

func (h *Handler) getTopAdverts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	adverts, err := h.service.Adverts.GetMyAdverts(c.Request.Context(), userId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	c.JSON(http.StatusOK, adverts)
}

func (h *Handler) getSimilarAdverts(c *gin.Context) {
	title := c.Query("title")
	price, _ := strconv.Atoi(c.Query("price"))

	adverts, err := h.service.Adverts.GetSimilarAdverts(c.Request.Context(), title, price)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "adverts not found")
		return
	}

	c.JSON(http.StatusOK, adverts)
}

// @Summary Get advert by id
// @Tags adverts
// @Description Get advert by id
// @ID get-advert-by-id
// @Accept json
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {json} models.Advert
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/adverts/{id} [put]
func (h *Handler) updateAdvert(c *gin.Context) {
	id := c.Param("id")

	var newAdvert models.AdvertInput
	if err := c.BindJSON(&newAdvert); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, inputError)
		return
	}

	log.Println(newAdvert)

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
