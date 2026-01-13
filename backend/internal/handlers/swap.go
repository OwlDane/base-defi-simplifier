package handlers

import (
	"base-defi-api/internal/models"
	"base-defi-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type SwapHandler struct {
	blockchainService *services.BlockchainService
	supabaseService   *services.SupabaseService
}

func NewSwapHandler(blockchainService *services.BlockchainService, supabaseService *services.SupabaseService) *SwapHandler {
	return &SwapHandler{
		blockchainService: blockchainService,
		supabaseService:   supabaseService,
	}
}

func (h *SwapHandler) GetQuote(c *fiber.Ctx) error {
	var req struct {
		TokenIn  string `json:"tokenIn"`
		TokenOut string `json:"tokenOut"`
		AmountIn string `json:"amountIn"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ApiResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	quote, err := h.blockchainService.GetSwapQuote(req.TokenIn, req.TokenOut, req.AmountIn)
	if err != nil {
		return c.Status(500).JSON(models.ApiResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(models.ApiResponse{
		Success: true,
		Data:    quote,
	})
}

func (h *SwapHandler) ExecuteSwap(c *fiber.Ctx) error {
	var req struct {
		Quote       models.SwapQuote `json:"quote"`
		UserAddress string           `json:"userAddress"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ApiResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	txHash, err := h.blockchainService.ExecuteSwap(&req.Quote, req.UserAddress)
	if err != nil {
		return c.Status(500).JSON(models.ApiResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(models.ApiResponse{
		Success: true,
		Data: fiber.Map{
			"txHash": txHash,
		},
	})
}

func (h *SwapHandler) GetSwapHistory(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(400).JSON(models.ApiResponse{
			Success: false,
			Error:   "Wallet address is required",
		})
	}

	// Get user first
	user, err := h.supabaseService.GetUser(address)
	if err != nil {
		return c.Status(404).JSON(models.ApiResponse{
			Success: false,
			Error:   "User not found",
		})
	}

	transactions, err := h.supabaseService.GetTransactionHistory(user.ID)
	if err != nil {
		return c.Status(500).JSON(models.ApiResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(models.ApiResponse{
		Success: true,
		Data:    transactions,
	})
}