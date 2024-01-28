import React from 'react'

import { BlockUI } from '@/components/ui/block-ui'
import type { FC, ReactNode } from 'react'

interface MainLayoutProps {
  sidebar: ReactNode
  header: ReactNode
  blocking: {
    isBlocked: boolean
    theme: 'black' | 'white'
  }
  children: ReactNode
}

const MainLayout: FC<MainLayoutProps> = ({
  sidebar,
  header,
  blocking,
  children,
}) => {
  return (
    <div className="grid min-h-screen w-full lg:grid-cols-[280px_1fr]">
      {sidebar}
      <div className="flex flex-col">
        {header}
        <main className="bg-[#09090b] flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-6">
          {children}
        </main>
      </div>

      <BlockUI isBlocked={blocking.isBlocked} theme={blocking.theme} />
    </div>
  )
}

export { MainLayout }
