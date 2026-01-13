export interface Token {
  address: string
  symbol: string
  name: string
  decimals: number
  logoURI?: string
}

export interface SwapQuote {
  tokenIn: Token
  tokenOut: Token
  amountIn: string
  amountOut: string
  priceImpact: number
  fee: string
  route: string[]
}

export interface Pool {
  id: string
  token0: Token
  token1: Token
  fee: number
  tvl: number
  apr: number
}

export interface UserPosition {
  id: string
  pool: Pool
  liquidity: string
  token0Amount: string
  token1Amount: string
  uncollectedFees: {
    token0: string
    token1: string
  }
}

export interface ApiResponse<T> {
  success: boolean
  data?: T
  error?: string
}