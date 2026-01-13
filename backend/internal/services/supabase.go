package services

import (
	"base-defi-api/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type SupabaseService struct {
	url string
	key string
}

func NewSupabaseService(url, key string) *SupabaseService {
	return &SupabaseService{
		url: url,
		key: key,
	}
}

func (s *SupabaseService) CreateUser(user *models.User) error {
	// Placeholder implementation
	// In real implementation, you would make HTTP request to Supabase
	fmt.Printf("Creating user: %+v\n", user)
	return nil
}

func (s *SupabaseService) GetUser(walletAddress string) (*models.User, error) {
	// Placeholder implementation
	// In real implementation, you would make HTTP request to Supabase
	fmt.Printf("Getting user: %s\n", walletAddress)
	return &models.User{
		ID:            "1",
		WalletAddress: walletAddress,
	}, nil
}

func (s *SupabaseService) CreateTransaction(tx *models.Transaction) error {
	// Placeholder implementation
	fmt.Printf("Creating transaction: %+v\n", tx)
	return nil
}

func (s *SupabaseService) GetTransactionHistory(userID string) ([]models.Transaction, error) {
	// Placeholder implementation
	fmt.Printf("Getting transaction history for user: %s\n", userID)
	return []models.Transaction{}, nil
}

func (s *SupabaseService) makeRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	var bodyReader *strings.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = strings.NewReader(string(jsonBody))
	}

	url := fmt.Sprintf("%s/rest/v1/%s", s.url, endpoint)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("apikey", s.key)
	req.Header.Set("Authorization", "Bearer "+s.key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}