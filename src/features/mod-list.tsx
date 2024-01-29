import React from 'react'

import { ModCard } from '@/features/mod-card'
import useGetMods from '@/hooks/getMods'
import { useEffect } from 'react'

interface ModListProps {
  profile: string
  host: string | null
}

export const ModList: React.FC<ModListProps> = ({ profile, host }) => {
  const m = useGetMods(profile)

  useEffect(() => {
    if (!m.isLoading || m.error) {
      if (m.error) {
        console.error(m.error)
      }

      console.log('Loaded mod list:', m.mods.length)
    }
  }, [m.isLoading, m.error, m.mods])

  return (
    <div className="w-full grid gap-6 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-white p-4 md:p-6">
      {m.mods.length > 0 ? (
        m.mods.map((mod, key) => (
          <ModCard
            key={key}
            profile={profile}
            mod={mod}
            image={`http://${host}/images/${profile}/BepInEx/plugins/${mod.mod_path_name}/icon.png`}
          />
        ))
      ) : (
        <p className="">🙊 No mods in this profile</p>
      )}
    </div>
  )
}
