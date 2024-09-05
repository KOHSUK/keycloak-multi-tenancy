import { Button } from "@/components/ui/button"
import Link from "next/link"
import { PlusIcon } from "lucide-react"

export function Header() {
  return (
    <header className="sticky top-0 z-10 flex items-center justify-between bg-background px-4 py-4 shadow-sm sm:px-6">
      <Link href="#" className="text-2xl font-bold" prefetch={false}>
        Design Dashboard
      </Link>
      <Button variant="default" className="w-32 rounded-lg">
        <PlusIcon className="mr-2 h-4 w-4" />
        Create new
      </Button>
    </header>
  )
}