import React, { useState } from 'react'
import { Button } from '@/components/ui/button'
import {
  FileIcon,
  GlobeIcon,
  Loader2,
  Package2Icon,
  Rocket,
} from 'lucide-react'
import { Separator } from '@/components/ui/separator'
import type { FC } from 'react'
import { Tab } from '@/types/uiState'
import { ProfileSelect } from '@/features/profile-select'
import { toast } from 'sonner'

interface SidebarProps {
  profiles: { label: string; value: string }[]
  setProfile: (profile: string | null) => void
  profile: string | null
  selectedTab?: Tab
  setSelectedTab?: (tab: Tab) => void
}

const Sidebar: FC<SidebarProps> = ({
  selectedTab,
  setSelectedTab,
  profile,
  profiles,
  setProfile,
}) => {
  const [isLoadingGame, setLoadingGame] = useState(false)

  const runGame = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault() // Preventing default button click behavior

    if (!profile) {
      toast('⚠️ No profile selected', {
        description: 'Please select a profile to run the game.',
      })
      return
    }

    if (isLoadingGame) {
      return
    }

    window
      .runGame(profile)
      .then(() => {
        setLoadingGame(false)
        toast('🏃‍♀️💨 Started!', {
          description: `Running profile ${profile}`,
        })
      })
      .catch((out: string) => {
        setLoadingGame(false)
        toast('🤕 Whoops!', {
          description: `Something went wrong: ${out}`,
        })
      })
  }

  return (
    <div className="bg-[#09090b] hidden border-r border-[#27272a] lg:block">
      <div className="flex h-full max-h-screen flex-col gap-2">
        <div className="flex h-[60px] justify-center border-b border-[#27272a] px-6">
          <div className="flex items-center gap-2 font-semibold">
            <Package2Icon color="white" className="h-6 w-6" />
            <span className="text-white">Lethal Mod Manager</span>
          </div>
        </div>
        <div className="flex-1 overflow-auto py-2">
          <nav className="grid grid-cols-1 justify-center min-w-full text-sm font-medium gap-2 px-4">
            <div className="flex justify-between gap-2 min-w-full">
              <ProfileSelect
                profiles={profiles}
                setProfile={setProfile}
                profile={profile}
              />
              <Button
                className="bg-[#4BC732] hover:bg-[#6BD420] border-none text-white hover:text-white"
                onClick={runGame}
                variant="outline"
              >
                {isLoadingGame ? (
                  <Loader2 size={16} className="animate-spin" />
                ) : (
                  <Rocket size={16} strokeWidth={2} />
                )}
              </Button>
            </div>

            <Separator className="w-full my-2 bg-[#27272a]" />
            <Button
              variant="ghost"
              className={`w-full text-white justify-start pl-3 hover:bg-[#27272a] hover:text-white ${
                selectedTab === 'local-mods' ? 'bg-[#18181B]' : ''
              }`}
              onClick={() => setSelectedTab && setSelectedTab('local-mods')}
            >
              <FileIcon className="mr-2 h-4 w-4" />
              Local mods
            </Button>

            <Button
              variant="ghost"
              className={`w-full text-white justify-start pl-3 hover:bg-[#27272a] hover:text-white ${
                selectedTab === 'global-mods' ? 'bg-[#18181B]' : ''
              }`}
              onClick={() => setSelectedTab && setSelectedTab('global-mods')}
            >
              <GlobeIcon className="mr-2 h-4 w-4" />
              Global mods
            </Button>
          </nav>
        </div>
      </div>
    </div>
  )
}

export default Sidebar
