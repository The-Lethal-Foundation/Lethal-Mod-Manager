import React, { useEffect } from 'react'
import { useGetGlobalMods } from '@/hooks/getGlobalMods'
import { GlobalModCard } from './mod-card-global'
import { ordering, section } from '@/types/global'

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
} from '@/components/ui/select'
import { Button } from '@/components/ui/button'
import { toast } from 'sonner'

interface GlobalModListProps {
  ordering: ordering
  setOrdering: (o: ordering) => void
  sectionState: section
  setSection: (s: section) => void
  query: string
  page: number
  profile: string | null
}

const GlobalModList: React.FC<GlobalModListProps> = ({
  ordering,
  sectionState,
  setSection,
  query,
  page,
  profile,
}) => {
  const { globalMods, isLoading, error } = useGetGlobalMods(
    page,
    ordering,
    sectionState,
    query,
  )

  const installFromUrl = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault() // Preventing default button click behavior
    if (!profile) {
      return
    }

    const url = window.prompt('Enter the url to install the mod from')

    if (!url || url.length < 1 || url === null) {
      toast('🤕 Whoops!', {
        description: 'No url provided',
      })
    }

    toast('🧙‍♂️ Installing mod...', {
      description: `${url}`,
      duration: 0,
    })

    window
      .installModFromUrl(profile, url!)
      .then(() => {
        toast('✅ installed', {
          description: `${url} installed!`,
        })
      })
      .catch((out: string) => {
        toast('🤕 Whoops!', {
          description: `Something went wrong: ${out}`,
        })
      })
  }

  useEffect(() => {
    if (!isLoading || error) {
      if (error) {
        console.error(error)
      }

      if (globalMods) {
        console.log('Loaded global mod list:', globalMods.length)
      }
    }
  }, [isLoading, error, globalMods])

  return (
    <>
      <div className="flex gap-4 ml-4 mt-4 md:ml-6 md:mt-6 items-center">
        <Button variant="secondary" onClick={installFromUrl}>
          Install from url
        </Button>

        <Select
          defaultValue="mods"
          onValueChange={(value) => {
            setSection(value as section)
          }}
        >
          <SelectTrigger className="w-[180px] text-white">
            {sectionState.charAt(0).toUpperCase() + sectionState.slice(1)}
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectLabel>Section</SelectLabel>
              <SelectItem value="mods">Mods</SelectItem>
              <SelectItem value="asset-replacements">
                Asset Replacements
              </SelectItem>
              <SelectItem value="libraries">Libraries</SelectItem>
              <SelectItem value="modpacks">Modpacks</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>

      <div className="w-full grid gap-6 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-white p-4 md:p-6">
        {globalMods && globalMods.length > 0 ? (
          globalMods.map((mod, index) => (
            <GlobalModCard
              key={index}
              mod={mod}
              image={mod.mod_picture}
              profile={profile}
            />
          ))
        ) : (
          <p className="">🌍 No global mods found</p>
        )}
      </div>
    </>
  )
}

export { GlobalModList }
