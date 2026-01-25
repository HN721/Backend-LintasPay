package dto

type RegisterRequest struct {
	Name     string `json:"name" `
	Email    string `json:"email" `
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Status  bool        `json:"status" example:"true"`
	Message string      `json:"message" example:"Successfully Login"`
	Token   string      `json:"token"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"email already registered"`
}
type CreateWalletRequest struct {
	UserID uint `json:"user_id" example:"1"`
}

type HistoryRequest struct {
	WalletID uint `json:"wallet_id" example:"1" binding:"required"`
}

type TransactionSuccessResponse struct {
	Message   string `json:"message"`
	Reference string `json:"reference" `
}

type HistoryResponse struct {
	Success string      `json:"success"`
	Data    interface{} `json:"data"`
}
