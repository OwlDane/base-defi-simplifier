export function Hero() {
  return (
    <section className="text-center mb-16">
      <h1 className="text-5xl font-bold text-gray-900 mb-6">
        DeFi Made <span className="text-base-blue">Simple</span>
      </h1>
      <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
        Experience the power of decentralized finance without the complexity. 
        Built on Base L2 for fast, secure, and affordable transactions.
      </p>
      <div className="flex justify-center space-x-4">
        <button className="bg-base-blue text-white px-8 py-3 rounded-lg font-semibold hover:bg-base-blue-dark transition-colors">
          Start Trading
        </button>
        <button className="border border-gray-300 text-gray-700 px-8 py-3 rounded-lg font-semibold hover:bg-gray-50 transition-colors">
          Learn More
        </button>
      </div>
    </section>
  )
}