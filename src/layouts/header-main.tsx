import React, { useEffect, useState } from 'react'
import type { FC } from 'react'
import { Input } from '../components/ui/input'
import { Button } from '@/components/ui/button'
import { Play } from 'lucide-react'
import { toast } from 'sonner'
import type { Tab } from '@/types/uiState'

export interface HeaderProps {
  profile: string | null
  currentTab: Tab
  setLocalModQuery: (q: string) => void
  setGlobalModQuery: (q: string) => void
}

const Header: FC<HeaderProps> = ({
  profile,
  currentTab,
  setLocalModQuery,
  setGlobalModQuery,
}) => {
  const runGame = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault() // Preventing default button click behavior

    if (!profile) {
      toast('⚠️ No profile selected', {
        description: 'Please select a profile to run the game.',
      })
      return
    }

    window
      .runGame(profile)
      .then(() => {
        toast('🏃‍♀️💨 Started!', {
          description: `Running profile ${profile}`,
        })
      })
      .catch((out: string) => {
        toast('🤕 Whoops!', {
          description: `Something went wrong: ${out}`,
        })
      })
  }

  const [input, setInput] = useState<string>('')

  useEffect(() => {
    setInput('')
  }, [currentTab])

  useEffect(() => {
    const timeoutId = setTimeout(() => {
      // Update the query based on the currentTab
      if (currentTab === 'local-mods') {
        setLocalModQuery(input)
      } else {
        setGlobalModQuery(input)
      }
    }, 500)

    return () => clearTimeout(timeoutId) // Cleanup the timeout
  }, [input, currentTab, setLocalModQuery, setGlobalModQuery])

  return (
    <header className="bg-[#09090b] flex h-14 lg:h-[60px] items-center gap-4 border-b border-[#27272a] px-4">
      <div className="w-full flex-1">
        <form>
          <div className="relative flex justify-between">
            <Input
              value={input}
              onChange={(e) => setInput(e.target.value)}
              type="search"
              placeholder="Search mods..."
              className="w-3/5 text-white border-[#27272a] focus:border-white"
            />
            {profile && (
              <Button onClick={runGame}>
                <Play size={24} />
              </Button>
            )}
          </div>
        </form>
      </div>
    </header>
  )
}

export default Header
