package handler

import (
	"lintaspay/internal/dto"
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

// Create Wallet godoc
// @Summary Create wallet
// @Description Create wallet for authenticated user
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body dto.CreateWalletRequest true "Create Wallet Request"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /wallets [post]
func (h *WalletHandler) Create(c *gin.Context) {
	var req dto.CreateWalletRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	User := c.GetUint("user_id")
	wallet := &domain.Wallet{
		UserID:  User,
		Balance: 0,
	}

	if err := h.UseCase.Create(wallet); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse{
		Message: "wallet created",
	})
}
