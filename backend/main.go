package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"base-defi-api/internal/config"
	"base-defi-api/internal/handlers"
	"base-defi-api/internal/middleware"
	"base-defi-api/internal/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize configuration
	cfg := config.New()

	// Initialize services
	supabaseService := services.NewSupabaseService(cfg.SupabaseURL, cfg.SupabaseKey)
	blockchainService := services.NewBlockchainService(cfg.BaseRPCURL)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(supabaseService)
	swapHandler := handlers.NewSwapHandler(blockchainService, supabaseService)
	poolHandler := handlers.NewPoolHandler(blockchainService)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AllowedOrigins,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "Base DeFi API is running",
		})
	})

	// API routes
	api := app.Group("/api/v1")

	// User routes
	users := api.Group("/users")
	users.Post("/", userHandler.CreateUser)
	users.Get("/:address", userHandler.GetUser)

	// Swap routes
	swaps := api.Group("/swaps")
	swaps.Post("/quote", swapHandler.GetQuote)
	swaps.Post("/execute", swapHandler.ExecuteSwap)
	swaps.Get("/history/:address", swapHandler.GetSwapHistory)

	// Pool routes
	pools := api.Group("/pools")
	pools.Get("/", poolHandler.GetPools)
	pools.Get("/:id", poolHandler.GetPool)
	pools.Get("/user/:address", poolHandler.GetUserPositions)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}