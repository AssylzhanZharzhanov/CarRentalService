package handler

import "github.com/gin-gonic/gin"

type Handler struct {

}

func (h *Handler) InitRoutes() *gin.Engine  {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/")
			users.GET("/")
			users.GET("/:id")
			users.PUT("/:id")
			users.DELETE("/:id")
		}

		adverts := api.Group("/adverts")
		{
			adverts.POST("/")
			adverts.GET("/")
			adverts.GET("/:id")
			adverts.PUT("/:id")
			adverts.DELETE("/:id")
		}
	}

	return router
}