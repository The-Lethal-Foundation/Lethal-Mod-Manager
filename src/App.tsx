import * as React from 'react'
import type { FC } from 'react'
import { useEffect, useState } from 'react'

import useGetProfiles from '@/hooks/getProfiles'
import { MainLayout } from '@/layouts/MainLayout'
import Sidebar from '@/layouts/sidebar-main'
import Header from '@/layouts/header-main'
import { useBlockUI } from '@/components/ui/block-ui'
import { ModList } from './features/mod-list'

import { toast } from 'sonner'
import { Toaster } from './components/ui/sonner'
import { ScrollArea } from '@/components/ui/scroll-area'

const App: FC = () => {
  const { isBlocked, theme, unblock } = useBlockUI('black', true)
  const p = useGetProfiles()
  const [selectedProfile, setSelectedProfile] = useState<string | null>(null)

  useEffect(() => {
    if (!p.isLoading || p.error) {
      console.log('Profiles list loaded')
    }
  }, [p.isLoading, p.error, unblock])

  useEffect(() => {
    window
      .init()
      .then((profileName: string) => {
        toast('🧙‍♂️ Spellbound!', {
          description: 'Your grimoire is ready for enchanting adventures.',
        })
        unblock()
        setSelectedProfile(profileName)
        console.log('Remembered Profile:', profileName)
      })
      .catch((out: string) => {
        toast('🤕 Whoops!', {
          description: `Something went wrong: ${out}`,
        })
      })
  }, [])

  return (
    <>
      <Toaster theme="dark" />
      <MainLayout
        sidebar={
          <Sidebar
            profiles={p.profiles}
            setProfile={setSelectedProfile}
            profile={selectedProfile}
          />
        }
        header={<Header />}
        blocking={{ isBlocked, theme }}
      >
        <ScrollArea
          style={{ maxHeight: 'calc(100vh - 50px - 20px)', overflowY: 'auto' }}
        >
          <h1 className="pl-4 pt-4 md:pl-6 md:pt-6 font-semibold text-lg md:text-2xl text-white mb-4">
            Mods
          </h1>

          {selectedProfile ? (
            <ModList profile={selectedProfile} />
          ) : (
            <div className="w-full">
              <span className="p-4 md:p-6 text-white">
                ✌️ No mods here yet. Maybe select another profile.
              </span>
            </div>
          )}
        </ScrollArea>
      </MainLayout>
    </>
  )
}

export default App
