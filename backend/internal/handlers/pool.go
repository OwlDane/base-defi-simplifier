package handlers

import (
	"base-defi-api/internal/models"
	"base-defi-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type PoolHandler struct {
	blockchainService *services.BlockchainService
}

func NewPoolHandler(blockchainService *services.BlockchainService) *PoolHandler {
	return &PoolHandler{
		blockchainService: blockchainService,
	}
}

func (h *PoolHandler) GetPools(c *fiber.Ctx) error {
	pools, err := h.blockchainService.GetPools()
	if err != nil {
		return c.Status(500).JSON(models.ApiResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(models.ApiResponse{
		Success: true,
		Data:    pools,
	})
}

func (h *PoolHandler) GetPool(c *fiber.Ctx) error {
	poolID := c.Params("id")
	if poolID == "" {
		return c.Status(400).JSON(models.ApiResponse{
			Success: false,
			Error:   "Pool ID is required",
		})
	}

	pool, err := h.blockchainService.GetPool(poolID)
	if err != nil {
		return c.Status(404).JSON(models.ApiResponse{
			Success: false,
			Error:   "Pool not found",
		})
	}

	return c.JSON(models.ApiResponse{
		Success: true,
		Data:    pool,
	})
}

func (h *PoolHandler) GetUserPositions(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(400).JSON(models.ApiResponse{
			Success: false,
			Error:   "Wallet address is required",
		})
	}

	positions, err := h.blockchainService.GetUserPositions(address)
	if err != nil {
		return c.Status(500).JSON(models.ApiResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(models.ApiResponse{
		Success: true,
		Data:    positions,
	})
}