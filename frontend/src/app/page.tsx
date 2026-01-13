import { WalletConnect } from '@/components/wallet-connect'
import { Hero } from '@/components/hero'
import { Features } from '@/components/features'

export default function Home() {
  return (
    <main className="min-h-screen bg-gradient-to-br from-base-blue-light to-white">
      <div className="container mx-auto px-4 py-8">
        <header className="flex justify-between items-center mb-12">
          <div className="flex items-center space-x-2">
            <div className="w-8 h-8 bg-base-blue rounded-full"></div>
            <h1 className="text-2xl font-bold text-gray-900">Base DeFi</h1>
          </div>
          <WalletConnect />
        </header>
        
        <Hero />
        <Features />
      </div>
    </main>
  )
}