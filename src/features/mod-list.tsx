import React from 'react'

import { ModCard } from '@/features/mod-card'
import useGetMods from '@/hooks/getMods'
import { useEffect } from 'react'

interface ModListProps {
  profile: string
}

export const ModList: React.FC<ModListProps> = ({ profile }) => {
  const m = useGetMods(profile)

  useEffect(() => {
    if (!m.isLoading || m.error) {
      console.log('Mods:', m.mods.length, m.mods)
    }
  }, [m.isLoading, m.error, m.mods])

  return (
    <div className="w-full grid gap-6 md:grid-cols-2 lg:grid-cols-3 text-white p-4 md:p-6">
      {m.mods.map((mod, key) => (
        <ModCard key={key} mod={mod} />
      ))}
    </div>
  )
}
