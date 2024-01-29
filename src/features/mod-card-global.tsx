import React from 'react'
import type { FC } from 'react'

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import type { GlobalModView } from '@/types/mod'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { DownloadIcon, LinkIcon } from 'lucide-react'
import { toast } from 'sonner'

interface GlobalModCardProps {
  profile: string | null
  mod: GlobalModView
  image?: string
}

export const GlobalModCard: FC<GlobalModCardProps> = ({
  profile,
  mod,
  image = 'https://generated.vusercontent.net/placeholder.svg',
}) => {
  const visitMod = () => {
    window.open(
      `https://thunderstore.io/c/lethal-company/p/${mod.mod_author}/${mod.mod_name.replace(/\s/g, '_')}`,
    )
  }

  const installMod = () => {
    toast('🧙‍♂️ Installing mod...', {
      description: `${mod.mod_name} by ${mod.mod_author}`,
      duration: 0,
    })

    if (!profile) {
      toast('🤕 Whoops!', {
        description: 'No profile selected',
      })
      return
    }

    window
      .installMod(profile, mod.mod_author, mod.mod_name.replace(/\s/g, '_'))
      .then(() => {
        toast('✅', {
          description: `${mod.mod_name} installed!`,
        })
      })
      .catch((out: string) => {
        toast('🤕 Whoops!', {
          description: `Something went wrong: ${out}`,
        })
      })
  }

  return (
    <>
      <DropdownMenu>
        <DropdownMenuContent>
          <DropdownMenuLabel>Mod action</DropdownMenuLabel>
          <DropdownMenuSeparator />

          <DropdownMenuItem
            className="hover:cursor-pointer"
            onClick={installMod}
          >
            <DownloadIcon className="w-4 h-4 mr-2" />
            Install
          </DropdownMenuItem>

          <DropdownMenuItem className="hover:cursor-pointer" onClick={visitMod}>
            <LinkIcon className="w-4 h-4 mr-2" />
            Visit
          </DropdownMenuItem>
        </DropdownMenuContent>
        <DropdownMenuTrigger asChild>
          <Card className="bg-[#09090b] border-none hover:scale-110 duration-300 hover:cursor-pointer">
            <CardHeader>
              <img
                src={image}
                alt="Mod #1"
                className="w-full h-48 object-cover rounded"
              />
            </CardHeader>
            <CardContent>
              <CardTitle className="text-white">{mod.mod_name}</CardTitle>
              <CardDescription className="truncate text-xs mt-1">
                By {mod.mod_author}
              </CardDescription>
            </CardContent>
          </Card>
        </DropdownMenuTrigger>
      </DropdownMenu>
    </>
  )
}
