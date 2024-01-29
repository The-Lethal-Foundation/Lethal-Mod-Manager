import React from 'react'
import type { FC } from 'react'

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Mod } from '@/types/mod'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { FolderIcon, Trash2Icon } from 'lucide-react'
import { toast } from 'sonner'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog'

interface ModCardProps {
  profile: string
  mod: Mod
  image?: string
}

export const ModCard: FC<ModCardProps> = ({
  profile,
  mod,
  image = 'https://generated.vusercontent.net/placeholder.svg',
}) => {
  const [isMounted, setIsMounted] = React.useState(true)

  if (!isMounted) {
    return null
  }

  const openModFolder = () => {
    window
      .openModDir(profile, mod.mod_path_name)
      .then(() => {})
      .catch((err: string) => {
        toast('🤕 Whoops!', {
          description: `Something went wrong: ${err}`,
        })
      })
  }

  const deleteMod = () => {
    setIsMounted(false)
    window
      .deleteMod(profile, mod.mod_path_name)
      .then(() => {})
      .catch((err: string) => {
        toast('🤕 Whoops!', {
          description: `Something went wrong: ${err}`,
        })
      })
  }

  return (
    <>
      <AlertDialog>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
            <AlertDialogDescription>
              This action cannot be undone. This will permanently delete this
              mod.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancel</AlertDialogCancel>
            <AlertDialogAction onClick={deleteMod}>
              Delete mod
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>

        <DropdownMenu>
          <DropdownMenuContent>
            <DropdownMenuLabel>Mod action</DropdownMenuLabel>
            <DropdownMenuSeparator />

            <DropdownMenuItem
              className="hover:cursor-pointer"
              onClick={openModFolder}
            >
              <FolderIcon className="mr-2 h-4 w-4" />
              Open folder
            </DropdownMenuItem>
            <AlertDialogTrigger className="w-full">
              <DropdownMenuItem className="hover:cursor-pointer">
                <Trash2Icon className="mr-2 h-4 w-4" />
                Delete
              </DropdownMenuItem>
            </AlertDialogTrigger>
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
                  By {mod.mod_author} • {mod.mod_version}
                </CardDescription>
              </CardContent>
            </Card>
          </DropdownMenuTrigger>
        </DropdownMenu>
      </AlertDialog>
    </>
  )
}
