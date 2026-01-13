package config

import "os"

type Config struct {
	Port           string
	SupabaseURL    string
	SupabaseKey    string
	BaseRPCURL     string
	AllowedOrigins string
	JWTSecret      string
}

func New() *Config {
	return &Config{
		Port:           getEnv("PORT", "8080"),
		SupabaseURL:    getEnv("SUPABASE_URL", ""),
		SupabaseKey:    getEnv("SUPABASE_SERVICE_ROLE_KEY", ""),
		BaseRPCURL:     getEnv("BASE_RPC_URL", "https://mainnet.base.org"),
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "http://localhost:3000"),
		JWTSecret:      getEnv("JWT_SECRET", "your-secret-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}