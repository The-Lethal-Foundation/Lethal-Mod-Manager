import * as React from 'react'
import type { FC } from 'react'
import { useEffect, useState } from 'react'

import useGetProfiles from '@/hooks/getProfiles'
import { MainLayout } from '@/layouts/MainLayout'
import Sidebar from '@/layouts/sidebar-main'
import Header from '@/layouts/header-main'
import { useBlockUI } from '@/components/ui/block-ui'
import { ModList } from './features/mod-list'

const App: FC = () => {
  const { isBlocked, theme, unblock } = useBlockUI('black', true)
  const p = useGetProfiles()
  const [selectedProfile, setSelectedProfile] = useState<string | null>(null)

  useEffect(() => {
    if (!p.isLoading || p.error) {
      unblock()
    }
  }, [p.isLoading, p.error, unblock])

  return (
    <>
      <MainLayout
        sidebar={
          <Sidebar profiles={p.profiles} setProfile={setSelectedProfile} />
        }
        header={<Header />}
        blocking={{ isBlocked, theme }}
      >
        <h1 className="font-semibold text-lg md:text-2xl text-white">Mods</h1>

        {selectedProfile ? (
          <ModList profile={selectedProfile} />
        ) : (
          <div className="w-full">
            <span className="text-white">
              ✌️ No mods here yet. Maybe select another profile.
            </span>
          </div>
        )}
      </MainLayout>
    </>
  )
}

export default App
