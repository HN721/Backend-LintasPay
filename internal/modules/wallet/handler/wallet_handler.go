package handler

import (
	"lintaspay/internal/modules/wallet/domain"
	"lintaspay/internal/modules/wallet/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	UseCase usecase.WalletUsecase
}

func NewWalletHandler(uc usecase.WalletUsecase) *WalletHandler {
	return &WalletHandler{
		UseCase: uc,
	}
}
func (h *WalletHandler) Create(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	wallet := &domain.Wallet{
		UserID:  req.UserID,
		Balance: 0,
	}

	if err := h.UseCase.Create(wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "wallet created",
	})
}
func (h *WalletHandler) Update(c *gin.Context) {
	var req struct {
		ID      uint  `json:"id" binding:"required"`
		Balance int64 `json:"balance" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	wallet := &domain.Wallet{
		ID:      req.ID,
		Balance: req.Balance,
	}

	if err := h.UseCase.Update(wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "wallet updated",
	})
}
