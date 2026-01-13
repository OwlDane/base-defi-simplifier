import { Zap, Shield, DollarSign, Users } from 'lucide-react'

const features = [
  {
    icon: Zap,
    title: 'Lightning Fast',
    description: 'Execute trades in seconds with Base L2 technology'
  },
  {
    icon: Shield,
    title: 'Secure & Reliable',
    description: 'Built on battle-tested smart contracts and audited protocols'
  },
  {
    icon: DollarSign,
    title: 'Low Fees',
    description: 'Minimal transaction costs compared to Ethereum mainnet'
  },
  {
    icon: Users,
    title: 'User Friendly',
    description: 'Intuitive interface designed for both beginners and pros'
  }
]

export function Features() {
  return (
    <section className="grid md:grid-cols-2 lg:grid-cols-4 gap-8 mb-16">
      {features.map((feature, index) => (
        <div key={index} className="bg-white p-6 rounded-xl shadow-sm border border-gray-100">
          <feature.icon className="w-12 h-12 text-base-blue mb-4" />
          <h3 className="text-xl font-semibold text-gray-900 mb-2">
            {feature.title}
          </h3>
          <p className="text-gray-600">
            {feature.description}
          </p>
        </div>
      ))}
    </section>
  )
}