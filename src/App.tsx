import * as React from 'react'
import type { FC } from 'react'
import { useEffect, useState } from 'react'

import useGetProfiles from '@/hooks/getProfiles'
import { MainLayout } from '@/layouts/MainLayout'
import Sidebar from '@/layouts/sidebar-main'
import Header from '@/layouts/header-main'
import { useBlockUI } from '@/components/ui/block-ui'
import { ModList } from '@/features/mod-list'

import { toast } from 'sonner'
import { Toaster } from '@/components/ui/sonner'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Tab } from '@/types/uiState'
import { GlobalModList } from './features/mod-list-global'
import { section, type ordering } from './types/global'

const App: FC = () => {
  const { isBlocked, theme, unblock } = useBlockUI('black', true)
  const p = useGetProfiles()

  const [host, setHost] = useState<string | null>(null)
  const [selectedProfile, setSelectedProfile] = useState<string | null>(null)
  const [selectedTab, setSelectedTab] = useState<Tab>('local-mods')

  const [orderingState, setOrderingState] = useState<ordering>('top-rated')
  const [sectionState, setSectionState] = useState<section>('mods')
  const [globalQueryState, setGlobalQueryState] = useState<string>('')
  const [pageState] = useState(1)

  useEffect(() => {
    if (!p.isLoading || p.error) {
      console.log('Profiles list loaded')
    }
  }, [p.isLoading, p.error, unblock])

  useEffect(() => {
    window
      .init()
      .then((lastUsedProfile) => {
        toast('🧙‍♂️ Spellbound!', {
          description: 'Your grimoire is ready for enchanting adventures.',
        })
        setSelectedProfile(lastUsedProfile)
      })
      .catch((out: string) => {
        toast('🤕 Whoops!', {
          description: `Something went wrong: ${out}`,
        })
      })
    window.getAddr().then((addr) => {
      setHost(addr)
      unblock()
    })
  }, [])

  return (
    <>
      <Toaster theme="dark" />
      <MainLayout
        sidebar={
          <Sidebar
            profiles={p.profiles}
            profile={selectedProfile}
            setProfile={setSelectedProfile}
            refetchProfiles={p.fetchProfiles}
            selectedTab={selectedTab}
            setSelectedTab={setSelectedTab}
          />
        }
        header={
          <Header
            selectedProfile={selectedProfile}
            setSelectedProfile={setSelectedProfile}
            refetchProfiles={p.fetchProfiles}
            currentTab={selectedTab}
            setGlobalModQuery={setGlobalQueryState}
            setLocalModQuery={() => {}}
          />
        }
        blocking={{ isBlocked, theme }}
      >
        <ScrollArea
          style={{ maxHeight: 'calc(100vh - 50px - 20px)', overflowY: 'auto' }}
        >
          <h1 className="pl-4 pt-4 md:pl-6 md:pt-6 font-semibold text-lg md:text-2xl text-white">
            Mods
          </h1>

          {selectedTab === 'local-mods' && (
            <>
              {selectedProfile ? (
                <ModList
                  key={selectedProfile}
                  profile={selectedProfile}
                  host={host}
                />
              ) : (
                <div className="w-full mt-4">
                  <span className="p-4 md:p-6 text-white">
                    ✌️ No mods here yet. Maybe select another profile.
                  </span>
                </div>
              )}
            </>
          )}

          {selectedTab === 'global-mods' && (
            <GlobalModList
              sectionState={sectionState}
              setSection={setSectionState}
              ordering={orderingState}
              setOrdering={setOrderingState}
              query={globalQueryState}
              page={pageState}
              profile={selectedProfile}
            />
          )}
        </ScrollArea>
      </MainLayout>
    </>
  )
}

export default App
