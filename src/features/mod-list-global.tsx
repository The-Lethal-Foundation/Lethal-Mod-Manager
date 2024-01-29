import React, { useEffect } from 'react'
import { useGetGlobalMods } from '@/hooks/getGlobalMods'
import { GlobalModCard } from './mod-card-global'

interface GlobalModListProps {
  page: number
  ordering: string
  profile: string | null
}

const GlobalModList: React.FC<GlobalModListProps> = ({
  page,
  ordering,
  profile,
}) => {
  const { globalMods, isLoading, error } = useGetGlobalMods(page, ordering)

  useEffect(() => {
    if (!isLoading || error) {
      if (error) {
        console.error(error)
      }

      console.log('Loaded global mod list:', globalMods.length)
    }
  }, [isLoading, error, globalMods])

  return (
    <div className="w-full grid gap-6 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 text-white p-4 md:p-6">
      {globalMods.length > 0 ? (
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
  )
}

export { GlobalModList }
