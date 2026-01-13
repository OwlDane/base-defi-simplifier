'use client'

import { OnchainKitProvider } from '@coinbase/onchainkit'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { WagmiProvider, createConfig, http, createStorage, cookieStorage } from 'wagmi'
import { base, baseGoerli } from 'wagmi/chains'
import { coinbaseWallet, walletConnect } from 'wagmi/connectors'
import { useState, useEffect } from 'react'
import { Loading } from './loading'

const config = createConfig({
  chains: [base, baseGoerli],
  storage: createStorage({
    storage: typeof window !== 'undefined' ? window.localStorage : cookieStorage,
  }),
  connectors: [
    coinbaseWallet({ 
      appName: 'Base DeFi Simplifier',
      preference: 'all'
    }),
    ...(typeof window !== 'undefined' ? [
      walletConnect({ 
        projectId: process.env.NEXT_PUBLIC_WALLETCONNECT_PROJECT_ID || 'demo',
        metadata: {
          name: 'Base DeFi Simplifier',
          description: 'Simplifying DeFi on Base L2',
          url: 'http://localhost:3000',
          icons: ['https://avatars.githubusercontent.com/u/37784886']
        }
      })
    ] : []),
  ],
  transports: {
    [base.id]: http(),
    [baseGoerli.id]: http(),
  },
  ssr: true,
})

export function Providers({ children }: { children: React.ReactNode }) {
  const [queryClient] = useState(() => new QueryClient({
    defaultOptions: {
      queries: {
        staleTime: 60 * 1000, // 1 minute
        refetchOnWindowFocus: false,
      },
    },
  }))
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  if (!mounted) {
    return <Loading />
  }

  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        <OnchainKitProvider
          apiKey={process.env.NEXT_PUBLIC_ONCHAINKIT_API_KEY}
          chain={base}
        >
          {children}
        </OnchainKitProvider>
      </QueryClientProvider>
    </WagmiProvider>
  )
}