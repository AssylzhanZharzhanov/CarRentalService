package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "gitlab.com/zharzhanov/region/docs"
	"gitlab.com/zharzhanov/region/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(CORSMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/code", h.generateCode)
		auth.POST("/code/verify", h.verifyCode)
	}

	api := router.Group("/api")
	{
		users := api.Group("/users", h.GetUserIdentity)
		{
			users.POST("/", h.createUser)
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id", h.deleteUser)
		}

		adverts := api.Group("/adverts", h.GetUserIdentity)
		{
			adverts.POST("/", h.createAdvert)
			adverts.GET("/", h.getAllAdverts)
			adverts.GET("/:id", h.getAdvertById)
			adverts.PUT("/:id", h.updateAdvert)
			adverts.DELETE("/:id", h.deleteAdvert)
		}

		feedback := api.Group("/feedback", h.GetUserIdentity)
		{
			feedback.POST("/", h.addFeedback)
			feedback.GET("/:id",h.getFeedback)
		}

		bookmark := api.Group("/bookmark", h.GetUserIdentity)
		{
			bookmark.POST("/", h.addBookmark)
			bookmark.GET("/", h.getBookmarks)
		}

		//filters := api.Group("/")
		//{
		//	filters.GET("/")
		//}

		images := api.Group("/images", h.GetUserIdentity)
		{
			images.POST("/", h.uploadImage)
			images.GET("/:id", h.getImageById)
			images.DELETE("/", h.deleteImage)
		}
	}

	return router
}
