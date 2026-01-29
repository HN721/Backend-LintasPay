package handler

import (
	"lintaspay/internal/dto"
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

// TopUp godoc
// @Summary Top Up Wallet
// @Description Add balance to wallet
// @Tags Transaction
// @Accept json
// @Produce json
// @Param request body TopUpRequest true "Top Up Request"
// @Success 200 {object} dto.TransactionSuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /trx/top-up [post]
func (h *TransactionHandler) TopUp(c *gin.Context) {
	var req TopUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	//  userID dari JWT
	userID := c.GetUint("user_id")

	ref := refrence.GenerateReference("TOPUP")

	if err := h.UseCase.TopUp(userID, req.Amount, ref); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.TransactionSuccessResponse{
		Message:   "top up success",
		Reference: ref,
	})
}

// Transfer godoc
// @Summary Transfer Balance
// @Description Transfer balance to another user
// @Tags Transaction
// @Accept json
// @Produce json
// @Param request body TransferRequest true "Transfer Request"
// @Success 200 {object} dto.TransactionSuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /trx/transfer [post]
func (h *TransactionHandler) Transfer(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	fromUserID := c.GetUint("user_id")
	ref := refrence.GenerateReference("TRF")

	if err := h.UseCase.Transfer(fromUserID, req.ToUserID, req.Amount, ref); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.TransactionSuccessResponse{
		Message:   "transfer success",
		Reference: ref,
	})
}

// HistoryTransaction godoc
// @Summary Transaction History
// @Description Get transaction history by wallet ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param request body dto.HistoryRequest true "History Request"
// @Success 200 {object} dto.HistoryResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /trx/history [get]
func (h *TransactionHandler) HistoryTransaction(c *gin.Context) {
	var req dto.HistoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	data, err := h.UseCase.History(req.WalletID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.HistoryResponse{
		Success: "Success Get Data",
		Data:    data,
	})
}
