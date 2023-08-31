package handler

import (
	_ "avito-backend-task-2023/docs"
	"avito-backend-task-2023/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine { // Инициализатор всех энд-поинтов
	router := gin.New()

	usersSegService := router.Group("/users-segmentation")
	{
		segments := usersSegService.Group("/segment")
		{
			segments.POST("/", h.createSegment)
			segments.DELETE("/", h.deleteSegment)
		}

		users := usersSegService.Group("/user")
		{
			users.POST("/", h.createUser)
			users.DELETE("/:id", h.deleteUser)
		}

		usersSegments := usersSegService.Group("/users-segments")
		{
			usersSegments.POST("/logs", h.getUserSegmentsLogs)
			usersSegments.POST("/", h.updateUserSegments)
			usersSegments.GET("/:id", h.getUserSegments)
		}

		usersSegService.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
