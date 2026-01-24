package handler

import (
	"lintaspay/internal/modules/users/usecase"
	"lintaspay/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UseCase usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		UseCase: uc,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UseCase.Register(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered successfully"})
}
func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UseCase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, err := jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Successfully Login",
		"token":   token,
		"data":    user,
	})
}
