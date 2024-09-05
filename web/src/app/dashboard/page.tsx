import { Dashboard } from "@/components/dashboard"
import { auth } from "@/lib/auth"

export default async function DashboardPage() {
  const session = await auth()
  if (!session) return <div>Not authenticated</div>

  return <Dashboard />
}