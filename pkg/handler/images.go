package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) UploadImages(c *gin.Context) {

	//id := c.Param("id")

	form, _ := c.MultipartForm()
	files := form.File["images[]"]

	var fileNames []string
	for _, file := range files {
		log.Println(file.Filename)
		fileNames = append(fileNames, file.Filename)
		//c.SaveUploadedFile(file, dst)
	}



	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func generateUrl(string) string {

	return ""
}

func (h *Handler) getImageById(c *gin.Context)  {
	id := c.Query("id")

	c.JSON(http.StatusOK, id)
}