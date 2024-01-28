import React from 'react'

import { ModCard } from '@/features/mod-card'
import useGetMods from '@/hooks/getMods'
import { useEffect } from 'react'

interface ModListProps {
  profile: string
}

export const ModList: React.FC<ModListProps> = ({ profile }) => {
  const { mods } = useGetMods(profile)

  // const mods = testModList

  useEffect(() => {
    console.log('New profile:', profile)
    console.log('New mods:', mods)
  }, [profile])

  return (
    <div className="w-full grid gap-6 md:grid-cols-2 lg:grid-cols-3 text-white">
      {mods.map((mod, key) => (
        <ModCard key={key} mod={mod} />
      ))}
    </div>
  )
}
