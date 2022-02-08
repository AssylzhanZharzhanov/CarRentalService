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

	router.MaxMultipartMemory = 100 << 20

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/code", h.generateCode)
		auth.POST("/code/verify", h.verifyCode)

		admin := auth.Group("/admin")
		{
			admin.POST("/sign-in", h.adminSignIn)
		}
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

		adverts := api.Group("/adverts")
		{
			adverts.POST("/", h.GetUserIdentity, h.createAdvert)
			adverts.GET("/", h.getAllAdverts)
			adverts.GET("/:id", h.getAdvertById)
			adverts.GET("/top", h.getTopAdverts)
			adverts.GET("/similar", h.getSimilarAdverts)
			adverts.GET("/my", h.getUserAdverts)
			adverts.PUT("/:id", h.GetUserIdentity, h.updateAdvert)
			adverts.DELETE("/:id", h.GetUserIdentity, h.deleteAdvert)

			usersAdverts := adverts.Group("/users", h.GetUserIdentity)
			{
				usersAdverts.GET("/active", h.getUserActiveAdverts)
				usersAdverts.GET("/archive", h.getUserArchiveAdverts)
				usersAdverts.GET("/moderation", h.getUserModerationAdverts)
			}
		}

		advertisements := api.Group("/", h.GetUserIdentity)
		{
			advertisements.POST("/", h.addAdvertisement)
			advertisements.GET("/", h.getAdvertisements)
			advertisements.GET("/:id", h.getAdvertisementByID)
			advertisements.PUT("/:id", h.updateAdvertisement)
			advertisements.DELETE("/:id", h.deleteAdvertisement)
			advertisements.POST("/connect", h.connectAdvertisement)
			advertisements.PUT("/connect", h.updateConnectedAdvertisement)
		}

		feedback := api.Group("/feedback", h.GetUserIdentity)
		{
			feedback.POST("/", h.addFeedback)
			feedback.GET("/:id",h.getFeedback)
			feedback.PUT("/:id", h.updateFeedback)
			feedback.DELETE("/:id", h.deleteFeedback)
		}

		bookmark := api.Group("/bookmarks", h.GetUserIdentity)
		{
			bookmark.POST("/", h.addBookmark)
			bookmark.GET("/", h.getBookmarks)
			bookmark.DELETE("/:id",  h.deleteBookmark)
		}

		filters := api.Group("/filters")
		{
			categories := filters.Group("/categories")
			{
				categories.GET("", h.getCategories)
				categories.POST("", h.addCategory)
				categories.DELETE("", h.deleteCategory)
			}

			cities := filters.Group("/cities")
			{
				cities.GET("", h.getCities)
				cities.POST("", h.addCity)
				cities.DELETE("/:id", h.deleteCity)
			}

			rentTypes := filters.Group("/rent_types")
			{
				rentTypes.GET("", h.getRentTypes)
				rentTypes.POST("", h.addRentType)
				rentTypes.DELETE("", h.deleteRentType)
			}

			prices := filters.Group("/prices")
			{
				prices.GET("", h.getPrice)
				prices.POST("", h.addPrice)
				prices.DELETE("", h.deletePrice)
			}

			statuses := filters.Group("/statuses")
			{
				statuses.POST("", h.addStatus)
				statuses.GET("", h.getStatuses)
				statuses.PUT("", h.updateStatus)
			}
		}

		images := api.Group("/images", h.GetUserIdentity)
		{
			images.POST("/", h.uploadImage)
			images.GET("/:id", h.getImageById)
			images.DELETE("/", h.deleteImage)
		}

		search := api.Group("/search")
		{
			search.GET("/autocomplete", h.searchCarMark)
			search.GET("/adverts", h.getSearchAdverts)
		}
	}

	return router
}
