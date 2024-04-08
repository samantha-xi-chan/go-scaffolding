package handler

import (
	"github.com/gin-gonic/gin"
	"go-scaffolding/internal/app01/service"
	"gorm.io/gorm"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler(db *gorm.DB) *Handler {
	userService := service.NewUserService(db)

	return &Handler{
		userService: userService,
	}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {

	// v1 user
	userV1 := r.Group("api/v1/user")
	{
		userV1.POST("", h.postUser)
		userV1.GET("/:id", h.getUserByID)
		userV1.PUT("/:id", h.putUserByID)
	}
}
