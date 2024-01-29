import React from 'react'
import { ProfileSelect } from '@/features/profile-select'
import { Button } from '@/components/ui/button'
import { FileIcon, GlobeIcon, Package2Icon } from 'lucide-react'
import { Separator } from '@/components/ui/separator'
import type { FC } from 'react'

interface SidebarProps {
  profiles: { label: string; value: string }[]
  setProfile: (profile: string | null) => void
  profile: string | null
}

const Sidebar: FC<SidebarProps> = ({ profiles, setProfile, profile }) => {
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
          <nav className="grid justify-center px-4 text-sm font-medium gap-2">
            <ProfileSelect
              profiles={profiles}
              setProfile={setProfile}
              profile={profile}
            />
            <Separator className="my-2 bg-[#27272a]" />

            <Button variant="link" className="text-white justify-start pl-3">
              <FileIcon className="mr-2 h-4 w-4" />
              Local mods
            </Button>

            <Button variant="link" className="text-white justify-start pl-3">
              <GlobeIcon className="mr-2 h-4 w-4" />
              Online mods
            </Button>
          </nav>
        </div>
      </div>
    </div>
  )
}

export default Sidebar
