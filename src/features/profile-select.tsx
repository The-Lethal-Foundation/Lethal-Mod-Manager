import * as React from 'react'
import { CaretSortIcon, CheckIcon } from '@radix-ui/react-icons'

import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from '@/components/ui/command'
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover'
import { PlusCircle } from 'lucide-react'

interface ProfileSelectProps {
  profiles: { label: string; value: string }[]
  profile: string | null
  setProfile: (profile: string | null) => void
  refetchProfiles: () => void
}

export function ProfileSelect({
  profiles,
  profile,
  setProfile,
  refetchProfiles,
}: ProfileSelectProps) {
  const [open, setOpen] = React.useState(false)
  const [value, setValue] = React.useState('')

  React.useEffect(() => {
    if (profile && profile.length > 0) {
      setValue(profile)
      window.saveLastUsedProfile(profile)
    }
  }, [profile])

  const createNewProfile = () => {
    const profileName = window.prompt('Enter a name for the new profile')
    if (!profileName) return

    window.createProfile(profileName).then(() => {
      refetchProfiles()
      setValue(profileName)
      setProfile(profileName)
    })
  }

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="default"
          role="combobox"
          aria-expanded={open}
          className="min-w-[75%] max-w-[75%] justify-between hover:bg-[#27272a]"
        >
          {value
            ? profiles.find((profile) => profile.value === value)?.label
            : 'Select profile...'}
          <CaretSortIcon className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[185px] p-0">
        <Command>
          <CommandInput placeholder="Search profile..." className="h-9" />
          <CommandItem className="flex h-full items-center border-b px-3 py-1 bg-gray-50 hover:cursor-pointer hover:opacity-80 duration-300">
            <span
              className="flex items-center h-full"
              onClick={createNewProfile}
            >
              <PlusCircle className="mr-2 h-3.5 w-3.5" />
              New Profile
            </span>
          </CommandItem>
          <CommandEmpty
            firstRender
            className="flex flex-col gap-4 items-center"
          >
            🙊 No profiles found
          </CommandEmpty>
          <CommandGroup>
            {profiles.map((profile) => {
              return (
                <CommandItem
                  className="hover:cursor-pointer"
                  value={profile.value}
                  key={profile.value}
                  onSelect={() => {
                    setValue(profile.value === value ? '' : profile.value)
                    setProfile(profile.value === value ? '' : profile.value)
                    setOpen(false)
                  }}
                >
                  {profile.label}
                  <CheckIcon
                    className={cn(
                      'ml-auto h-4 w-4',
                      value === profile.value ? 'opacity-100' : 'opacity-0',
                    )}
                  />
                </CommandItem>
              )
            })}
          </CommandGroup>
        </Command>
      </PopoverContent>
    </Popover>
  )
}
