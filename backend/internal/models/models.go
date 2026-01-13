package models

import "time"

type User struct {
	ID            string    `json:"id" db:"id"`
	WalletAddress string    `json:"wallet_address" db:"wallet_address"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type Transaction struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	TxHash    string    `json:"tx_hash" db:"tx_hash"`
	Type      string    `json:"type" db:"type"`
	AmountIn  string    `json:"amount_in" db:"amount_in"`
	AmountOut string    `json:"amount_out" db:"amount_out"`
	TokenIn   string    `json:"token_in" db:"token_in"`
	TokenOut  string    `json:"token_out" db:"token_out"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Token struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	LogoURI  string `json:"logoURI,omitempty"`
}

type SwapQuote struct {
	TokenIn     Token    `json:"tokenIn"`
	TokenOut    Token    `json:"tokenOut"`
	AmountIn    string   `json:"amountIn"`
	AmountOut   string   `json:"amountOut"`
	PriceImpact float64  `json:"priceImpact"`
	Fee         string   `json:"fee"`
	Route       []string `json:"route"`
}

type Pool struct {
	ID     string  `json:"id"`
	Token0 Token   `json:"token0"`
	Token1 Token   `json:"token1"`
	Fee    int     `json:"fee"`
	TVL    float64 `json:"tvl"`
	APR    float64 `json:"apr"`
}

type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}