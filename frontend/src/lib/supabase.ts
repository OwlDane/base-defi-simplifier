import { createClient } from '@supabase/supabase-js'

const supabaseUrl = process.env.NEXT_PUBLIC_SUPABASE_URL!
const supabaseAnonKey = process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!

export const supabase = createClient(supabaseUrl, supabaseAnonKey)

export function createSupabaseClient() {
  return supabase
}

// Types for database tables
export interface User {
  id: string
  wallet_address: string
  created_at: string
  updated_at: string
}

export interface Transaction {
  id: string
  user_id: string
  tx_hash: string
  type: 'swap' | 'liquidity' | 'stake'
  amount_in: string
  amount_out: string
  token_in: string
  token_out: string
  status: 'pending' | 'completed' | 'failed'
  created_at: string
}