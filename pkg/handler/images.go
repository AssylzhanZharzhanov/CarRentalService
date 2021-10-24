package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path"
	"path/filepath"
)

const staticFileHost = "http://161.35.196.24/static/"
//const staticFileHost = "http://localhost/static"

func (h *Handler) uploadImage(c *gin.Context) {
	advertId := c.Query("advertId")

	file, err := c.FormFile("image")
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	fileName := filepath.Base(file.Filename)
	newFileName := uuid.New().String() + fileName
	dst := path.Join("./static", newFileName)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	imageUrl := staticFileHost + newFileName
	if err := h.service.UploadImage(c.Request.Context(), advertId, imageUrl); err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) UploadMultipleImages(c *gin.Context) {
	//id, _ := (c.Query("advertId")
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	files := form.File["images[]"]

	var fileNames []string
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		fileNames = append(fileNames, filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}


}

func (h *Handler) getImageById(c *gin.Context)  {
	id := c.Query("id")

	c.JSON(http.StatusOK, id)
}

func (h *Handler) deleteImage(c *gin.Context)  {
	imageId := c.Query("imageId")
	advertId := c.Query("advertId")

	err := h.service.DeleteImage(c.Request.Context(), imageId, advertId)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}