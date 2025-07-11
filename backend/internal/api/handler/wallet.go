package handler

import (
	"mpc/internal/model"
	"mpc/internal/service"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	BaseHandler
	walletService *service.WalletService
}

func NewWalletHandler(walletService *service.WalletService) *WalletHandler {
	return &WalletHandler{
		BaseHandler:   NewBaseHandler(),
		walletService: walletService,
	}
}

// GetBalance godoc
// @Summary      Get wallet balance
// @Description  Get wallet balance
// @Tags         wallets
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Response{payload=model.WalletResponse}
// @Failure      400  {object}  model.ErrorResponse
// @Failure      401  {object}  model.ErrorResponse
// @Router       /wallet [get]
func (h *WalletHandler) GetBalance(c *gin.Context) {
	userID, err := h.GetUserID(c)
	if err != nil {
		c.Error(err)
		return
	}

	var req model.GetBalanceRequest
	res, err := h.walletService.GetBalanceByAddress(c.Request.Context(), req, userID)
	if err != nil {
		c.Error(err)
		return
	}
	h.SuccessResponse(c, res)
}
