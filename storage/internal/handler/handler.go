package handler

import (
	"github.com/gin-gonic/gin"
	"storage/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRouters() *gin.Engine {
	routes := gin.New()

	auth := routes.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := routes.Group("/api")
	{
		api.POST("/upload", h.UploadFile)
		api.GET("/get/:id", h.DownloadFile)
		api.DELETE("/delete/:id", h.DeleteObject)
	}

	return routes
}
