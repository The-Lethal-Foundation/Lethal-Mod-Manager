import React from 'react'
import type { FC } from 'react'
import { Input } from '../components/ui/input'
import { Button } from '@/components/ui/button'
import { Play } from 'lucide-react'
import { toast } from 'sonner'

export interface HeaderProps {
  profile: string | null
}

const Header: FC<HeaderProps> = ({ profile }) => {
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

  return (
    <header className="bg-[#09090b] flex h-14 lg:h-[60px] items-center gap-4 border-b border-[#27272a] px-4">
      <div className="w-full flex-1">
        <form>
          <div className="relative flex justify-between">
            <Input
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
