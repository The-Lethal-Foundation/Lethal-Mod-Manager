import React, { useEffect, useState } from 'react'
import type { FC } from 'react'
import { Input } from '../components/ui/input'
import type { Tab } from '@/types/uiState'
import { Button } from '@/components/ui/button'
import { Settings } from 'lucide-react'

import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from '@/components/ui/sheet'
import { Label } from '@radix-ui/react-dropdown-menu'
import { toast } from 'sonner'

export interface HeaderProps {
  selectedProfile: string | null
  setSelectedProfile: (p: string) => void
  refetchProfiles: () => void
  currentTab: Tab
  setLocalModQuery: (q: string) => void
  setGlobalModQuery: (q: string) => void
}

const Header: FC<HeaderProps> = ({
  selectedProfile,
  setSelectedProfile,
  refetchProfiles,
  currentTab,
  setLocalModQuery,
  setGlobalModQuery,
}) => {
  const [input, setInput] = useState<string>('')

  // Settings
  const [editedProfile, setEditedProfile] = useState<string | null>(null)

  useEffect(() => {
    setEditedProfile(selectedProfile)
  }, [selectedProfile])

  useEffect(() => {
    setInput('')
  }, [currentTab])

  useEffect(() => {
    const timeoutId = setTimeout(() => {
      // Update the query based on the currentTab
      if (currentTab === 'local-mods') {
        setLocalModQuery(input)
      } else {
        setGlobalModQuery(input)
      }
    }, 500)

    return () => clearTimeout(timeoutId) // Cleanup the timeout
  }, [input, currentTab, setLocalModQuery, setGlobalModQuery])

  const saveChanges = () => {
    if (!editedProfile) return
    if (!selectedProfile) return

    // Rename the profile
    if (editedProfile === selectedProfile) return
    if (!editedProfile.match(/^[a-zA-Z0-9_ ]+$/)) return
    window.renameProfile(selectedProfile, editedProfile).then(() => {
      refetchProfiles()
      setSelectedProfile(editedProfile)
      toast('🧙‍♂️ Profile renamed!', {
        description: `Your profile has been renamed to ${editedProfile}`,
      })
    })
  }

  return (
    <header className="bg-[#09090b] flex h-14 lg:h-[60px] items-center gap-4 border-b border-[#27272a] px-4">
      <Sheet>
        <SheetContent className="bg-[#09090b] border-[#27272a] text-white">
          <SheetHeader>
            <SheetTitle className="text-white">⚙️ Settings</SheetTitle>
            {/* <SheetDescription></SheetDescription> */}
          </SheetHeader>
          <div className="grid gap-4 py-4">
            <div className="grid grid-cols-1 items-center gap-2">
              <Label className="">Profile Name</Label>
              <Input
                id="name"
                pattern="^[a-zA-Z0-9_ ]+$"
                value={editedProfile ?? ''}
                onChange={(e) => setEditedProfile(e.target.value)}
                className="col-span-3"
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4"></div>
          </div>
          <SheetFooter>
            <SheetClose asChild>
              <Button onClick={saveChanges}>Save changes</Button>
            </SheetClose>
          </SheetFooter>
        </SheetContent>

        <div className="w-full flex-1">
          <form>
            <div className="relative flex justify-between gap-4">
              <Input
                value={input}
                onChange={(e) => setInput(e.target.value)}
                type="search"
                placeholder="Search mods..."
                className="w-full h-10 text-white border-none focus:border-white"
              />

              <SheetTrigger asChild>
                <Button className="hover:bg-[#27272a]">
                  <Settings size={16} />
                </Button>
              </SheetTrigger>
            </div>
          </form>
        </div>
      </Sheet>
    </header>
  )
}

export default Header
