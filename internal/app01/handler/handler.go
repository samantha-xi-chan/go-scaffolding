package handler

import (
	"github.com/gin-gonic/gin"
	"go-scaffolding/internal/app01/service"
	"gorm.io/gorm"
	"net/http"
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
	r.GET("/user/:id", h.getUserByID)
}

func (h *Handler) getUserByID(c *gin.Context) {
	idStr := c.Param("id")

	user, err := h.userService.GetUserByID(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.Id,
		"username": user.Username,
		"email":    user.Email,
	})
}
