import { LogIn } from "@/components/login"
import { auth } from "@/lib/auth"
import { redirect } from "next/navigation"

export default async function Page() {
  const session = await auth()
  if (session) return redirect("/dashboard")

  return <LogIn />
}
