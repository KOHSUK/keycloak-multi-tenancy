import React from 'react'
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { signOutAction } from './signout-action'

interface SignOutDialogProps {
  children: React.ReactNode
}

export function SignOutDialog({ children }: SignOutDialogProps) {
  return (
    <Dialog>
      <DialogTrigger asChild>
        {children}
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Are you sure you want to sign out?</DialogTitle>
          <DialogDescription>
            You will be logged out of your account.
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <form action={signOutAction}>
            <Button type="submit" variant="destructive">
              Sign Out
            </Button>
          </form>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}