export function Loading() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-base-blue-light to-white flex items-center justify-center">
      <div className="flex flex-col items-center space-y-4">
        <div className="w-12 h-12 bg-base-blue rounded-full animate-pulse"></div>
        <p className="text-gray-600">Loading Base DeFi...</p>
      </div>
    </div>
  )
}