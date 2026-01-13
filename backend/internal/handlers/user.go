package handlers

import (
	"base-defi-api/internal/models"
	"base-defi-api/internal/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	supabaseService *services.SupabaseService
}

func NewUserHandler(supabaseService *services.SupabaseService) *UserHandler {
	return &UserHandler{
		supabaseService: supabaseService,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req struct {
		WalletAddress string `json:"wallet_address"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(models.ApiResponse{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	user := &models.User{
		ID:            uuid.New().String(),
		WalletAddress: req.WalletAddress,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := h.supabaseService.CreateUser(user); err != nil {
		return c.Status(500).JSON(models.ApiResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(models.ApiResponse{
		Success: true,
		Data:    user,
	})
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(400).JSON(models.ApiResponse{
			Success: false,
			Error:   "Wallet address is required",
		})
	}

	user, err := h.supabaseService.GetUser(address)
	if err != nil {
		return c.Status(404).JSON(models.ApiResponse{
			Success: false,
			Error:   "User not found",
		})
	}

	return c.JSON(models.ApiResponse{
		Success: true,
		Data:    user,
	})
}