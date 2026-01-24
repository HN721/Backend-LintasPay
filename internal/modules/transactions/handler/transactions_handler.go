package handler

import (
	"lintaspay/internal/modules/transactions/usecase"
	"lintaspay/pkg/refrence"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	UseCase usecase.TransactionUsecase
}

func NewTransactionHandler(uc usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		UseCase: uc,
	}
}

type TopUpRequest struct {
	Amount int64 `json:"amount" binding:"required,min=1"`
}
type TransferRequest struct {
	ToUserID uint  `json:"to_user_id" binding:"required"`
	Amount   int64 `json:"amount" binding:"required,min=1"`
}

func (h *TransactionHandler) TopUp(c *gin.Context) {
	var req TopUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//  userID dari JWT
	userID := c.GetUint("user_id")

	ref := refrence.GenerateReference("TOPUP")

	if err := h.UseCase.TopUp(userID, req.Amount, ref); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "top up success",
		"reference": ref,
	})
}
func (h *TransactionHandler) Transfer(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fromUserID := c.GetUint("user_id")
	ref := refrence.GenerateReference("TRF")

	if err := h.UseCase.Transfer(fromUserID, req.ToUserID, req.Amount, ref); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "transfer success",
		"reference": ref,
	})
}
